package bonds

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types2.MsgCreateBond:
			return handleMsgCreateBond(ctx, keeper, msg)
		case types2.MsgEditBond:
			return handleMsgEditBond(ctx, keeper, msg)
		case types2.MsgSetNextAlpha:
			return handleMsgSetNextAlpha(ctx, keeper, msg)
		case types2.MsgUpdateBondState:
			return handleMsgUpdateBondState(ctx, keeper, msg)
		case types2.MsgBuy:
			return handleMsgBuy(ctx, keeper, msg)
		case types2.MsgSell:
			return handleMsgSell(ctx, keeper, msg)
		case types2.MsgSwap:
			return handleMsgSwap(ctx, keeper, msg)
		case types2.MsgMakeOutcomePayment:
			return handleMsgMakeOutcomePayment(ctx, keeper, msg)
		case types2.MsgWithdrawShare:
			return handleMsgWithdrawShare(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized bonds Msg type: %v", msg.Type())
		}
	}
}

func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) []abci.ValidatorUpdate {

	iterator := keeper.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		bond := keeper.MustGetBondByKey(ctx, iterator.Key())
		batch := keeper.MustGetBatch(ctx, bond.BondDid)

		// Subtract one block
		batch.BlocksRemaining = batch.BlocksRemaining.SubUint64(1)
		keeper.SetBatch(ctx, bond.BondDid, batch)

		// If blocks remaining > 0 do not perform orders
		if !batch.BlocksRemaining.IsZero() {
			continue
		}

		// Store current reserve to check if this has changed later on
		reserveBeforeOrderProcessing := bond.CurrentReserve

		// Perform orders
		keeper.PerformOrders(ctx, bond.BondDid)

		// Get bond again just in case current supply was updated
		// Get batch again just in case orders were cancelled
		bond = keeper.MustGetBond(ctx, bond.BondDid)
		batch = keeper.MustGetBatch(ctx, bond.BondDid)

		// For augmented, if hatch phase and newSupply >= S0, go to open phase
		if bond.FunctionType == types2.AugmentedFunction &&
			bond.State == types2.HatchState {
			args := bond.FunctionParameters.AsMap()
			if bond.CurrentSupply.Amount.ToDec().GTE(args["S0"]) {
				keeper.SetBondState(ctx, bond.BondDid, types2.OpenState)
				bond = keeper.MustGetBond(ctx, bond.BondDid) // get updated bond
			}
		}

		// Update alpha value if in open state and next alpha is not null
		if bond.State == types2.OpenState && batch.HasNextAlpha() {
			keeper.UpdateAlpha(ctx, bond.BondDid)
		}

		// Save current batch as last batch and reset current batch
		keeper.SetLastBatch(ctx, bond.BondDid, batch)
		keeper.SetBatch(ctx, bond.BondDid, types2.NewBatch(bond.BondDid, bond.Token, bond.BatchBlocks))

		// If reserve has not changed, no need to recalculate I0; rest of function can be skipped
		if bond.CurrentReserve.IsEqual(reserveBeforeOrderProcessing) {
			continue
		}

		// Recalculate and re-set I0 if alpha bond
		if bond.AlphaBond {
			paramsMap := bond.FunctionParameters.AsMap()
			newI0 := types2.InvariantI(bond.OutcomePayment, paramsMap["systemAlpha"],
				bond.CurrentReserve[0].Amount)
			bond.FunctionParameters.ReplaceParam("I0", newI0)
		}

		// Save bond
		keeper.SetBond(ctx, bond.BondDid, bond)
	}
	return []abci.ValidatorUpdate{}
}

func handleMsgCreateBond(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgCreateBond) (*sdk.Result, error) {
	if keeper.BankKeeper.BlacklistedAddr(msg.FeeAddress) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive transactions", msg.FeeAddress)
	}

	// Check that bond and bond DID do not already exist
	if keeper.BondExists(ctx, msg.BondDid) {
		return nil, sdkerrors.Wrap(types2.ErrBondAlreadyExists, msg.BondDid)
	} else if keeper.BondDidExists(ctx, msg.Token) {
		return nil, sdkerrors.Wrap(types2.ErrBondTokenIsTaken, msg.Token)
	} else if msg.Token == keeper.StakingKeeper.GetParams(ctx).BondDenom {
		return nil, types2.ErrBondTokenCannotBeStakingToken
	}

	// Check that bond token not reserved
	if keeper.ReservedBondToken(ctx, msg.Token) {
		return nil, types2.ErrReservedBondToken
	}

	// Set state to open by default (overridden below if augmented function)
	state := types2.OpenState

	// If augmented, add R0, S0, V0 as parameters for quick access
	// Also, override AllowSells and set to False if S0 > 0
	if msg.FunctionType == types2.AugmentedFunction {
		paramsMap := msg.FunctionParameters.AsMap()
		d0, _ := paramsMap["d0"]
		p0, _ := paramsMap["p0"]
		theta, _ := paramsMap["theta"]
		kappa, _ := paramsMap["kappa"]

		R0 := d0.Mul(sdk.OneDec().Sub(theta))
		S0 := d0.Quo(p0)
		V0, err := types2.Invariant(R0, S0, kappa)
		if err != nil {
			return nil, err
		}

		msg.FunctionParameters = msg.FunctionParameters.AddParams(
			types2.FunctionParams{
				types2.NewFunctionParam("R0", R0),
				types2.NewFunctionParam("S0", S0),
				types2.NewFunctionParam("V0", V0),
			})

		if msg.AlphaBond {
			publicAlpha := types2.StartingPublicAlpha
			systemAlpha := types2.SystemAlpha(publicAlpha, sdk.OneInt(),
				sdk.OneInt(), R0.TruncateInt(), msg.OutcomePayment)

			I0 := types2.InvariantI(msg.OutcomePayment, systemAlpha, sdk.ZeroInt())

			msg.FunctionParameters = msg.FunctionParameters.AddParams(
				types2.FunctionParams{
					types2.NewFunctionParam("I0", I0),
					types2.NewFunctionParam("publicAlpha", publicAlpha),
					types2.NewFunctionParam("systemAlpha", systemAlpha),
				})
		}

		// The starting state for augmented bonding curves is the Hatch state.
		// Note that we can never start with OpenState since S0>0 (S0=d0/p0 and d0>0).
		state = types2.HatchState
	}

	bond := types2.NewBond(msg.Token, msg.Name, msg.Description, msg.CreatorDid,
		msg.ControllerDid, msg.FunctionType, msg.FunctionParameters,
		msg.ReserveTokens, msg.TxFeePercentage, msg.ExitFeePercentage,
		msg.FeeAddress, msg.MaxSupply, msg.OrderQuantityLimits, msg.SanityRate,
		msg.SanityMarginPercentage, msg.AllowSells, msg.AlphaBond, msg.BatchBlocks,
		msg.OutcomePayment, state, msg.BondDid)

	keeper.SetBond(ctx, bond.BondDid, bond)
	keeper.SetBondDid(ctx, bond.Token, bond.BondDid)
	keeper.SetBatch(ctx, bond.BondDid, types2.NewBatch(bond.BondDid, bond.Token, msg.BatchBlocks))

	logger := keeper.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s [%s] with reserve(s) [%s] created by %s", msg.Token,
		msg.FunctionType, strings.Join(bond.ReserveTokens, ","), msg.CreatorDid))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeCreateBond,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types2.AttributeKeyToken, msg.Token),
			sdk.NewAttribute(types2.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types2.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types2.AttributeKeyFunctionType, msg.FunctionType),
			sdk.NewAttribute(types2.AttributeKeyFunctionParameters, msg.FunctionParameters.String()),
			sdk.NewAttribute(types2.AttributeKeyCreatorDid, msg.CreatorDid),
			sdk.NewAttribute(types2.AttributeKeyControllerDid, msg.ControllerDid),
			sdk.NewAttribute(types2.AttributeKeyReserveTokens, types2.StringsToString(msg.ReserveTokens)),
			sdk.NewAttribute(types2.AttributeKeyTxFeePercentage, msg.TxFeePercentage.String()),
			sdk.NewAttribute(types2.AttributeKeyExitFeePercentage, msg.ExitFeePercentage.String()),
			sdk.NewAttribute(types2.AttributeKeyFeeAddress, msg.FeeAddress.String()),
			sdk.NewAttribute(types2.AttributeKeyMaxSupply, msg.MaxSupply.String()),
			sdk.NewAttribute(types2.AttributeKeyOrderQuantityLimits, msg.OrderQuantityLimits.String()),
			sdk.NewAttribute(types2.AttributeKeySanityRate, msg.SanityRate.String()),
			sdk.NewAttribute(types2.AttributeKeySanityMarginPercentage, msg.SanityMarginPercentage.String()),
			sdk.NewAttribute(types2.AttributeKeyAllowSells, strconv.FormatBool(msg.AllowSells)),
			sdk.NewAttribute(types2.AttributeKeyAlphaBond, strconv.FormatBool(msg.AlphaBond)),
			sdk.NewAttribute(types2.AttributeKeyBatchBlocks, msg.BatchBlocks.String()),
			sdk.NewAttribute(types2.AttributeKeyOutcomePayment, msg.OutcomePayment.String()),
			sdk.NewAttribute(types2.AttributeKeyState, string(state)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.CreatorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgEditBond(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgEditBond) (*sdk.Result, error) {

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	if bond.CreatorDid != msg.EditorDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the creator of the bond")
	}

	if msg.Name != types2.DoNotModifyField {
		bond.Name = msg.Name
	}
	if msg.Description != types2.DoNotModifyField {
		bond.Description = msg.Description
	}

	if msg.OrderQuantityLimits != types2.DoNotModifyField {
		orderQuantityLimits, err := sdk.ParseCoins(msg.OrderQuantityLimits)
		if err != nil {
			return nil, err
		}
		bond.OrderQuantityLimits = orderQuantityLimits
	}

	if msg.SanityRate != types2.DoNotModifyField {
		var sanityRate, sanityMarginPercentage sdk.Dec
		if msg.SanityRate == "" {
			sanityRate = sdk.ZeroDec()
			sanityMarginPercentage = sdk.ZeroDec()
		} else {
			parsedSanityRate, err := sdk.NewDecFromStr(msg.SanityRate)
			if err != nil {
				return nil, sdkerrors.Wrap(types2.ErrArgumentMissingOrNonFloat, "sanity rate")
			} else if parsedSanityRate.IsNegative() {
				return nil, sdkerrors.Wrap(types2.ErrArgumentCannotBeNegative, "sanity rate")
			}
			parsedSanityMarginPercentage, err := sdk.NewDecFromStr(msg.SanityMarginPercentage)
			if err != nil {
				return nil, sdkerrors.Wrap(types2.ErrArgumentMissingOrNonFloat, "sanity margin percentage")
			} else if parsedSanityMarginPercentage.IsNegative() {
				return nil, sdkerrors.Wrap(types2.ErrArgumentCannotBeNegative, "sanity margin percentage")
			}
			sanityRate = parsedSanityRate
			sanityMarginPercentage = parsedSanityMarginPercentage
		}
		bond.SanityRate = sanityRate
		bond.SanityMarginPercentage = sanityMarginPercentage
	}

	keeper.SetBond(ctx, bond.BondDid, bond)

	logger := keeper.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s edited by %s",
		msg.BondDid, msg.EditorDid))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeEditBond,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types2.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types2.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types2.AttributeKeyOrderQuantityLimits, msg.OrderQuantityLimits),
			sdk.NewAttribute(types2.AttributeKeySanityRate, msg.SanityRate),
			sdk.NewAttribute(types2.AttributeKeySanityMarginPercentage, msg.SanityMarginPercentage),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.EditorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSetNextAlpha(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgSetNextAlpha) (*sdk.Result, error) {

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	newPublicAlpha := msg.Alpha

	if bond.FunctionType != types2.AugmentedFunction {
		return nil, sdkerrors.Wrap(types2.ErrFunctionNotAvailableForFunctionType,
			"bond is not an augmented bonding curve")
	} else if !bond.AlphaBond {
		return nil, sdkerrors.Wrap(types2.ErrFunctionNotAvailableForFunctionType,
			"bond is not an alpha bond")
	} else if bond.State != types2.OpenState {
		return nil, types2.ErrInvalidStateForAction
	}

	if bond.ControllerDid != msg.EditorDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the controller of the bond")
	}

	// Get supply, reserve, outcome payment. Note that we get the adjusted
	// supply in order to take into consideration the influence of the buys and
	// sells in the current batch. We then get the reserve based on this supply.
	S := keeper.GetSupplyAdjustedForAlphaEdit(ctx, bond.BondDid).Amount
	R, err := bond.ReserveAtSupply(S)
	if err != nil {
		return nil, err
	}
	C := bond.OutcomePayment

	// Get current parameters
	paramsMap := bond.FunctionParameters.AsMap()

	// Check (newPublicAlpha != publicAlpha)
	if newPublicAlpha.Equal(paramsMap["publicAlpha"]) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidAlpha,
			"cannot change public alpha to the current value of public alpha")
	}

	// Calculate scaled delta public alpha, to calculate new system alpha
	prevPublicAlpha := paramsMap["publicAlpha"]
	deltaPublicAlpha := newPublicAlpha.Sub(prevPublicAlpha)
	temp, err := types2.ApproxPower(
		prevPublicAlpha.Mul(sdk.OneDec().Sub(types2.StartingPublicAlpha)),
		sdk.MustNewDecFromStr("2"))
	if err != nil {
		return nil, err
	}
	scaledDeltaPublicAlpha := deltaPublicAlpha.Mul(temp)

	// Calculate new system alpha
	prevSystemAlpha := paramsMap["systemAlpha"]
	var newSystemAlpha sdk.Dec
	if deltaPublicAlpha.IsPositive() {
		// 1 - (1 - scaled_delta_public_alpha) * (1 - previous_alpha)
		temp1 := sdk.OneDec().Sub(scaledDeltaPublicAlpha)
		temp2 := sdk.OneDec().Sub(prevSystemAlpha)
		newSystemAlpha = sdk.OneDec().Sub(temp1.Mul(temp2))
	} else {
		// (1 - scaled_delta_public_alpha) * (previous_alpha)
		temp1 := sdk.OneDec().Sub(scaledDeltaPublicAlpha)
		temp2 := prevSystemAlpha
		newSystemAlpha = temp1.Mul(temp2)
	}

	// Check 1 (newSystemAlpha != prevSystemAlpha)
	if newSystemAlpha.Equal(prevSystemAlpha) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidAlpha,
			"resultant system alpha based on public alpha is unchanged")
	}
	// Check 2 (I > C * newSystemAlpha)
	if paramsMap["I0"].LTE(newSystemAlpha.MulInt(C)) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidAlpha,
			"cannot change alpha to that value due to violated restriction [1]")
	}
	// Check 3 (R / C > newSystemAlpha - prevSystemAlpha)
	if R.QuoInt(C).LTE(newSystemAlpha.Sub(prevSystemAlpha)) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidAlpha,
			"cannot change alpha to that value due to violated restriction [2]")
	}

	// Recalculate kappa and V0 using new alpha
	newKappa := types2.Kappa(paramsMap["I0"], C, newSystemAlpha)
	_, err = types2.Invariant(R, S.ToDec(), newKappa)
	if err != nil {
		return nil, err
	}

	// Get batch to set new alpha
	batch := keeper.MustGetBatch(ctx, bond.BondDid)
	batch.NextPublicAlpha = newPublicAlpha
	keeper.SetBatch(ctx, bond.BondDid, batch)

	logger := keeper.Logger(ctx)
	logger.Info(fmt.Sprintf("bond %s next alpha set by %s",
		msg.BondDid, msg.EditorDid))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeSetNextAlpha,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types2.AttributeKeyPublicAlpha, newPublicAlpha.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.EditorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUpdateBondState(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgUpdateBondState) (*sdk.Result, error) {

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}
	batch := keeper.MustGetBatch(ctx, msg.BondDid)

	if bond.FunctionType != types2.AugmentedFunction {
		return nil, types2.ErrFunctionNotAvailableForFunctionType
	} else if !msg.State.IsValidProgressionFrom(bond.State) {
		return nil, types2.ErrInvalidStateProgression
	} // Also, next state must be SETTLE or FAILED -- checked by ValidateBasic

	if bond.ControllerDid != msg.EditorDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"editor must be the controller of the bond")
	}

	// If state is settle or failed, move all outcome payment to reserve, so
	// that it is available for share withdrawal (MsgWithdrawShare)
	if msg.State == types2.SettleState || msg.State == types2.FailedState {
		if !batch.Empty() {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				"cannot update bond state to SETTLE/FAILED while there are orders in the batch")
		}
		keeper.MoveOutcomePaymentToReserve(ctx, bond.BondDid)
	}

	// Update bond state
	keeper.SetBondState(ctx, bond.BondDid, msg.State)

	ctx.EventManager().EmitEvents(sdk.Events{
		// No need to emit event/log for state change, as SetBondState does this
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.EditorDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgBuy(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgBuy) (*sdk.Result, error) {
	buyerAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.BuyerDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types2.ErrBondTokenDoesNotMatchBond
	}

	// Check current state is HATCH/OPEN, max prices, order quantity limits
	if bond.State != types2.OpenState && bond.State != types2.HatchState {
		return nil, types2.ErrInvalidStateForAction
	} else if !bond.ReserveDenomsEqualTo(msg.MaxPrices) {
		return nil, sdkerrors.Wrap(types2.ErrReserveDenomsMismatch, msg.MaxPrices.String())
	} else if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.Amount}) {
		return nil, types2.ErrOrderQuantityLimitExceeded
	}

	// For the swapper, the first buy is the initialisation of the reserves
	// The max prices are used as the actual prices and one token is minted
	// The amount of token serves to define the price of adding more liquidity
	if bond.CurrentSupply.IsZero() && bond.FunctionType == types2.SwapperFunction {
		return performFirstSwapperFunctionBuy(ctx, keeper, msg)
	}

	// Take max that buyer is willing to pay (enforces maxPrice <= balance)
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, buyerAddr,
		types2.BatchesIntermediaryAccount, msg.MaxPrices)
	if err != nil {
		return nil, err
	}

	// Create order
	order := types2.NewBuyOrder(msg.BuyerDid, msg.Amount, msg.MaxPrices)

	// Get buy price and check if can add buy order to batch
	buyPrices, sellPrices, err := keeper.GetUpdatedBatchPricesAfterBuy(ctx, bond.BondDid, order)
	if err != nil {
		return nil, err
	}

	// Add buy order to batch
	keeper.AddBuyOrder(ctx, bond.BondDid, order, buyPrices, sellPrices)

	// Cancel unfulfillable orders
	keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeBuy,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types2.AttributeKeyMaxPrices, msg.MaxPrices.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.BuyerDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func performFirstSwapperFunctionBuy(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgBuy) (*sdk.Result, error) {
	buyerAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.BuyerDid).Address()

	// TODO: investigate effect that a high amount has on future buyers' ability to buy.

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types2.ErrBondTokenDoesNotMatchBond
	}

	// Check if initial liquidity violates sanity rate
	if bond.ReservesViolateSanityRate(msg.MaxPrices) {
		return nil, types2.ErrValuesViolateSanityRate
	}

	// Use max prices as the amount to send to the liquidity pool (i.e. price)
	err := keeper.DepositReserve(ctx, bond.BondDid, buyerAddr, msg.MaxPrices)
	if err != nil {
		return nil, err
	}

	// Mint bond tokens
	err = keeper.SupplyKeeper.MintCoins(ctx, types2.BondsMintBurnAccount,
		sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Send bond tokens to buyer
	err = keeper.SupplyKeeper.SendCoinsFromModuleToAccount(ctx,
		types2.BondsMintBurnAccount, buyerAddr, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Update supply
	keeper.SetCurrentSupply(ctx, bond.BondDid, bond.CurrentSupply.Add(msg.Amount))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeInitSwapper,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
			sdk.NewAttribute(types2.AttributeKeyChargedPrices, msg.MaxPrices.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.BuyerDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSell(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgSell) (*sdk.Result, error) {
	sellerAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.SellerDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check sells allowed, current state is OPEN, and order limits not exceeded
	if !bond.AllowSells {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotAllowSelling, msg.BondDid)
	} else if bond.State != types2.OpenState {
		return nil, types2.ErrInvalidStateForAction
	} else if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.Amount}) {
		return nil, types2.ErrOrderQuantityLimitExceeded
	}

	// Check that bond token used belongs to this bond
	if msg.Amount.Denom != bond.Token {
		return nil, types2.ErrBondTokenDoesNotMatchBond
	}

	// Send coins to be burned from seller (enforces sellAmount <= balance)
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, sellerAddr,
		types2.BondsMintBurnAccount, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Burn bond tokens to be sold
	err = keeper.SupplyKeeper.BurnCoins(ctx, types2.BondsMintBurnAccount,
		sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	// Create order
	order := types2.NewSellOrder(msg.SellerDid, msg.Amount)

	// Get sell price and check if can add sell order to batch
	buyPrices, sellPrices, err := keeper.GetUpdatedBatchPricesAfterSell(ctx, bond.BondDid, order)
	if err != nil {
		return nil, err
	}

	// Add sell order to batch
	keeper.AddSellOrder(ctx, bond.BondDid, order, buyPrices, sellPrices)

	//// Cancel unfulfillable orders (Note: no need)
	//keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeSell,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SellerDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSwap(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgSwap) (*sdk.Result, error) {
	swapperAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.SwapperDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrap(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	// Confirm that function type is swapper_function and state is OPEN
	if bond.FunctionType != types2.SwapperFunction {
		return nil, types2.ErrFunctionNotAvailableForFunctionType
	} else if bond.State != types2.OpenState {
		return nil, types2.ErrInvalidStateForAction
	}

	// Check that from and to use reserve token names
	fromAndTo := sdk.NewCoins(msg.From, sdk.NewCoin(msg.ToToken, sdk.OneInt()))
	fromAndToDenoms := msg.From.Denom + "," + msg.ToToken
	if !bond.ReserveDenomsEqualTo(fromAndTo) {
		return nil, sdkerrors.Wrap(types2.ErrReserveDenomsMismatch, fromAndToDenoms)
	}

	// Check if order quantity limit exceeded
	if bond.AnyOrderQuantityLimitsExceeded(sdk.Coins{msg.From}) {
		return nil, types2.ErrOrderQuantityLimitExceeded
	}

	// Take coins to be swapped from swapper (enforces swapAmount <= balance)
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(ctx, swapperAddr,
		types2.BatchesIntermediaryAccount, sdk.Coins{msg.From})
	if err != nil {
		return nil, err
	}

	// Create order
	order := types2.NewSwapOrder(msg.SwapperDid, msg.From, msg.ToToken)

	// Add swap order to batch
	keeper.AddSwapOrder(ctx, bond.BondDid, order)

	//// Cancel unfulfillable orders (Note: no need)
	//keeper.CancelUnfulfillableOrders(ctx, bond.BondDid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeSwap,
			sdk.NewAttribute(types2.AttributeKeyBondDid, bond.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.From.Amount.String()),
			sdk.NewAttribute(types2.AttributeKeySwapFromToken, msg.From.Denom),
			sdk.NewAttribute(types2.AttributeKeySwapToToken, msg.ToToken),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SwapperDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgMakeOutcomePayment(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgMakeOutcomePayment) (*sdk.Result, error) {
	senderAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.SenderDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	// Confirm that state is OPEN and that outcome payment is not nil
	if bond.State != types2.OpenState {
		return nil, types2.ErrInvalidStateForAction
	}

	// Send outcome payment to outcome payment reserve
	outcomePayment := bond.GetNewReserveCoins(msg.Amount)
	err := keeper.DepositOutcomePayment(ctx, bond.BondDid, senderAddr, outcomePayment)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeMakeOutcomePayment,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(sdk.AttributeKeyAmount, outcomePayment.String()),
			sdk.NewAttribute(types2.AttributeKeyAddress, senderAddr.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.SenderDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgWithdrawShare(ctx sdk.Context, keeper keeper.Keeper, msg types2.MsgWithdrawShare) (*sdk.Result, error) {
	recipientAddr := keeper.DidKeeper.MustGetDidDoc(ctx, msg.RecipientDid).Address()

	bond, found := keeper.GetBond(ctx, msg.BondDid)
	if !found {
		return nil, sdkerrors.Wrapf(types2.ErrBondDoesNotExist, msg.BondDid)
	}

	// Check that state is SETTLE or FAILED
	if bond.State != types2.SettleState && bond.State != types2.FailedState {
		return nil, types2.ErrInvalidStateForAction
	}

	// Get number of bond tokens owned by the recipient
	bondTokensOwnedAmount := keeper.BankKeeper.GetCoins(ctx, recipientAddr).AmountOf(bond.Token)
	if bondTokensOwnedAmount.IsZero() {
		return nil, types2.ErrNoBondTokensOwned
	}
	bondTokensOwned := sdk.NewCoin(bond.Token, bondTokensOwnedAmount)

	// Send coins to be burned from recipient
	err := keeper.SupplyKeeper.SendCoinsFromAccountToModule(
		ctx, recipientAddr, types2.BondsMintBurnAccount, sdk.NewCoins(bondTokensOwned))
	if err != nil {
		return nil, err
	}

	// Burn bond tokens
	err = keeper.SupplyKeeper.BurnCoins(ctx, types2.BondsMintBurnAccount,
		sdk.NewCoins(sdk.NewCoin(bond.Token, bondTokensOwnedAmount)))
	if err != nil {
		return nil, err
	}

	// Calculate amount owned
	remainingReserve := keeper.GetReserveBalances(ctx, bond.BondDid)
	bondTokensShare := bondTokensOwnedAmount.ToDec().QuoInt(bond.CurrentSupply.Amount)
	reserveOwedDec := sdk.NewDecCoinsFromCoins(remainingReserve...).MulDec(bondTokensShare)
	reserveOwed, _ := reserveOwedDec.TruncateDecimal()

	// Send coins owed to recipient
	err = keeper.WithdrawReserve(ctx, bond.BondDid, recipientAddr, reserveOwed)
	if err != nil {
		return nil, err
	}

	// Update supply
	keeper.SetCurrentSupply(ctx, bond.BondDid, bond.CurrentSupply.Sub(bondTokensOwned))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types2.EventTypeWithdrawShare,
			sdk.NewAttribute(types2.AttributeKeyBondDid, msg.BondDid),
			sdk.NewAttribute(types2.AttributeKeyAddress, recipientAddr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, reserveOwed.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.RecipientDid),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
