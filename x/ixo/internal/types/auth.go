package types

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto"
	ed25519tm "github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/multisig"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"os"
)

var (
	expectedMinGasPrices       = "0.025" + IxoNativeToken
	approximationGasAdjustment = float64(1.5)
	// TODO: parameterise (or remove) hard-coded gas prices and adjustments

	// simulation pubkey to estimate gas consumption
	simEd25519Pubkey ed25519tm.PubKeyEd25519
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey[:], bz)
}

type PubKeyGetter func(ctx sdk.Context, msg IxoMsg) (crypto.PubKey, error)

func NewDefaultPubKeyGetter(didKeeper DidKeeper) PubKeyGetter {
	return func(ctx sdk.Context, msg IxoMsg) (pubKey crypto.PubKey, err error) {

		signerDidDoc, err := didKeeper.GetDidDoc(ctx, msg.GetSignerDid())
		if err != nil {
			return pubKey, err
		}

		var pubKeyRaw ed25519tm.PubKeyEd25519
		copy(pubKeyRaw[:], base58.Decode(signerDidDoc.GetPubKey()))
		return pubKeyRaw, nil
	}
}

func NewDefaultAnteHandler(ak auth.AccountKeeper, supplyKeeper supply.Keeper, sigGasConsumer ante.SignatureVerificationGasConsumer, pubKeyGetter PubKeyGetter) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		NewDeductFeeDecorator(ak, supplyKeeper, pubKeyGetter),
		NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
		NewSigVerificationDecorator(ak, pubKeyGetter),
		NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
	)
}

func ApproximateFeeForTx(cliCtx context.CLIContext, tx auth.StdTx, chainId string) (auth.StdFee, error) {

	// Set up a transaction builder
	cdc := cliCtx.Codec
	txEncoder := auth.DefaultTxEncoder
	gasAdjustment := approximationGasAdjustment
	fees := sdk.NewCoins(sdk.NewCoin(IxoNativeToken, sdk.OneInt()))
	txBldr := auth.NewTxBuilder(txEncoder(cdc), 0, 0, 0, gasAdjustment, true, chainId, tx.Memo, fees, nil)

	// Approximate gas consumption
	txBldr, err := utils.EnrichWithGas(txBldr, cliCtx, tx.Msgs)
	if err != nil {
		return auth.StdFee{}, err
	}

	// Clear fees and set gas-prices to deduce updated fee = (gas * gas-prices)
	signMsg, err := txBldr.WithFees("").WithGasPrices(expectedMinGasPrices).BuildSignMsg(tx.Msgs)
	if err != nil {
		return auth.StdFee{}, err
	}

	return signMsg.Fee, nil
}

func GenerateOrBroadcastMsgs(cliCtx context.CLIContext, msg sdk.Msg, ixoDid exported.IxoDid) error {
	msgs := []sdk.Msg{msg}
	txBldr := auth.NewTxBuilderFromCLI(cliCtx.Input).WithTxEncoder(utils.GetTxEncoder(cliCtx.Codec))

	if cliCtx.GenerateOnly {
		return utils.PrintUnsignedStdTx(txBldr, cliCtx, msgs)
	}

	return CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs, ixoDid)
}

func Sign(cliCtx context.CLIContext, msg auth.StdSignMsg,
	ixoDid exported.IxoDid) ([]byte, error) {
	var privateKey ed25519tm.PrivKeyEd25519
	copy(privateKey[:], base58.Decode(ixoDid.Secret.SignKey))
	copy(privateKey[32:], base58.Decode(ixoDid.VerifyKey))

	sig, err := MakeSignature(msg.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	encoder := utils.GetTxEncoder(cliCtx.Codec)
	return encoder(auth.NewStdTx(msg.Msgs, msg.Fee, []auth.StdSignature{sig}, msg.Memo))
}

func BuildAndSign(txBldr auth.TxBuilder, ctx context.CLIContext,
	msgs []sdk.Msg, ixoDid exported.IxoDid) ([]byte, error) {
	msg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, err
	}

	return Sign(ctx, msg, ixoDid)
}

func CompleteAndBroadcastTxCLI(txBldr auth.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg, ixoDid exported.IxoDid) error {
	txBldr, err := utils.PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return err
	}

	if txBldr.SimulateAndExecute() || cliCtx.Simulate {
		txBldr, err = utils.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			return err
		}

		gasEst := utils.GasEstimateResponse{GasEstimate: txBldr.Gas()}
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", gasEst.String())
	}

	if cliCtx.Simulate {
		return nil
	}

	if !cliCtx.SkipConfirm {
		stdSignMsg, err := txBldr.BuildSignMsg(msgs)
		if err != nil {
			return err
		}

		var json []byte
		if viper.GetBool(flags.FlagIndentResponse) {
			json, err = cliCtx.Codec.MarshalJSONIndent(stdSignMsg, "", "  ")
			if err != nil {
				panic(err)
			}
		} else {
			json = cliCtx.Codec.MustMarshalJSON(stdSignMsg)
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", json)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf)
		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return err
		}
	}

	// build and sign the transaction
	txBytes, err := BuildAndSign(txBldr, cliCtx, msgs, ixoDid)
	if err != nil {
		return err
	}

	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err
	}

	return cliCtx.PrintOutput(res)
}

func SignAndBroadcastTxFromStdSignMsg(cliCtx context.CLIContext,
	msg auth.StdSignMsg, ixoDid exported.IxoDid) (sdk.TxResponse, error) {

	// sign the transaction
	txBytes, err := Sign(cliCtx, msg, ixoDid)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	return res, nil
}

func MakeSignature(signBytes []byte,
	privateKey ed25519tm.PrivKeyEd25519) (auth.StdSignature, error) {
	sig, err := privateKey.Sign(signBytes)
	if err != nil {
		return auth.StdSignature{}, err
	}

	return auth.StdSignature{
		PubKey:    privateKey.PubKey(),
		Signature: sig,
	}, nil
}

// Identical to DefaultSigVerificationGasConsumer, but with ed25519 allowed
func IxoSigVerificationGasConsumer(
	meter sdk.GasMeter, sig []byte, pubkey crypto.PubKey, params auth.Params,
) error {
	switch pubkey := pubkey.(type) {
	case ed25519tm.PubKeyEd25519:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
		return nil

	case secp256k1.PubKeySecp256k1:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
		return nil

	case multisig.PubKeyMultisigThreshold:
		var multisignature multisig.Multisignature
		codec.Cdc.MustUnmarshalBinaryBare(sig, &multisignature)

		ante.ConsumeMultisignatureVerificationGas(meter, multisignature, pubkey, params)
		return nil

	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "unrecognized public key type: %T", pubkey)
	}
}
