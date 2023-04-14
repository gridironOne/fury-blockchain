# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [fury/bonds/v1beta1/bonds.proto](#fury/bonds/v1beta1/bonds.proto)
    - [BaseOrder](#fury.bonds.v1beta1.BaseOrder)
    - [Batch](#fury.bonds.v1beta1.Batch)
    - [Bond](#fury.bonds.v1beta1.Bond)
    - [BondDetails](#fury.bonds.v1beta1.BondDetails)
    - [BuyOrder](#fury.bonds.v1beta1.BuyOrder)
    - [FunctionParam](#fury.bonds.v1beta1.FunctionParam)
    - [Params](#fury.bonds.v1beta1.Params)
    - [SellOrder](#fury.bonds.v1beta1.SellOrder)
    - [SwapOrder](#fury.bonds.v1beta1.SwapOrder)
  
- [fury/bonds/v1beta1/genesis.proto](#fury/bonds/v1beta1/genesis.proto)
    - [GenesisState](#fury.bonds.v1beta1.GenesisState)
  
- [fury/bonds/v1beta1/query.proto](#fury/bonds/v1beta1/query.proto)
    - [QueryAlphaMaximumsRequest](#fury.bonds.v1beta1.QueryAlphaMaximumsRequest)
    - [QueryAlphaMaximumsResponse](#fury.bonds.v1beta1.QueryAlphaMaximumsResponse)
    - [QueryAvailableReserveRequest](#fury.bonds.v1beta1.QueryAvailableReserveRequest)
    - [QueryAvailableReserveResponse](#fury.bonds.v1beta1.QueryAvailableReserveResponse)
    - [QueryBatchRequest](#fury.bonds.v1beta1.QueryBatchRequest)
    - [QueryBatchResponse](#fury.bonds.v1beta1.QueryBatchResponse)
    - [QueryBondRequest](#fury.bonds.v1beta1.QueryBondRequest)
    - [QueryBondResponse](#fury.bonds.v1beta1.QueryBondResponse)
    - [QueryBondsDetailedRequest](#fury.bonds.v1beta1.QueryBondsDetailedRequest)
    - [QueryBondsDetailedResponse](#fury.bonds.v1beta1.QueryBondsDetailedResponse)
    - [QueryBondsRequest](#fury.bonds.v1beta1.QueryBondsRequest)
    - [QueryBondsResponse](#fury.bonds.v1beta1.QueryBondsResponse)
    - [QueryBuyPriceRequest](#fury.bonds.v1beta1.QueryBuyPriceRequest)
    - [QueryBuyPriceResponse](#fury.bonds.v1beta1.QueryBuyPriceResponse)
    - [QueryCurrentPriceRequest](#fury.bonds.v1beta1.QueryCurrentPriceRequest)
    - [QueryCurrentPriceResponse](#fury.bonds.v1beta1.QueryCurrentPriceResponse)
    - [QueryCurrentReserveRequest](#fury.bonds.v1beta1.QueryCurrentReserveRequest)
    - [QueryCurrentReserveResponse](#fury.bonds.v1beta1.QueryCurrentReserveResponse)
    - [QueryCustomPriceRequest](#fury.bonds.v1beta1.QueryCustomPriceRequest)
    - [QueryCustomPriceResponse](#fury.bonds.v1beta1.QueryCustomPriceResponse)
    - [QueryLastBatchRequest](#fury.bonds.v1beta1.QueryLastBatchRequest)
    - [QueryLastBatchResponse](#fury.bonds.v1beta1.QueryLastBatchResponse)
    - [QueryParamsRequest](#fury.bonds.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#fury.bonds.v1beta1.QueryParamsResponse)
    - [QuerySellReturnRequest](#fury.bonds.v1beta1.QuerySellReturnRequest)
    - [QuerySellReturnResponse](#fury.bonds.v1beta1.QuerySellReturnResponse)
    - [QuerySwapReturnRequest](#fury.bonds.v1beta1.QuerySwapReturnRequest)
    - [QuerySwapReturnResponse](#fury.bonds.v1beta1.QuerySwapReturnResponse)
  
    - [Query](#fury.bonds.v1beta1.Query)
  
- [fury/bonds/v1beta1/tx.proto](#fury/bonds/v1beta1/tx.proto)
    - [MsgBuy](#fury.bonds.v1beta1.MsgBuy)
    - [MsgBuyResponse](#fury.bonds.v1beta1.MsgBuyResponse)
    - [MsgCreateBond](#fury.bonds.v1beta1.MsgCreateBond)
    - [MsgCreateBondResponse](#fury.bonds.v1beta1.MsgCreateBondResponse)
    - [MsgEditBond](#fury.bonds.v1beta1.MsgEditBond)
    - [MsgEditBondResponse](#fury.bonds.v1beta1.MsgEditBondResponse)
    - [MsgMakeOutcomePayment](#fury.bonds.v1beta1.MsgMakeOutcomePayment)
    - [MsgMakeOutcomePaymentResponse](#fury.bonds.v1beta1.MsgMakeOutcomePaymentResponse)
    - [MsgSell](#fury.bonds.v1beta1.MsgSell)
    - [MsgSellResponse](#fury.bonds.v1beta1.MsgSellResponse)
    - [MsgSetNextAlpha](#fury.bonds.v1beta1.MsgSetNextAlpha)
    - [MsgSetNextAlphaResponse](#fury.bonds.v1beta1.MsgSetNextAlphaResponse)
    - [MsgSwap](#fury.bonds.v1beta1.MsgSwap)
    - [MsgSwapResponse](#fury.bonds.v1beta1.MsgSwapResponse)
    - [MsgUpdateBondState](#fury.bonds.v1beta1.MsgUpdateBondState)
    - [MsgUpdateBondStateResponse](#fury.bonds.v1beta1.MsgUpdateBondStateResponse)
    - [MsgWithdrawReserve](#fury.bonds.v1beta1.MsgWithdrawReserve)
    - [MsgWithdrawReserveResponse](#fury.bonds.v1beta1.MsgWithdrawReserveResponse)
    - [MsgWithdrawShare](#fury.bonds.v1beta1.MsgWithdrawShare)
    - [MsgWithdrawShareResponse](#fury.bonds.v1beta1.MsgWithdrawShareResponse)
  
    - [Msg](#fury.bonds.v1beta1.Msg)
  
- [fury/claims/v1beta1/claims.proto](#fury/claims/v1beta1/claims.proto)
    - [Claim](#fury.claims.v1beta1.Claim)
    - [ClaimPayments](#fury.claims.v1beta1.ClaimPayments)
    - [Collection](#fury.claims.v1beta1.Collection)
    - [Contract1155Payment](#fury.claims.v1beta1.Contract1155Payment)
    - [Dispute](#fury.claims.v1beta1.Dispute)
    - [DisputeData](#fury.claims.v1beta1.DisputeData)
    - [Evaluation](#fury.claims.v1beta1.Evaluation)
    - [Params](#fury.claims.v1beta1.Params)
    - [Payment](#fury.claims.v1beta1.Payment)
    - [Payments](#fury.claims.v1beta1.Payments)
  
    - [CollectionState](#fury.claims.v1beta1.CollectionState)
    - [EvaluationStatus](#fury.claims.v1beta1.EvaluationStatus)
    - [PaymentStatus](#fury.claims.v1beta1.PaymentStatus)
    - [PaymentType](#fury.claims.v1beta1.PaymentType)
  
- [fury/claims/v1beta1/cosmos.proto](#fury/claims/v1beta1/cosmos.proto)
    - [Input](#fury.claims.v1beta1.Input)
    - [Output](#fury.claims.v1beta1.Output)
  
- [fury/claims/v1beta1/authz.proto](#fury/claims/v1beta1/authz.proto)
    - [EvaluateClaimAuthorization](#fury.claims.v1beta1.EvaluateClaimAuthorization)
    - [EvaluateClaimConstraints](#fury.claims.v1beta1.EvaluateClaimConstraints)
    - [SubmitClaimAuthorization](#fury.claims.v1beta1.SubmitClaimAuthorization)
    - [SubmitClaimConstraints](#fury.claims.v1beta1.SubmitClaimConstraints)
    - [WithdrawPaymentAuthorization](#fury.claims.v1beta1.WithdrawPaymentAuthorization)
    - [WithdrawPaymentConstraints](#fury.claims.v1beta1.WithdrawPaymentConstraints)
  
- [fury/claims/v1beta1/event.proto](#fury/claims/v1beta1/event.proto)
    - [ClaimDisputedEvent](#fury.claims.v1beta1.ClaimDisputedEvent)
    - [ClaimEvaluatedEvent](#fury.claims.v1beta1.ClaimEvaluatedEvent)
    - [ClaimSubmittedEvent](#fury.claims.v1beta1.ClaimSubmittedEvent)
    - [ClaimUpdatedEvent](#fury.claims.v1beta1.ClaimUpdatedEvent)
    - [CollectionCreatedEvent](#fury.claims.v1beta1.CollectionCreatedEvent)
    - [CollectionUpdatedEvent](#fury.claims.v1beta1.CollectionUpdatedEvent)
    - [PaymentWithdrawCreatedEvent](#fury.claims.v1beta1.PaymentWithdrawCreatedEvent)
    - [PaymentWithdrawnEvent](#fury.claims.v1beta1.PaymentWithdrawnEvent)
  
- [fury/claims/v1beta1/genesis.proto](#fury/claims/v1beta1/genesis.proto)
    - [GenesisState](#fury.claims.v1beta1.GenesisState)
  
- [fury/claims/v1beta1/query.proto](#fury/claims/v1beta1/query.proto)
    - [QueryClaimListRequest](#fury.claims.v1beta1.QueryClaimListRequest)
    - [QueryClaimListResponse](#fury.claims.v1beta1.QueryClaimListResponse)
    - [QueryClaimRequest](#fury.claims.v1beta1.QueryClaimRequest)
    - [QueryClaimResponse](#fury.claims.v1beta1.QueryClaimResponse)
    - [QueryCollectionListRequest](#fury.claims.v1beta1.QueryCollectionListRequest)
    - [QueryCollectionListResponse](#fury.claims.v1beta1.QueryCollectionListResponse)
    - [QueryCollectionRequest](#fury.claims.v1beta1.QueryCollectionRequest)
    - [QueryCollectionResponse](#fury.claims.v1beta1.QueryCollectionResponse)
    - [QueryDisputeListRequest](#fury.claims.v1beta1.QueryDisputeListRequest)
    - [QueryDisputeListResponse](#fury.claims.v1beta1.QueryDisputeListResponse)
    - [QueryDisputeRequest](#fury.claims.v1beta1.QueryDisputeRequest)
    - [QueryDisputeResponse](#fury.claims.v1beta1.QueryDisputeResponse)
    - [QueryParamsRequest](#fury.claims.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#fury.claims.v1beta1.QueryParamsResponse)
  
    - [Query](#fury.claims.v1beta1.Query)
  
- [fury/claims/v1beta1/tx.proto](#fury/claims/v1beta1/tx.proto)
    - [MsgCreateCollection](#fury.claims.v1beta1.MsgCreateCollection)
    - [MsgCreateCollectionResponse](#fury.claims.v1beta1.MsgCreateCollectionResponse)
    - [MsgDisputeClaim](#fury.claims.v1beta1.MsgDisputeClaim)
    - [MsgDisputeClaimResponse](#fury.claims.v1beta1.MsgDisputeClaimResponse)
    - [MsgEvaluateClaim](#fury.claims.v1beta1.MsgEvaluateClaim)
    - [MsgEvaluateClaimResponse](#fury.claims.v1beta1.MsgEvaluateClaimResponse)
    - [MsgSubmitClaim](#fury.claims.v1beta1.MsgSubmitClaim)
    - [MsgSubmitClaimResponse](#fury.claims.v1beta1.MsgSubmitClaimResponse)
    - [MsgWithdrawPayment](#fury.claims.v1beta1.MsgWithdrawPayment)
    - [MsgWithdrawPaymentResponse](#fury.claims.v1beta1.MsgWithdrawPaymentResponse)
  
    - [Msg](#fury.claims.v1beta1.Msg)
  
- [fury/entity/v1beta1/cosmos.proto](#fury/entity/v1beta1/cosmos.proto)
    - [Grant](#fury.entity.v1beta1.Grant)
  
- [fury/iid/v1beta1/types.proto](#fury/iid/v1beta1/types.proto)
    - [AccordedRight](#fury.iid.v1beta1.AccordedRight)
    - [Context](#fury.iid.v1beta1.Context)
    - [IidMetadata](#fury.iid.v1beta1.IidMetadata)
    - [LinkedClaim](#fury.iid.v1beta1.LinkedClaim)
    - [LinkedEntity](#fury.iid.v1beta1.LinkedEntity)
    - [LinkedResource](#fury.iid.v1beta1.LinkedResource)
    - [Service](#fury.iid.v1beta1.Service)
    - [VerificationMethod](#fury.iid.v1beta1.VerificationMethod)
  
- [fury/iid/v1beta1/iid.proto](#fury/iid/v1beta1/iid.proto)
    - [IidDocument](#fury.iid.v1beta1.IidDocument)
  
- [fury/entity/v1beta1/entity.proto](#fury/entity/v1beta1/entity.proto)
    - [Entity](#fury.entity.v1beta1.Entity)
    - [EntityAccount](#fury.entity.v1beta1.EntityAccount)
    - [EntityMetadata](#fury.entity.v1beta1.EntityMetadata)
    - [Params](#fury.entity.v1beta1.Params)
  
- [fury/entity/v1beta1/event.proto](#fury/entity/v1beta1/event.proto)
    - [EntityAccountAuthzCreatedEvent](#fury.entity.v1beta1.EntityAccountAuthzCreatedEvent)
    - [EntityAccountCreatedEvent](#fury.entity.v1beta1.EntityAccountCreatedEvent)
    - [EntityCreatedEvent](#fury.entity.v1beta1.EntityCreatedEvent)
    - [EntityTransferredEvent](#fury.entity.v1beta1.EntityTransferredEvent)
    - [EntityUpdatedEvent](#fury.entity.v1beta1.EntityUpdatedEvent)
    - [EntityVerifiedUpdatedEvent](#fury.entity.v1beta1.EntityVerifiedUpdatedEvent)
  
- [fury/entity/v1beta1/genesis.proto](#fury/entity/v1beta1/genesis.proto)
    - [GenesisState](#fury.entity.v1beta1.GenesisState)
  
- [fury/entity/v1beta1/proposal.proto](#fury/entity/v1beta1/proposal.proto)
    - [InitializeNftContract](#fury.entity.v1beta1.InitializeNftContract)
  
- [fury/entity/v1beta1/query.proto](#fury/entity/v1beta1/query.proto)
    - [QueryEntityIidDocumentRequest](#fury.entity.v1beta1.QueryEntityIidDocumentRequest)
    - [QueryEntityIidDocumentResponse](#fury.entity.v1beta1.QueryEntityIidDocumentResponse)
    - [QueryEntityListRequest](#fury.entity.v1beta1.QueryEntityListRequest)
    - [QueryEntityListResponse](#fury.entity.v1beta1.QueryEntityListResponse)
    - [QueryEntityMetadataRequest](#fury.entity.v1beta1.QueryEntityMetadataRequest)
    - [QueryEntityMetadataResponse](#fury.entity.v1beta1.QueryEntityMetadataResponse)
    - [QueryEntityRequest](#fury.entity.v1beta1.QueryEntityRequest)
    - [QueryEntityResponse](#fury.entity.v1beta1.QueryEntityResponse)
    - [QueryEntityVerifiedRequest](#fury.entity.v1beta1.QueryEntityVerifiedRequest)
    - [QueryEntityVerifiedResponse](#fury.entity.v1beta1.QueryEntityVerifiedResponse)
    - [QueryParamsRequest](#fury.entity.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#fury.entity.v1beta1.QueryParamsResponse)
  
    - [Query](#fury.entity.v1beta1.Query)
  
- [fury/iid/v1beta1/tx.proto](#fury/iid/v1beta1/tx.proto)
    - [MsgAddAccordedRight](#fury.iid.v1beta1.MsgAddAccordedRight)
    - [MsgAddAccordedRightResponse](#fury.iid.v1beta1.MsgAddAccordedRightResponse)
    - [MsgAddController](#fury.iid.v1beta1.MsgAddController)
    - [MsgAddControllerResponse](#fury.iid.v1beta1.MsgAddControllerResponse)
    - [MsgAddIidContext](#fury.iid.v1beta1.MsgAddIidContext)
    - [MsgAddIidContextResponse](#fury.iid.v1beta1.MsgAddIidContextResponse)
    - [MsgAddLinkedClaim](#fury.iid.v1beta1.MsgAddLinkedClaim)
    - [MsgAddLinkedClaimResponse](#fury.iid.v1beta1.MsgAddLinkedClaimResponse)
    - [MsgAddLinkedEntity](#fury.iid.v1beta1.MsgAddLinkedEntity)
    - [MsgAddLinkedEntityResponse](#fury.iid.v1beta1.MsgAddLinkedEntityResponse)
    - [MsgAddLinkedResource](#fury.iid.v1beta1.MsgAddLinkedResource)
    - [MsgAddLinkedResourceResponse](#fury.iid.v1beta1.MsgAddLinkedResourceResponse)
    - [MsgAddService](#fury.iid.v1beta1.MsgAddService)
    - [MsgAddServiceResponse](#fury.iid.v1beta1.MsgAddServiceResponse)
    - [MsgAddVerification](#fury.iid.v1beta1.MsgAddVerification)
    - [MsgAddVerificationResponse](#fury.iid.v1beta1.MsgAddVerificationResponse)
    - [MsgCreateIidDocument](#fury.iid.v1beta1.MsgCreateIidDocument)
    - [MsgCreateIidDocumentResponse](#fury.iid.v1beta1.MsgCreateIidDocumentResponse)
    - [MsgDeactivateIID](#fury.iid.v1beta1.MsgDeactivateIID)
    - [MsgDeactivateIIDResponse](#fury.iid.v1beta1.MsgDeactivateIIDResponse)
    - [MsgDeleteAccordedRight](#fury.iid.v1beta1.MsgDeleteAccordedRight)
    - [MsgDeleteAccordedRightResponse](#fury.iid.v1beta1.MsgDeleteAccordedRightResponse)
    - [MsgDeleteController](#fury.iid.v1beta1.MsgDeleteController)
    - [MsgDeleteControllerResponse](#fury.iid.v1beta1.MsgDeleteControllerResponse)
    - [MsgDeleteIidContext](#fury.iid.v1beta1.MsgDeleteIidContext)
    - [MsgDeleteIidContextResponse](#fury.iid.v1beta1.MsgDeleteIidContextResponse)
    - [MsgDeleteLinkedClaim](#fury.iid.v1beta1.MsgDeleteLinkedClaim)
    - [MsgDeleteLinkedClaimResponse](#fury.iid.v1beta1.MsgDeleteLinkedClaimResponse)
    - [MsgDeleteLinkedEntity](#fury.iid.v1beta1.MsgDeleteLinkedEntity)
    - [MsgDeleteLinkedEntityResponse](#fury.iid.v1beta1.MsgDeleteLinkedEntityResponse)
    - [MsgDeleteLinkedResource](#fury.iid.v1beta1.MsgDeleteLinkedResource)
    - [MsgDeleteLinkedResourceResponse](#fury.iid.v1beta1.MsgDeleteLinkedResourceResponse)
    - [MsgDeleteService](#fury.iid.v1beta1.MsgDeleteService)
    - [MsgDeleteServiceResponse](#fury.iid.v1beta1.MsgDeleteServiceResponse)
    - [MsgRevokeVerification](#fury.iid.v1beta1.MsgRevokeVerification)
    - [MsgRevokeVerificationResponse](#fury.iid.v1beta1.MsgRevokeVerificationResponse)
    - [MsgSetVerificationRelationships](#fury.iid.v1beta1.MsgSetVerificationRelationships)
    - [MsgSetVerificationRelationshipsResponse](#fury.iid.v1beta1.MsgSetVerificationRelationshipsResponse)
    - [MsgUpdateIidDocument](#fury.iid.v1beta1.MsgUpdateIidDocument)
    - [MsgUpdateIidDocumentResponse](#fury.iid.v1beta1.MsgUpdateIidDocumentResponse)
    - [Verification](#fury.iid.v1beta1.Verification)
  
    - [Msg](#fury.iid.v1beta1.Msg)
  
- [fury/entity/v1beta1/tx.proto](#fury/entity/v1beta1/tx.proto)
    - [MsgCreateEntity](#fury.entity.v1beta1.MsgCreateEntity)
    - [MsgCreateEntityAccount](#fury.entity.v1beta1.MsgCreateEntityAccount)
    - [MsgCreateEntityAccountResponse](#fury.entity.v1beta1.MsgCreateEntityAccountResponse)
    - [MsgCreateEntityResponse](#fury.entity.v1beta1.MsgCreateEntityResponse)
    - [MsgGrantEntityAccountAuthz](#fury.entity.v1beta1.MsgGrantEntityAccountAuthz)
    - [MsgGrantEntityAccountAuthzResponse](#fury.entity.v1beta1.MsgGrantEntityAccountAuthzResponse)
    - [MsgTransferEntity](#fury.entity.v1beta1.MsgTransferEntity)
    - [MsgTransferEntityResponse](#fury.entity.v1beta1.MsgTransferEntityResponse)
    - [MsgUpdateEntity](#fury.entity.v1beta1.MsgUpdateEntity)
    - [MsgUpdateEntityResponse](#fury.entity.v1beta1.MsgUpdateEntityResponse)
    - [MsgUpdateEntityVerified](#fury.entity.v1beta1.MsgUpdateEntityVerified)
    - [MsgUpdateEntityVerifiedResponse](#fury.entity.v1beta1.MsgUpdateEntityVerifiedResponse)
  
    - [Msg](#fury.entity.v1beta1.Msg)
  
- [fury/iid/v1beta1/event.proto](#fury/iid/v1beta1/event.proto)
    - [IidDocumentCreatedEvent](#fury.iid.v1beta1.IidDocumentCreatedEvent)
    - [IidDocumentUpdatedEvent](#fury.iid.v1beta1.IidDocumentUpdatedEvent)
  
- [fury/iid/v1beta1/genesis.proto](#fury/iid/v1beta1/genesis.proto)
    - [GenesisState](#fury.iid.v1beta1.GenesisState)
  
- [fury/iid/v1beta1/query.proto](#fury/iid/v1beta1/query.proto)
    - [QueryIidDocumentRequest](#fury.iid.v1beta1.QueryIidDocumentRequest)
    - [QueryIidDocumentResponse](#fury.iid.v1beta1.QueryIidDocumentResponse)
    - [QueryIidDocumentsRequest](#fury.iid.v1beta1.QueryIidDocumentsRequest)
    - [QueryIidDocumentsResponse](#fury.iid.v1beta1.QueryIidDocumentsResponse)
  
    - [Query](#fury.iid.v1beta1.Query)
  
- [fury/token/v1beta1/token.proto](#fury/token/v1beta1/token.proto)
    - [Params](#fury.token.v1beta1.Params)
    - [Token](#fury.token.v1beta1.Token)
    - [TokenData](#fury.token.v1beta1.TokenData)
    - [TokenProperties](#fury.token.v1beta1.TokenProperties)
    - [TokensCancelled](#fury.token.v1beta1.TokensCancelled)
    - [TokensRetired](#fury.token.v1beta1.TokensRetired)
  
- [fury/token/v1beta1/authz.proto](#fury/token/v1beta1/authz.proto)
    - [MintAuthorization](#fury.token.v1beta1.MintAuthorization)
    - [MintConstraints](#fury.token.v1beta1.MintConstraints)
  
- [fury/token/v1beta1/tx.proto](#fury/token/v1beta1/tx.proto)
    - [MintBatch](#fury.token.v1beta1.MintBatch)
    - [MsgCancelToken](#fury.token.v1beta1.MsgCancelToken)
    - [MsgCancelTokenResponse](#fury.token.v1beta1.MsgCancelTokenResponse)
    - [MsgCreateToken](#fury.token.v1beta1.MsgCreateToken)
    - [MsgCreateTokenResponse](#fury.token.v1beta1.MsgCreateTokenResponse)
    - [MsgMintToken](#fury.token.v1beta1.MsgMintToken)
    - [MsgMintTokenResponse](#fury.token.v1beta1.MsgMintTokenResponse)
    - [MsgPauseToken](#fury.token.v1beta1.MsgPauseToken)
    - [MsgPauseTokenResponse](#fury.token.v1beta1.MsgPauseTokenResponse)
    - [MsgRetireToken](#fury.token.v1beta1.MsgRetireToken)
    - [MsgRetireTokenResponse](#fury.token.v1beta1.MsgRetireTokenResponse)
    - [MsgStopToken](#fury.token.v1beta1.MsgStopToken)
    - [MsgStopTokenResponse](#fury.token.v1beta1.MsgStopTokenResponse)
    - [MsgTransferToken](#fury.token.v1beta1.MsgTransferToken)
    - [MsgTransferTokenResponse](#fury.token.v1beta1.MsgTransferTokenResponse)
    - [TokenBatch](#fury.token.v1beta1.TokenBatch)
  
    - [Msg](#fury.token.v1beta1.Msg)
  
- [fury/token/v1beta1/event.proto](#fury/token/v1beta1/event.proto)
    - [TokenCancelledEvent](#fury.token.v1beta1.TokenCancelledEvent)
    - [TokenCreatedEvent](#fury.token.v1beta1.TokenCreatedEvent)
    - [TokenMintedEvent](#fury.token.v1beta1.TokenMintedEvent)
    - [TokenPausedEvent](#fury.token.v1beta1.TokenPausedEvent)
    - [TokenRetiredEvent](#fury.token.v1beta1.TokenRetiredEvent)
    - [TokenStoppedEvent](#fury.token.v1beta1.TokenStoppedEvent)
    - [TokenTransferredEvent](#fury.token.v1beta1.TokenTransferredEvent)
    - [TokenUpdatedEvent](#fury.token.v1beta1.TokenUpdatedEvent)
  
- [fury/token/v1beta1/genesis.proto](#fury/token/v1beta1/genesis.proto)
    - [GenesisState](#fury.token.v1beta1.GenesisState)
  
- [fury/token/v1beta1/proposal.proto](#fury/token/v1beta1/proposal.proto)
    - [SetTokenContractCodes](#fury.token.v1beta1.SetTokenContractCodes)
  
- [fury/token/v1beta1/query.proto](#fury/token/v1beta1/query.proto)
    - [QueryParamsRequest](#fury.token.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#fury.token.v1beta1.QueryParamsResponse)
    - [QueryTokenDocRequest](#fury.token.v1beta1.QueryTokenDocRequest)
    - [QueryTokenDocResponse](#fury.token.v1beta1.QueryTokenDocResponse)
    - [QueryTokenListRequest](#fury.token.v1beta1.QueryTokenListRequest)
    - [QueryTokenListResponse](#fury.token.v1beta1.QueryTokenListResponse)
    - [QueryTokenMetadataRequest](#fury.token.v1beta1.QueryTokenMetadataRequest)
    - [QueryTokenMetadataResponse](#fury.token.v1beta1.QueryTokenMetadataResponse)
    - [TokenMetadataProperties](#fury.token.v1beta1.TokenMetadataProperties)
  
    - [Query](#fury.token.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="fury/bonds/v1beta1/bonds.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/bonds/v1beta1/bonds.proto



<a name="fury.bonds.v1beta1.BaseOrder"></a>

### BaseOrder
BaseOrder defines a base order type. It contains all the necessary fields for
specifying the general details about a buy, sell, or swap order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| cancelled | [bool](#bool) |  |  |
| cancel_reason | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.Batch"></a>

### Batch
Batch holds a collection of outstanding buy, sell, and swap orders on a
particular bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| blocks_remaining | [string](#string) |  |  |
| next_public_alpha | [string](#string) |  |  |
| total_buy_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| total_sell_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| buy_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| sell_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| buys | [BuyOrder](#fury.bonds.v1beta1.BuyOrder) | repeated |  |
| sells | [SellOrder](#fury.bonds.v1beta1.SellOrder) | repeated |  |
| swaps | [SwapOrder](#fury.bonds.v1beta1.SwapOrder) | repeated |  |






<a name="fury.bonds.v1beta1.Bond"></a>

### Bond
Bond defines a token bonding curve type with all of its parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#fury.bonds.v1beta1.FunctionParam) | repeated |  |
| reserve_tokens | [string](#string) | repeated |  |
| tx_fee_percentage | [string](#string) |  |  |
| exit_fee_percentage | [string](#string) |  |  |
| fee_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| current_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| current_outcome_payment_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| allow_sells | [bool](#bool) |  |  |
| allow_reserve_withdrawals | [bool](#bool) |  |  |
| alpha_bond | [bool](#bool) |  |  |
| batch_blocks | [string](#string) |  |  |
| outcome_payment | [string](#string) |  |  |
| state | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.BondDetails"></a>

### BondDetails
BondDetails contains details about the current state of a given bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| spot_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.bonds.v1beta1.BuyOrder"></a>

### BuyOrder
BuyOrder defines a type for submitting a buy order on a bond, together with
the maximum amount of reserve tokens the buyer is willing to pay.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#fury.bonds.v1beta1.BaseOrder) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.bonds.v1beta1.FunctionParam"></a>

### FunctionParam
FunctionParam is a key-value pair used for specifying a specific bond
parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| param | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.Params"></a>

### Params
Params defines the parameters for the bonds module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reserved_bond_tokens | [string](#string) | repeated |  |






<a name="fury.bonds.v1beta1.SellOrder"></a>

### SellOrder
SellOrder defines a type for submitting a sell order on a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#fury.bonds.v1beta1.BaseOrder) |  |  |






<a name="fury.bonds.v1beta1.SwapOrder"></a>

### SwapOrder
SwapOrder defines a type for submitting a swap order between two tokens on a
bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#fury.bonds.v1beta1.BaseOrder) |  |  |
| to_token | [string](#string) |  |  |





 

 

 

 



<a name="fury/bonds/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/bonds/v1beta1/genesis.proto



<a name="fury.bonds.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the bonds module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [Bond](#fury.bonds.v1beta1.Bond) | repeated |  |
| batches | [Batch](#fury.bonds.v1beta1.Batch) | repeated |  |
| params | [Params](#fury.bonds.v1beta1.Params) |  |  |





 

 

 

 



<a name="fury/bonds/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/bonds/v1beta1/query.proto



<a name="fury.bonds.v1beta1.QueryAlphaMaximumsRequest"></a>

### QueryAlphaMaximumsRequest
QueryAlphaMaximumsRequest is the request type for the Query/AlphaMaximums RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryAlphaMaximumsResponse"></a>

### QueryAlphaMaximumsResponse
QueryAlphaMaximumsResponse is the response type for the Query/AlphaMaximums
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_system_alpha_increase | [string](#string) |  |  |
| max_system_alpha | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryAvailableReserveRequest"></a>

### QueryAvailableReserveRequest
QueryAvailableReserveRequest is the request type for the
Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryAvailableReserveResponse"></a>

### QueryAvailableReserveResponse
QueryAvailableReserveResponse is the response type for the
Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.bonds.v1beta1.QueryBatchRequest"></a>

### QueryBatchRequest
QueryBatchRequest is the request type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryBatchResponse"></a>

### QueryBatchResponse
QueryBatchResponse is the response type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch | [Batch](#fury.bonds.v1beta1.Batch) |  |  |






<a name="fury.bonds.v1beta1.QueryBondRequest"></a>

### QueryBondRequest
QueryBondRequest is the request type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryBondResponse"></a>

### QueryBondResponse
QueryBondResponse is the response type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#fury.bonds.v1beta1.Bond) |  |  |






<a name="fury.bonds.v1beta1.QueryBondsDetailedRequest"></a>

### QueryBondsDetailedRequest
QueryBondsDetailedRequest is the request type for the Query/BondsDetailed RPC
method.






<a name="fury.bonds.v1beta1.QueryBondsDetailedResponse"></a>

### QueryBondsDetailedResponse
QueryBondsDetailedResponse is the response type for the Query/BondsDetailed
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds_detailed | [BondDetails](#fury.bonds.v1beta1.BondDetails) | repeated |  |






<a name="fury.bonds.v1beta1.QueryBondsRequest"></a>

### QueryBondsRequest
QueryBondsRequest is the request type for the Query/Bonds RPC method.






<a name="fury.bonds.v1beta1.QueryBondsResponse"></a>

### QueryBondsResponse
QueryBondsResponse is the response type for the Query/Bonds RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [string](#string) | repeated |  |






<a name="fury.bonds.v1beta1.QueryBuyPriceRequest"></a>

### QueryBuyPriceRequest
QueryCustomPriceRequest is the request type for the Query/BuyPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryBuyPriceResponse"></a>

### QueryBuyPriceResponse
QueryCustomPriceResponse is the response type for the Query/BuyPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.bonds.v1beta1.QueryCurrentPriceRequest"></a>

### QueryCurrentPriceRequest
QueryCurrentPriceRequest is the request type for the Query/CurrentPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryCurrentPriceResponse"></a>

### QueryCurrentPriceResponse
QueryCurrentPriceResponse is the response type for the Query/CurrentPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="fury.bonds.v1beta1.QueryCurrentReserveRequest"></a>

### QueryCurrentReserveRequest
QueryCurrentReserveRequest is the request type for the Query/CurrentReserve
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryCurrentReserveResponse"></a>

### QueryCurrentReserveResponse
QueryCurrentReserveResponse is the response type for the Query/CurrentReserve
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.bonds.v1beta1.QueryCustomPriceRequest"></a>

### QueryCustomPriceRequest
QueryCustomPriceRequest is the request type for the Query/CustomPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryCustomPriceResponse"></a>

### QueryCustomPriceResponse
QueryCustomPriceResponse is the response type for the Query/CustomPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="fury.bonds.v1beta1.QueryLastBatchRequest"></a>

### QueryLastBatchRequest
QueryLastBatchRequest is the request type for the Query/LastBatch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QueryLastBatchResponse"></a>

### QueryLastBatchResponse
QueryLastBatchResponse is the response type for the Query/LastBatch RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_batch | [Batch](#fury.bonds.v1beta1.Batch) |  |  |






<a name="fury.bonds.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="fury.bonds.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#fury.bonds.v1beta1.Params) |  |  |






<a name="fury.bonds.v1beta1.QuerySellReturnRequest"></a>

### QuerySellReturnRequest
QuerySellReturnRequest is the request type for the Query/SellReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QuerySellReturnResponse"></a>

### QuerySellReturnResponse
QuerySellReturnResponse is the response type for the Query/SellReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| exit_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.bonds.v1beta1.QuerySwapReturnRequest"></a>

### QuerySwapReturnRequest
QuerySwapReturnRequest is the request type for the Query/SwapReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| from_token_with_amount | [string](#string) |  |  |
| to_token | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.QuerySwapReturnResponse"></a>

### QuerySwapReturnResponse
QuerySwapReturnResponse is the response type for the Query/SwapReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 


<a name="fury.bonds.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bonds | [QueryBondsRequest](#fury.bonds.v1beta1.QueryBondsRequest) | [QueryBondsResponse](#fury.bonds.v1beta1.QueryBondsResponse) | Bonds returns all existing bonds. |
| BondsDetailed | [QueryBondsDetailedRequest](#fury.bonds.v1beta1.QueryBondsDetailedRequest) | [QueryBondsDetailedResponse](#fury.bonds.v1beta1.QueryBondsDetailedResponse) | BondsDetailed returns a list of all existing bonds with some details about their current state. |
| Params | [QueryParamsRequest](#fury.bonds.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#fury.bonds.v1beta1.QueryParamsResponse) | Params queries the paramaters of x/bonds module. |
| Bond | [QueryBondRequest](#fury.bonds.v1beta1.QueryBondRequest) | [QueryBondResponse](#fury.bonds.v1beta1.QueryBondResponse) | Bond queries info of a specific bond. |
| Batch | [QueryBatchRequest](#fury.bonds.v1beta1.QueryBatchRequest) | [QueryBatchResponse](#fury.bonds.v1beta1.QueryBatchResponse) | Batch queries info of a specific bond&#39;s current batch. |
| LastBatch | [QueryLastBatchRequest](#fury.bonds.v1beta1.QueryLastBatchRequest) | [QueryLastBatchResponse](#fury.bonds.v1beta1.QueryLastBatchResponse) | LastBatch queries info of a specific bond&#39;s last batch. |
| CurrentPrice | [QueryCurrentPriceRequest](#fury.bonds.v1beta1.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#fury.bonds.v1beta1.QueryCurrentPriceResponse) | CurrentPrice queries the current price/s of a specific bond. |
| CurrentReserve | [QueryCurrentReserveRequest](#fury.bonds.v1beta1.QueryCurrentReserveRequest) | [QueryCurrentReserveResponse](#fury.bonds.v1beta1.QueryCurrentReserveResponse) | CurrentReserve queries the current balance/s of the reserve pool for a specific bond. |
| AvailableReserve | [QueryAvailableReserveRequest](#fury.bonds.v1beta1.QueryAvailableReserveRequest) | [QueryAvailableReserveResponse](#fury.bonds.v1beta1.QueryAvailableReserveResponse) | AvailableReserve queries current available balance/s of the reserve pool for a specific bond. |
| CustomPrice | [QueryCustomPriceRequest](#fury.bonds.v1beta1.QueryCustomPriceRequest) | [QueryCustomPriceResponse](#fury.bonds.v1beta1.QueryCustomPriceResponse) | CustomPrice queries price/s of a specific bond at a specific supply. |
| BuyPrice | [QueryBuyPriceRequest](#fury.bonds.v1beta1.QueryBuyPriceRequest) | [QueryBuyPriceResponse](#fury.bonds.v1beta1.QueryBuyPriceResponse) | BuyPrice queries price/s of buying an amount of tokens from a specific bond. |
| SellReturn | [QuerySellReturnRequest](#fury.bonds.v1beta1.QuerySellReturnRequest) | [QuerySellReturnResponse](#fury.bonds.v1beta1.QuerySellReturnResponse) | SellReturn queries return/s on selling an amount of tokens of a specific bond. |
| SwapReturn | [QuerySwapReturnRequest](#fury.bonds.v1beta1.QuerySwapReturnRequest) | [QuerySwapReturnResponse](#fury.bonds.v1beta1.QuerySwapReturnResponse) | SwapReturn queries return/s on swapping an amount of tokens to another token of a specific bond. |
| AlphaMaximums | [QueryAlphaMaximumsRequest](#fury.bonds.v1beta1.QueryAlphaMaximumsRequest) | [QueryAlphaMaximumsResponse](#fury.bonds.v1beta1.QueryAlphaMaximumsResponse) | AlphaMaximums queries alpha maximums for a specific augmented bonding curve. |

 



<a name="fury/bonds/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/bonds/v1beta1/tx.proto



<a name="fury.bonds.v1beta1.MsgBuy"></a>

### MsgBuy
MsgBuy defines a message for buying from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |
| buyer_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgBuyResponse"></a>

### MsgBuyResponse
MsgBuyResponse defines the Msg/Buy response type.






<a name="fury.bonds.v1beta1.MsgCreateBond"></a>

### MsgCreateBond
MsgCreateBond defines a message for creating a new bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#fury.bonds.v1beta1.FunctionParam) | repeated |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |
| reserve_tokens | [string](#string) | repeated |  |
| tx_fee_percentage | [string](#string) |  |  |
| exit_fee_percentage | [string](#string) |  |  |
| fee_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| allow_sells | [bool](#bool) |  |  |
| allow_reserve_withdrawals | [bool](#bool) |  |  |
| alpha_bond | [bool](#bool) |  |  |
| batch_blocks | [string](#string) |  |  |
| outcome_payment | [string](#string) |  |  |
| creator_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgCreateBondResponse"></a>

### MsgCreateBondResponse
MsgCreateBondResponse defines the Msg/CreateBond response type.






<a name="fury.bonds.v1beta1.MsgEditBond"></a>

### MsgEditBond
MsgEditBond defines a message for editing an existing bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| order_quantity_limits | [string](#string) |  |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |
| editor_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgEditBondResponse"></a>

### MsgEditBondResponse
MsgEditBondResponse defines the Msg/EditBond response type.






<a name="fury.bonds.v1beta1.MsgMakeOutcomePayment"></a>

### MsgMakeOutcomePayment
MsgMakeOutcomePayment defines a message for making an outcome payment to a
bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| sender_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgMakeOutcomePaymentResponse"></a>

### MsgMakeOutcomePaymentResponse
MsgMakeOutcomePaymentResponse defines the Msg/MakeOutcomePayment response
type.






<a name="fury.bonds.v1beta1.MsgSell"></a>

### MsgSell
MsgSell defines a message for selling from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seller_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| bond_did | [string](#string) |  |  |
| seller_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgSellResponse"></a>

### MsgSellResponse
MsgSellResponse defines the Msg/Sell response type.






<a name="fury.bonds.v1beta1.MsgSetNextAlpha"></a>

### MsgSetNextAlpha
MsgSetNextAlpha defines a message for editing a bond&#39;s alpha parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| alpha | [string](#string) |  |  |
| delta | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |
| oracle_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgSetNextAlphaResponse"></a>

### MsgSetNextAlphaResponse







<a name="fury.bonds.v1beta1.MsgSwap"></a>

### MsgSwap
MsgSwap defines a message for swapping from one reserve bond token to
another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| swapper_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| from | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| to_token | [string](#string) |  |  |
| swapper_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgSwapResponse"></a>

### MsgSwapResponse
MsgSwapResponse defines the Msg/Swap response type.






<a name="fury.bonds.v1beta1.MsgUpdateBondState"></a>

### MsgUpdateBondState
MsgUpdateBondState defines a message for updating a bond&#39;s current state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| state | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |
| editor_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgUpdateBondStateResponse"></a>

### MsgUpdateBondStateResponse
MsgUpdateBondStateResponse defines the Msg/UpdateBondState response type.






<a name="fury.bonds.v1beta1.MsgWithdrawReserve"></a>

### MsgWithdrawReserve
MsgWithdrawReserve defines a message for withdrawing reserve from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdrawer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |
| withdrawer_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgWithdrawReserveResponse"></a>

### MsgWithdrawReserveResponse
MsgWithdrawReserveResponse defines the Msg/WithdrawReserve response type.






<a name="fury.bonds.v1beta1.MsgWithdrawShare"></a>

### MsgWithdrawShare
MsgWithdrawShare defines a message for withdrawing a share from a bond that
is in the SETTLE stage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| recipient_address | [string](#string) |  |  |






<a name="fury.bonds.v1beta1.MsgWithdrawShareResponse"></a>

### MsgWithdrawShareResponse
MsgWithdrawShareResponse defines the Msg/WithdrawShare response type.





 

 

 


<a name="fury.bonds.v1beta1.Msg"></a>

### Msg
Msg defines the bonds Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBond | [MsgCreateBond](#fury.bonds.v1beta1.MsgCreateBond) | [MsgCreateBondResponse](#fury.bonds.v1beta1.MsgCreateBondResponse) | CreateBond defines a method for creating a bond. |
| EditBond | [MsgEditBond](#fury.bonds.v1beta1.MsgEditBond) | [MsgEditBondResponse](#fury.bonds.v1beta1.MsgEditBondResponse) | EditBond defines a method for editing a bond. |
| SetNextAlpha | [MsgSetNextAlpha](#fury.bonds.v1beta1.MsgSetNextAlpha) | [MsgSetNextAlphaResponse](#fury.bonds.v1beta1.MsgSetNextAlphaResponse) | SetNextAlpha defines a method for editing a bond&#39;s alpha parameter. |
| UpdateBondState | [MsgUpdateBondState](#fury.bonds.v1beta1.MsgUpdateBondState) | [MsgUpdateBondStateResponse](#fury.bonds.v1beta1.MsgUpdateBondStateResponse) | UpdateBondState defines a method for updating a bond&#39;s current state. |
| Buy | [MsgBuy](#fury.bonds.v1beta1.MsgBuy) | [MsgBuyResponse](#fury.bonds.v1beta1.MsgBuyResponse) | Buy defines a method for buying from a bond. |
| Sell | [MsgSell](#fury.bonds.v1beta1.MsgSell) | [MsgSellResponse](#fury.bonds.v1beta1.MsgSellResponse) | Sell defines a method for selling from a bond. |
| Swap | [MsgSwap](#fury.bonds.v1beta1.MsgSwap) | [MsgSwapResponse](#fury.bonds.v1beta1.MsgSwapResponse) | Swap defines a method for swapping from one reserve bond token to another. |
| MakeOutcomePayment | [MsgMakeOutcomePayment](#fury.bonds.v1beta1.MsgMakeOutcomePayment) | [MsgMakeOutcomePaymentResponse](#fury.bonds.v1beta1.MsgMakeOutcomePaymentResponse) | MakeOutcomePayment defines a method for making an outcome payment to a bond. |
| WithdrawShare | [MsgWithdrawShare](#fury.bonds.v1beta1.MsgWithdrawShare) | [MsgWithdrawShareResponse](#fury.bonds.v1beta1.MsgWithdrawShareResponse) | WithdrawShare defines a method for withdrawing a share from a bond that is in the SETTLE stage. |
| WithdrawReserve | [MsgWithdrawReserve](#fury.bonds.v1beta1.MsgWithdrawReserve) | [MsgWithdrawReserveResponse](#fury.bonds.v1beta1.MsgWithdrawReserveResponse) | WithdrawReserve defines a method for withdrawing reserve from a bond. |

 



<a name="fury/claims/v1beta1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/claims.proto



<a name="fury.claims.v1beta1.Claim"></a>

### Claim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| agent_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| agent_address | [string](#string) |  |  |
| submission_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submissionDate is the date and time that the claim was submitted on-chain |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim in the cid hash format |
| evaluation | [Evaluation](#fury.claims.v1beta1.Evaluation) |  | evaluation is the result of one or more claim evaluations |
| payments_status | [ClaimPayments](#fury.claims.v1beta1.ClaimPayments) |  |  |






<a name="fury.claims.v1beta1.ClaimPayments"></a>

### ClaimPayments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission | [PaymentStatus](#fury.claims.v1beta1.PaymentStatus) |  |  |
| evaluation | [PaymentStatus](#fury.claims.v1beta1.PaymentStatus) |  |  |
| approval | [PaymentStatus](#fury.claims.v1beta1.PaymentStatus) |  |  |
| rejection | [PaymentStatus](#fury.claims.v1beta1.PaymentStatus) |  | PaymentStatus penalty = 5; |






<a name="fury.claims.v1beta1.Collection"></a>

### Collection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | collection id is the incremented internal id for the collection of claims |
| entity | [string](#string) |  | entity is the DID of the entity for which the claims are being created |
| admin | [string](#string) |  | admin is the account address that will authorize or revoke agents and payments (the grantor) |
| protocol | [string](#string) |  | protocol is the DID of the claim protocol |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| quota | [uint64](#uint64) |  | quota is the maximum number of claims that may be submitted, 0 is unlimited |
| count | [uint64](#uint64) |  | count is the number of claims already submitted (internally calculated) |
| evaluated | [uint64](#uint64) |  | evaluated is the number of claims that have been evaluated (internally calculated) |
| approved | [uint64](#uint64) |  | approved is the number of claims that have been evaluated and approved (internally calculated) |
| rejected | [uint64](#uint64) |  | rejected is the number of claims that have been evaluated and rejected (internally calculated) |
| disputed | [uint64](#uint64) |  | disputed is the number of claims that have disputed status (internally calculated) |
| state | [CollectionState](#fury.claims.v1beta1.CollectionState) |  | state is the current state of this Collection (open, paused, closed) |
| payments | [Payments](#fury.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |
| signer | [string](#string) |  | signer address |






<a name="fury.claims.v1beta1.Contract1155Payment"></a>

### Contract1155Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| token_id | [string](#string) |  |  |
| amount | [uint32](#uint32) |  |  |






<a name="fury.claims.v1beta1.Dispute"></a>

### Dispute



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subject_id | [string](#string) |  |  |
| type | [int32](#int32) |  | type is expressed as an integer, interpreted by the client |
| data | [DisputeData](#fury.claims.v1beta1.DisputeData) |  |  |






<a name="fury.claims.v1beta1.DisputeData"></a>

### DisputeData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | dispute link ***.ipfs |
| type | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [bool](#bool) |  |  |






<a name="fury.claims.v1beta1.Evaluation"></a>

### Evaluation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id indicates which Claim this evaluation is for |
| collection_id | [string](#string) |  | collection_id indicates to which Collection the claim being evaluated belongs to |
| oracle | [string](#string) |  | oracle is the DID of the Oracle entity that evaluates the claim |
| agent_did | [string](#string) |  | agent is the DID of the agent that submits the evaluation |
| agent_address | [string](#string) |  |  |
| status | [EvaluationStatus](#fury.claims.v1beta1.EvaluationStatus) |  | status is the evaluation status expressed as an integer (2=approved, 3=rejected, ...) |
| reason | [uint32](#uint32) |  | reason is the code expressed as an integer, for why the evaluation result was given (codes defined by evaluator) |
| verification_proof | [string](#string) |  | verificationProof is the cid of the evaluation Verfiable Credential |
| evaluation_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | evaluationDate is the date and time that the claim evaluation was submitted on-chain |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by evaluator for claim approval, if empty list then use default by Collection |






<a name="fury.claims.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_sequence | [uint64](#uint64) |  |  |
| fury_account | [string](#string) |  |  |
| network_fee_percentage | [string](#string) |  |  |
| node_fee_percentage | [string](#string) |  |  |






<a name="fury.claims.v1beta1.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account is the entity account address from which the payment will be made |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| contract_1155_payment | [Contract1155Payment](#fury.claims.v1beta1.Contract1155Payment) |  | if empty(nil) then no contract payment, not allowed for Evaluation Payment |
| timeout_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  | timeout after claim/evaluation to create authZ for payment, if 0 then immidiate direct payment |






<a name="fury.claims.v1beta1.Payments"></a>

### Payments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission | [Payment](#fury.claims.v1beta1.Payment) |  |  |
| evaluation | [Payment](#fury.claims.v1beta1.Payment) |  |  |
| approval | [Payment](#fury.claims.v1beta1.Payment) |  |  |
| rejection | [Payment](#fury.claims.v1beta1.Payment) |  | Payment penalty = 5; |





 


<a name="fury.claims.v1beta1.CollectionState"></a>

### CollectionState


| Name | Number | Description |
| ---- | ------ | ----------- |
| OPEN | 0 |  |
| PAUSED | 1 |  |
| CLOSED | 2 |  |



<a name="fury.claims.v1beta1.EvaluationStatus"></a>

### EvaluationStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 |  |
| APPROVED | 1 |  |
| REJECTED | 2 |  |
| DISPUTED | 3 |  |



<a name="fury.claims.v1beta1.PaymentStatus"></a>

### PaymentStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| NO_PAYMENT | 0 |  |
| PROMISED | 1 | agent is contracted to receive payment |
| AUTHORIZED | 2 | authz set up, no guarantee |
| GAURANTEED | 3 | escrow set up with funds blocked |
| PAID | 4 |  |
| FAILED | 5 |  |
| DISPUTED | 6 |  |



<a name="fury.claims.v1beta1.PaymentType"></a>

### PaymentType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBMISSION | 0 |  |
| APPROVAL | 1 |  |
| EVALUATION | 2 |  |
| REJECTION | 3 |  |


 

 

 



<a name="fury/claims/v1beta1/cosmos.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/cosmos.proto



<a name="fury.claims.v1beta1.Input"></a>

### Input
Input models transaction input.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| coins | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="fury.claims.v1beta1.Output"></a>

### Output
Output models transaction outputs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| coins | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 

 



<a name="fury/claims/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/authz.proto



<a name="fury.claims.v1beta1.EvaluateClaimAuthorization"></a>

### EvaluateClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [EvaluateClaimConstraints](#fury.claims.v1beta1.EvaluateClaimConstraints) | repeated |  |






<a name="fury.claims.v1beta1.EvaluateClaimConstraints"></a>

### EvaluateClaimConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| claim_ids | [string](#string) | repeated | either collection_id or claim_ids is needed |
| agent_quota | [uint64](#uint64) |  |  |
| before_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | if null then no before_date validation done |
| max_custom_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | max custom amount evaluator can change, if empty list must use amount defined in Token payments |






<a name="fury.claims.v1beta1.SubmitClaimAuthorization"></a>

### SubmitClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [SubmitClaimConstraints](#fury.claims.v1beta1.SubmitClaimConstraints) | repeated |  |






<a name="fury.claims.v1beta1.SubmitClaimConstraints"></a>

### SubmitClaimConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| agent_quota | [uint64](#uint64) |  |  |






<a name="fury.claims.v1beta1.WithdrawPaymentAuthorization"></a>

### WithdrawPaymentAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [WithdrawPaymentConstraints](#fury.claims.v1beta1.WithdrawPaymentConstraints) | repeated |  |






<a name="fury.claims.v1beta1.WithdrawPaymentConstraints"></a>

### WithdrawPaymentConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id the withdrawal is for |
| inputs | [Input](#fury.claims.v1beta1.Input) | repeated | Inputs to the multisend tx to run to withdraw payment |
| outputs | [Output](#fury.claims.v1beta1.Output) | repeated | Outputs for the multisend tx to run to withdraw payment |
| payment_type | [PaymentType](#fury.claims.v1beta1.PaymentType) |  | payment type to keep track what payment is for and mark claim payment accordingly |
| contract_1155_payment | [Contract1155Payment](#fury.claims.v1beta1.Contract1155Payment) |  | if empty(nil) then no contract payment |
| toAddress | [string](#string) |  | for contract payment |
| fromAddress | [string](#string) |  | for contract payment |
| release_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date that grantee can execute authorization, calculated from created date plus the timeout on Collection payments, if null then none |





 

 

 

 



<a name="fury/claims/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/event.proto



<a name="fury.claims.v1beta1.ClaimDisputedEvent"></a>

### ClaimDisputedEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dispute | [Dispute](#fury.claims.v1beta1.Dispute) |  |  |






<a name="fury.claims.v1beta1.ClaimEvaluatedEvent"></a>

### ClaimEvaluatedEvent
ClaimEvaluatedEvent is an event triggered on a Claim evaluation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| evaluation | [Evaluation](#fury.claims.v1beta1.Evaluation) |  |  |






<a name="fury.claims.v1beta1.ClaimSubmittedEvent"></a>

### ClaimSubmittedEvent
CollectionCreatedEvent is an event triggered on a Claim submission


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#fury.claims.v1beta1.Claim) |  |  |






<a name="fury.claims.v1beta1.ClaimUpdatedEvent"></a>

### ClaimUpdatedEvent
ClaimUpdatedEvent is an event triggered on a Claim update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#fury.claims.v1beta1.Claim) |  |  |






<a name="fury.claims.v1beta1.CollectionCreatedEvent"></a>

### CollectionCreatedEvent
CollectionCreatedEvent is an event triggered on a Collection creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#fury.claims.v1beta1.Collection) |  |  |






<a name="fury.claims.v1beta1.CollectionUpdatedEvent"></a>

### CollectionUpdatedEvent
CollectionUpdatedEvent is an event triggered on a Collection update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#fury.claims.v1beta1.Collection) |  |  |






<a name="fury.claims.v1beta1.PaymentWithdrawCreatedEvent"></a>

### PaymentWithdrawCreatedEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw | [WithdrawPaymentConstraints](#fury.claims.v1beta1.WithdrawPaymentConstraints) |  |  |






<a name="fury.claims.v1beta1.PaymentWithdrawnEvent"></a>

### PaymentWithdrawnEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw | [WithdrawPaymentConstraints](#fury.claims.v1beta1.WithdrawPaymentConstraints) |  |  |





 

 

 

 



<a name="fury/claims/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/genesis.proto



<a name="fury.claims.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the claims module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#fury.claims.v1beta1.Params) |  |  |
| collections | [Collection](#fury.claims.v1beta1.Collection) | repeated |  |
| claims | [Claim](#fury.claims.v1beta1.Claim) | repeated |  |
| disputes | [Dispute](#fury.claims.v1beta1.Dispute) | repeated |  |





 

 

 

 



<a name="fury/claims/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/query.proto



<a name="fury.claims.v1beta1.QueryClaimListRequest"></a>

### QueryClaimListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="fury.claims.v1beta1.QueryClaimListResponse"></a>

### QueryClaimListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claims | [Claim](#fury.claims.v1beta1.Claim) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="fury.claims.v1beta1.QueryClaimRequest"></a>

### QueryClaimRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.claims.v1beta1.QueryClaimResponse"></a>

### QueryClaimResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#fury.claims.v1beta1.Claim) |  |  |






<a name="fury.claims.v1beta1.QueryCollectionListRequest"></a>

### QueryCollectionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="fury.claims.v1beta1.QueryCollectionListResponse"></a>

### QueryCollectionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [Collection](#fury.claims.v1beta1.Collection) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="fury.claims.v1beta1.QueryCollectionRequest"></a>

### QueryCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.claims.v1beta1.QueryCollectionResponse"></a>

### QueryCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#fury.claims.v1beta1.Collection) |  |  |






<a name="fury.claims.v1beta1.QueryDisputeListRequest"></a>

### QueryDisputeListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="fury.claims.v1beta1.QueryDisputeListResponse"></a>

### QueryDisputeListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disputes | [Dispute](#fury.claims.v1beta1.Dispute) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="fury.claims.v1beta1.QueryDisputeRequest"></a>

### QueryDisputeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proof | [string](#string) |  |  |






<a name="fury.claims.v1beta1.QueryDisputeResponse"></a>

### QueryDisputeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dispute | [Dispute](#fury.claims.v1beta1.Dispute) |  |  |






<a name="fury.claims.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="fury.claims.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#fury.claims.v1beta1.Params) |  | params holds all the parameters of this module. |





 

 

 


<a name="fury.claims.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#fury.claims.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#fury.claims.v1beta1.QueryParamsResponse) | Parameters queries the parameters of the module. |
| Collection | [QueryCollectionRequest](#fury.claims.v1beta1.QueryCollectionRequest) | [QueryCollectionResponse](#fury.claims.v1beta1.QueryCollectionResponse) |  |
| CollectionList | [QueryCollectionListRequest](#fury.claims.v1beta1.QueryCollectionListRequest) | [QueryCollectionListResponse](#fury.claims.v1beta1.QueryCollectionListResponse) |  |
| Claim | [QueryClaimRequest](#fury.claims.v1beta1.QueryClaimRequest) | [QueryClaimResponse](#fury.claims.v1beta1.QueryClaimResponse) |  |
| ClaimList | [QueryClaimListRequest](#fury.claims.v1beta1.QueryClaimListRequest) | [QueryClaimListResponse](#fury.claims.v1beta1.QueryClaimListResponse) |  |
| Dispute | [QueryDisputeRequest](#fury.claims.v1beta1.QueryDisputeRequest) | [QueryDisputeResponse](#fury.claims.v1beta1.QueryDisputeResponse) |  |
| DisputeList | [QueryDisputeListRequest](#fury.claims.v1beta1.QueryDisputeListRequest) | [QueryDisputeListResponse](#fury.claims.v1beta1.QueryDisputeListResponse) |  |

 



<a name="fury/claims/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/claims/v1beta1/tx.proto



<a name="fury.claims.v1beta1.MsgCreateCollection"></a>

### MsgCreateCollection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [string](#string) |  | entity is the DID of the entity for which the claims are being created |
| signer | [string](#string) |  | signer address |
| protocol | [string](#string) |  | protocol is the DID of the claim protocol |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| quota | [uint64](#uint64) |  | quota is the maximum number of claims that may be submitted, 0 is unlimited |
| state | [CollectionState](#fury.claims.v1beta1.CollectionState) |  | state is the current state of this Collection (open, paused, closed) |
| payments | [Payments](#fury.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |






<a name="fury.claims.v1beta1.MsgCreateCollectionResponse"></a>

### MsgCreateCollectionResponse







<a name="fury.claims.v1beta1.MsgDisputeClaim"></a>

### MsgDisputeClaim
Agent laying dispute must be admin for Collection, or controller on
Collection entity, or have authz cap, aka is agent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subject_id | [string](#string) |  | subject_id for which this dispute is against, for now can only lay disputes against claims |
| agent_did | [string](#string) |  | agent is the DID of the agent disputing the claim, agent detials wont be saved in kvStore |
| agent_address | [string](#string) |  |  |
| dispute_type | [int32](#int32) |  | type is expressed as an integer, interpreted by the client |
| data | [DisputeData](#fury.claims.v1beta1.DisputeData) |  |  |






<a name="fury.claims.v1beta1.MsgDisputeClaimResponse"></a>

### MsgDisputeClaimResponse







<a name="fury.claims.v1beta1.MsgEvaluateClaim"></a>

### MsgEvaluateClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim to make evaluation against |
| collection_id | [string](#string) |  | claimID is the unique identifier of the claim to make evaluation against |
| oracle | [string](#string) |  | oracle is the DID of the Oracle entity that evaluates the claim |
| agent_did | [string](#string) |  | agent is the DID of the agent that submits the evaluation |
| agent_address | [string](#string) |  |  |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |
| status | [EvaluationStatus](#fury.claims.v1beta1.EvaluationStatus) |  | status is the evaluation status expressed as an integer (2=approved, 3=rejected, ...) |
| reason | [uint32](#uint32) |  | reason is the code expressed as an integer, for why the evaluation result was given (codes defined by evaluator) |
| verification_proof | [string](#string) |  | verificationProof is the cid of the evaluation Verfiable Credential |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by evaluator for claim approval, if empty list then use default by Collection |






<a name="fury.claims.v1beta1.MsgEvaluateClaimResponse"></a>

### MsgEvaluateClaimResponse







<a name="fury.claims.v1beta1.MsgSubmitClaim"></a>

### MsgSubmitClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim in the cid hash format |
| agent_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| agent_address | [string](#string) |  |  |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="fury.claims.v1beta1.MsgSubmitClaimResponse"></a>

### MsgSubmitClaimResponse







<a name="fury.claims.v1beta1.MsgWithdrawPayment"></a>

### MsgWithdrawPayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id the withdrawal is for |
| inputs | [Input](#fury.claims.v1beta1.Input) | repeated | Inputs to the multisend tx to run to withdraw payment |
| outputs | [Output](#fury.claims.v1beta1.Output) | repeated | Outputs for the multisend tx to run to withdraw payment |
| payment_type | [PaymentType](#fury.claims.v1beta1.PaymentType) |  | payment type to keep track what payment is for and mark claim payment accordingly |
| contract_1155_payment | [Contract1155Payment](#fury.claims.v1beta1.Contract1155Payment) |  | if empty(nil) then no contract payment |
| toAddress | [string](#string) |  | for contract payment |
| fromAddress | [string](#string) |  | for contract payment |
| release_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date that grantee can execute authorization, calculated from created date plus the timeout on Collection payments |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="fury.claims.v1beta1.MsgWithdrawPaymentResponse"></a>

### MsgWithdrawPaymentResponse






 

 

 


<a name="fury.claims.v1beta1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCollection | [MsgCreateCollection](#fury.claims.v1beta1.MsgCreateCollection) | [MsgCreateCollectionResponse](#fury.claims.v1beta1.MsgCreateCollectionResponse) |  |
| SubmitClaim | [MsgSubmitClaim](#fury.claims.v1beta1.MsgSubmitClaim) | [MsgSubmitClaimResponse](#fury.claims.v1beta1.MsgSubmitClaimResponse) |  |
| EvaluateClaim | [MsgEvaluateClaim](#fury.claims.v1beta1.MsgEvaluateClaim) | [MsgEvaluateClaimResponse](#fury.claims.v1beta1.MsgEvaluateClaimResponse) |  |
| DisputeClaim | [MsgDisputeClaim](#fury.claims.v1beta1.MsgDisputeClaim) | [MsgDisputeClaimResponse](#fury.claims.v1beta1.MsgDisputeClaimResponse) |  |
| WithdrawPayment | [MsgWithdrawPayment](#fury.claims.v1beta1.MsgWithdrawPayment) | [MsgWithdrawPaymentResponse](#fury.claims.v1beta1.MsgWithdrawPaymentResponse) |  |

 



<a name="fury/entity/v1beta1/cosmos.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/cosmos.proto



<a name="fury.entity.v1beta1.Grant"></a>

### Grant
Grant gives permissions to execute
the provide method with expiration time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorization | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 

 



<a name="fury/iid/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/iid/v1beta1/types.proto



<a name="fury.iid.v1beta1.AccordedRight"></a>

### AccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| mechanism | [string](#string) |  |  |
| message | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="fury.iid.v1beta1.Context"></a>

### Context



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| val | [string](#string) |  |  |






<a name="fury.iid.v1beta1.IidMetadata"></a>

### IidMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versionId | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deactivated | [bool](#bool) |  |  |






<a name="fury.iid.v1beta1.LinkedClaim"></a>

### LinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [string](#string) |  |  |
| right | [string](#string) |  |  |






<a name="fury.iid.v1beta1.LinkedEntity"></a>

### LinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| relationship | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="fury.iid.v1beta1.LinkedResource"></a>

### LinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| mediaType | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [string](#string) |  |  |
| right | [string](#string) |  |  |






<a name="fury.iid.v1beta1.Service"></a>

### Service



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |






<a name="fury.iid.v1beta1.VerificationMethod"></a>

### VerificationMethod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| controller | [string](#string) |  |  |
| blockchainAccountID | [string](#string) |  |  |
| publicKeyHex | [string](#string) |  |  |
| publicKeyMultibase | [string](#string) |  |  |
| publicKeyBase58 | [string](#string) |  |  |





 

 

 

 



<a name="fury/iid/v1beta1/iid.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/iid/v1beta1/iid.proto



<a name="fury.iid.v1beta1.IidDocument"></a>

### IidDocument
type entity account
relationship entity account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| context | [Context](#fury.iid.v1beta1.Context) | repeated | @context is spec for did document. |
| id | [string](#string) |  | id represents the id for the did document. |
| controller | [string](#string) | repeated | A DID controller is an entity that is authorized to make changes to a DID document. cfr. https://www.w3.org/TR/did-core/#did-controller |
| verificationMethod | [VerificationMethod](#fury.iid.v1beta1.VerificationMethod) | repeated | A DID document can express verification methods, such as cryptographic public keys, which can be used to authenticate or authorize interactions with the DID subject or associated parties. https://www.w3.org/TR/did-core/#verification-methods |
| service | [Service](#fury.iid.v1beta1.Service) | repeated | Services are used in DID documents to express ways of communicating with the DID subject or associated entities. https://www.w3.org/TR/did-core/#services |
| authentication | [string](#string) | repeated | NOTE: below this line there are the relationships Authentication represents public key associated with the did document. cfr. https://www.w3.org/TR/did-core/#authentication |
| assertionMethod | [string](#string) | repeated | Used to specify how the DID subject is expected to express claims, such as for the purposes of issuing a Verifiable Credential. cfr. https://www.w3.org/TR/did-core/#assertion |
| keyAgreement | [string](#string) | repeated | used to specify how an entity can generate encryption material in order to transmit confidential information intended for the DID subject. https://www.w3.org/TR/did-core/#key-agreement |
| capabilityInvocation | [string](#string) | repeated | Used to specify a verification method that might be used by the DID subject to invoke a cryptographic capability, such as the authorization to update the DID Document. https://www.w3.org/TR/did-core/#capability-invocation |
| capabilityDelegation | [string](#string) | repeated | Used to specify a mechanism that might be used by the DID subject to delegate a cryptographic capability to another party. https://www.w3.org/TR/did-core/#capability-delegation |
| linkedResource | [LinkedResource](#fury.iid.v1beta1.LinkedResource) | repeated |  |
| linkedClaim | [LinkedClaim](#fury.iid.v1beta1.LinkedClaim) | repeated |  |
| accordedRight | [AccordedRight](#fury.iid.v1beta1.AccordedRight) | repeated |  |
| linkedEntity | [LinkedEntity](#fury.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| metadata | [IidMetadata](#fury.iid.v1beta1.IidMetadata) |  | Metadata concerning the IidDocument such as versionId, created, updated and deactivated |





 

 

 

 



<a name="fury/entity/v1beta1/entity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/entity.proto



<a name="fury.entity.v1beta1.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id represents the id for the entity document. |
| type | [string](#string) |  | Type of entity, eg protocol or asset |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| relayer_node | [string](#string) |  | Address of the operator through which the Entity was created |
| credentials | [string](#string) | repeated | Credentials of the enitity to be verified |
| entity_verified | [bool](#bool) |  | Used as check whether the credentials of entity is verified |
| metadata | [EntityMetadata](#fury.entity.v1beta1.EntityMetadata) |  | Metadata concerning the Entity such as versionId, created, updated and deactivated |
| accounts | [EntityAccount](#fury.entity.v1beta1.EntityAccount) | repeated | module accounts created for entity |






<a name="fury.entity.v1beta1.EntityAccount"></a>

### EntityAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| address | [string](#string) |  |  |






<a name="fury.entity.v1beta1.EntityMetadata"></a>

### EntityMetadata
EntityMetadata defines metadata associated to a entity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version_id | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="fury.entity.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nftContractAddress | [string](#string) |  |  |
| nftContractMinter | [string](#string) |  |  |
| createSequence | [uint64](#uint64) |  |  |





 

 

 

 



<a name="fury/entity/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/event.proto



<a name="fury.entity.v1beta1.EntityAccountAuthzCreatedEvent"></a>

### EntityAccountAuthzCreatedEvent
EntityAccountCreatedEvent is an event triggered on a entity account creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| account_name | [string](#string) |  |  |
| granter | [string](#string) |  |  |
| grantee | [string](#string) |  |  |
| grant | [Grant](#fury.entity.v1beta1.Grant) |  |  |






<a name="fury.entity.v1beta1.EntityAccountCreatedEvent"></a>

### EntityAccountCreatedEvent
EntityAccountCreatedEvent is an event triggered on a entity account creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| account_name | [string](#string) |  |  |
| account_address | [string](#string) |  |  |






<a name="fury.entity.v1beta1.EntityCreatedEvent"></a>

### EntityCreatedEvent
EntityCreatedEvent is an event triggered on a Entity creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#fury.entity.v1beta1.Entity) |  |  |
| signer | [string](#string) |  |  |






<a name="fury.entity.v1beta1.EntityTransferredEvent"></a>

### EntityTransferredEvent
EntityTransferredEvent is an event triggered on a entity transfer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |






<a name="fury.entity.v1beta1.EntityUpdatedEvent"></a>

### EntityUpdatedEvent
EntityUpdatedEvent is an event triggered on a entity document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#fury.entity.v1beta1.Entity) |  |  |
| signer | [string](#string) |  |  |






<a name="fury.entity.v1beta1.EntityVerifiedUpdatedEvent"></a>

### EntityVerifiedUpdatedEvent
EntityVerifiedUpdatedEvent is an event triggered on a entity verified
document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| entity_verified | [bool](#bool) |  |  |





 

 

 

 



<a name="fury/entity/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/genesis.proto



<a name="fury.entity.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the project module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entities | [Entity](#fury.entity.v1beta1.Entity) | repeated |  |
| params | [Params](#fury.entity.v1beta1.Params) |  |  |





 

 

 

 



<a name="fury/entity/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/proposal.proto



<a name="fury.entity.v1beta1.InitializeNftContract"></a>

### InitializeNftContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| NftContractCodeId | [uint64](#uint64) |  |  |
| NftMinterAddress | [string](#string) |  |  |





 

 

 

 



<a name="fury/entity/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/query.proto



<a name="fury.entity.v1beta1.QueryEntityIidDocumentRequest"></a>

### QueryEntityIidDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.entity.v1beta1.QueryEntityIidDocumentResponse"></a>

### QueryEntityIidDocumentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [fury.iid.v1beta1.IidDocument](#fury.iid.v1beta1.IidDocument) |  |  |






<a name="fury.entity.v1beta1.QueryEntityListRequest"></a>

### QueryEntityListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | string type = 2; string status = 3; |






<a name="fury.entity.v1beta1.QueryEntityListResponse"></a>

### QueryEntityListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entities | [Entity](#fury.entity.v1beta1.Entity) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="fury.entity.v1beta1.QueryEntityMetadataRequest"></a>

### QueryEntityMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.entity.v1beta1.QueryEntityMetadataResponse"></a>

### QueryEntityMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#fury.entity.v1beta1.Entity) |  |  |






<a name="fury.entity.v1beta1.QueryEntityRequest"></a>

### QueryEntityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.entity.v1beta1.QueryEntityResponse"></a>

### QueryEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#fury.entity.v1beta1.Entity) |  |  |
| iidDocument | [fury.iid.v1beta1.IidDocument](#fury.iid.v1beta1.IidDocument) |  |  |






<a name="fury.entity.v1beta1.QueryEntityVerifiedRequest"></a>

### QueryEntityVerifiedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.entity.v1beta1.QueryEntityVerifiedResponse"></a>

### QueryEntityVerifiedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_verified | [bool](#bool) |  |  |






<a name="fury.entity.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="fury.entity.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#fury.entity.v1beta1.Params) |  | params holds all the parameters of this module. |





 

 

 


<a name="fury.entity.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#fury.entity.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#fury.entity.v1beta1.QueryParamsResponse) |  |
| Entity | [QueryEntityRequest](#fury.entity.v1beta1.QueryEntityRequest) | [QueryEntityResponse](#fury.entity.v1beta1.QueryEntityResponse) |  |
| EntityMetaData | [QueryEntityMetadataRequest](#fury.entity.v1beta1.QueryEntityMetadataRequest) | [QueryEntityMetadataResponse](#fury.entity.v1beta1.QueryEntityMetadataResponse) |  |
| EntityIidDocument | [QueryEntityIidDocumentRequest](#fury.entity.v1beta1.QueryEntityIidDocumentRequest) | [QueryEntityIidDocumentResponse](#fury.entity.v1beta1.QueryEntityIidDocumentResponse) |  |
| EntityVerified | [QueryEntityVerifiedRequest](#fury.entity.v1beta1.QueryEntityVerifiedRequest) | [QueryEntityVerifiedResponse](#fury.entity.v1beta1.QueryEntityVerifiedResponse) |  |
| EntityList | [QueryEntityListRequest](#fury.entity.v1beta1.QueryEntityListRequest) | [QueryEntityListResponse](#fury.entity.v1beta1.QueryEntityListResponse) |  |

 



<a name="fury/iid/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/iid/v1beta1/tx.proto



<a name="fury.iid.v1beta1.MsgAddAccordedRight"></a>

### MsgAddAccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| accordedRight | [AccordedRight](#fury.iid.v1beta1.AccordedRight) |  | the Accorded right to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddAccordedRightResponse"></a>

### MsgAddAccordedRightResponse







<a name="fury.iid.v1beta1.MsgAddController"></a>

### MsgAddController



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did of the document |
| controller_did | [string](#string) |  | the did to add as a controller of the did document |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddControllerResponse"></a>

### MsgAddControllerResponse







<a name="fury.iid.v1beta1.MsgAddIidContext"></a>

### MsgAddIidContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| context | [Context](#fury.iid.v1beta1.Context) |  | the context to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddIidContextResponse"></a>

### MsgAddIidContextResponse







<a name="fury.iid.v1beta1.MsgAddLinkedClaim"></a>

### MsgAddLinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| linkedClaim | [LinkedClaim](#fury.iid.v1beta1.LinkedClaim) |  | the claim to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddLinkedClaimResponse"></a>

### MsgAddLinkedClaimResponse







<a name="fury.iid.v1beta1.MsgAddLinkedEntity"></a>

### MsgAddLinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the iid |
| linkedEntity | [LinkedEntity](#fury.iid.v1beta1.LinkedEntity) |  | the entity to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddLinkedEntityResponse"></a>

### MsgAddLinkedEntityResponse







<a name="fury.iid.v1beta1.MsgAddLinkedResource"></a>

### MsgAddLinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| linkedResource | [LinkedResource](#fury.iid.v1beta1.LinkedResource) |  | the verification to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddLinkedResourceResponse"></a>

### MsgAddLinkedResourceResponse







<a name="fury.iid.v1beta1.MsgAddService"></a>

### MsgAddService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| service_data | [Service](#fury.iid.v1beta1.Service) |  | the service data to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddServiceResponse"></a>

### MsgAddServiceResponse







<a name="fury.iid.v1beta1.MsgAddVerification"></a>

### MsgAddVerification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| verification | [Verification](#fury.iid.v1beta1.Verification) |  | the verification to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgAddVerificationResponse"></a>

### MsgAddVerificationResponse







<a name="fury.iid.v1beta1.MsgCreateIidDocument"></a>

### MsgCreateIidDocument
MsgCreateDidDocument defines a SDK message for creating a new did.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| controllers | [string](#string) | repeated | the list of controller DIDs |
| context | [Context](#fury.iid.v1beta1.Context) | repeated |  |
| verifications | [Verification](#fury.iid.v1beta1.Verification) | repeated | the list of verification methods and relationships |
| services | [Service](#fury.iid.v1beta1.Service) | repeated |  |
| accordedRight | [AccordedRight](#fury.iid.v1beta1.AccordedRight) | repeated |  |
| linkedResource | [LinkedResource](#fury.iid.v1beta1.LinkedResource) | repeated |  |
| linkedEntity | [LinkedEntity](#fury.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| signer | [string](#string) |  | address of the account signing the message |
| linkedClaim | [LinkedClaim](#fury.iid.v1beta1.LinkedClaim) | repeated |  |






<a name="fury.iid.v1beta1.MsgCreateIidDocumentResponse"></a>

### MsgCreateIidDocumentResponse







<a name="fury.iid.v1beta1.MsgDeactivateIID"></a>

### MsgDeactivateIID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| state | [bool](#bool) |  |  |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeactivateIIDResponse"></a>

### MsgDeactivateIIDResponse







<a name="fury.iid.v1beta1.MsgDeleteAccordedRight"></a>

### MsgDeleteAccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| right_id | [string](#string) |  | the Accorded right id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteAccordedRightResponse"></a>

### MsgDeleteAccordedRightResponse







<a name="fury.iid.v1beta1.MsgDeleteController"></a>

### MsgDeleteController



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did of the document |
| controller_did | [string](#string) |  | the did to remove from the list of controllers of the did document |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteControllerResponse"></a>

### MsgDeleteControllerResponse







<a name="fury.iid.v1beta1.MsgDeleteIidContext"></a>

### MsgDeleteIidContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| contextKey | [string](#string) |  | the context key |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteIidContextResponse"></a>

### MsgDeleteIidContextResponse







<a name="fury.iid.v1beta1.MsgDeleteLinkedClaim"></a>

### MsgDeleteLinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| claim_id | [string](#string) |  | the claim id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteLinkedClaimResponse"></a>

### MsgDeleteLinkedClaimResponse







<a name="fury.iid.v1beta1.MsgDeleteLinkedEntity"></a>

### MsgDeleteLinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the iid |
| entity_id | [string](#string) |  | the entity id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteLinkedEntityResponse"></a>

### MsgDeleteLinkedEntityResponse







<a name="fury.iid.v1beta1.MsgDeleteLinkedResource"></a>

### MsgDeleteLinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| resource_id | [string](#string) |  | the service id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteLinkedResourceResponse"></a>

### MsgDeleteLinkedResourceResponse







<a name="fury.iid.v1beta1.MsgDeleteService"></a>

### MsgDeleteService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| service_id | [string](#string) |  | the service id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgDeleteServiceResponse"></a>

### MsgDeleteServiceResponse







<a name="fury.iid.v1beta1.MsgRevokeVerification"></a>

### MsgRevokeVerification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| method_id | [string](#string) |  | the verification method id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgRevokeVerificationResponse"></a>

### MsgRevokeVerificationResponse







<a name="fury.iid.v1beta1.MsgSetVerificationRelationships"></a>

### MsgSetVerificationRelationships



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| method_id | [string](#string) |  | the verification method id |
| relationships | [string](#string) | repeated | the list of relationships to set |
| signer | [string](#string) |  | address of the account signing the message |






<a name="fury.iid.v1beta1.MsgSetVerificationRelationshipsResponse"></a>

### MsgSetVerificationRelationshipsResponse







<a name="fury.iid.v1beta1.MsgUpdateIidDocument"></a>

### MsgUpdateIidDocument
Updates the entity with all the fields, so if field empty will be updated
with default go type, aka never null


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| controllers | [string](#string) | repeated | the list of controller DIDs |
| context | [Context](#fury.iid.v1beta1.Context) | repeated |  |
| verifications | [Verification](#fury.iid.v1beta1.Verification) | repeated | the list of verification methods and relationships |
| services | [Service](#fury.iid.v1beta1.Service) | repeated |  |
| accordedRight | [AccordedRight](#fury.iid.v1beta1.AccordedRight) | repeated |  |
| linkedResource | [LinkedResource](#fury.iid.v1beta1.LinkedResource) | repeated |  |
| linkedEntity | [LinkedEntity](#fury.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| signer | [string](#string) |  | address of the account signing the message |
| linkedClaim | [LinkedClaim](#fury.iid.v1beta1.LinkedClaim) | repeated |  |






<a name="fury.iid.v1beta1.MsgUpdateIidDocumentResponse"></a>

### MsgUpdateIidDocumentResponse







<a name="fury.iid.v1beta1.Verification"></a>

### Verification
Verification is a message that allows to assign a verification method
to one or more verification relationships


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relationships | [string](#string) | repeated | verificationRelationships defines which relationships are allowed to use the verification method

relationships that the method is allowed into. |
| method | [VerificationMethod](#fury.iid.v1beta1.VerificationMethod) |  | public key associated with the did document. |
| context | [string](#string) | repeated | additional contexts (json ld schemas) |





 

 

 


<a name="fury.iid.v1beta1.Msg"></a>

### Msg
Msg defines the identity Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateIidDocument | [MsgCreateIidDocument](#fury.iid.v1beta1.MsgCreateIidDocument) | [MsgCreateIidDocumentResponse](#fury.iid.v1beta1.MsgCreateIidDocumentResponse) | CreateDidDocument defines a method for creating a new identity. |
| UpdateIidDocument | [MsgUpdateIidDocument](#fury.iid.v1beta1.MsgUpdateIidDocument) | [MsgUpdateIidDocumentResponse](#fury.iid.v1beta1.MsgUpdateIidDocumentResponse) | UpdateDidDocument defines a method for updating an identity. |
| AddVerification | [MsgAddVerification](#fury.iid.v1beta1.MsgAddVerification) | [MsgAddVerificationResponse](#fury.iid.v1beta1.MsgAddVerificationResponse) | AddVerificationMethod adds a new verification method |
| RevokeVerification | [MsgRevokeVerification](#fury.iid.v1beta1.MsgRevokeVerification) | [MsgRevokeVerificationResponse](#fury.iid.v1beta1.MsgRevokeVerificationResponse) | RevokeVerification remove the verification method and all associated verification Relations |
| SetVerificationRelationships | [MsgSetVerificationRelationships](#fury.iid.v1beta1.MsgSetVerificationRelationships) | [MsgSetVerificationRelationshipsResponse](#fury.iid.v1beta1.MsgSetVerificationRelationshipsResponse) | SetVerificationRelationships overwrite current verification relationships |
| AddService | [MsgAddService](#fury.iid.v1beta1.MsgAddService) | [MsgAddServiceResponse](#fury.iid.v1beta1.MsgAddServiceResponse) | AddService add a new service |
| DeleteService | [MsgDeleteService](#fury.iid.v1beta1.MsgDeleteService) | [MsgDeleteServiceResponse](#fury.iid.v1beta1.MsgDeleteServiceResponse) | DeleteService delete an existing service |
| AddController | [MsgAddController](#fury.iid.v1beta1.MsgAddController) | [MsgAddControllerResponse](#fury.iid.v1beta1.MsgAddControllerResponse) | AddService add a new service |
| DeleteController | [MsgDeleteController](#fury.iid.v1beta1.MsgDeleteController) | [MsgDeleteControllerResponse](#fury.iid.v1beta1.MsgDeleteControllerResponse) | DeleteService delete an existing service |
| AddLinkedResource | [MsgAddLinkedResource](#fury.iid.v1beta1.MsgAddLinkedResource) | [MsgAddLinkedResourceResponse](#fury.iid.v1beta1.MsgAddLinkedResourceResponse) | Add / Delete Linked Resource |
| DeleteLinkedResource | [MsgDeleteLinkedResource](#fury.iid.v1beta1.MsgDeleteLinkedResource) | [MsgDeleteLinkedResourceResponse](#fury.iid.v1beta1.MsgDeleteLinkedResourceResponse) |  |
| AddLinkedClaim | [MsgAddLinkedClaim](#fury.iid.v1beta1.MsgAddLinkedClaim) | [MsgAddLinkedClaimResponse](#fury.iid.v1beta1.MsgAddLinkedClaimResponse) | Add / Delete Linked Claims |
| DeleteLinkedClaim | [MsgDeleteLinkedClaim](#fury.iid.v1beta1.MsgDeleteLinkedClaim) | [MsgDeleteLinkedClaimResponse](#fury.iid.v1beta1.MsgDeleteLinkedClaimResponse) |  |
| AddLinkedEntity | [MsgAddLinkedEntity](#fury.iid.v1beta1.MsgAddLinkedEntity) | [MsgAddLinkedEntityResponse](#fury.iid.v1beta1.MsgAddLinkedEntityResponse) | Add / Delete Linked Entity |
| DeleteLinkedEntity | [MsgDeleteLinkedEntity](#fury.iid.v1beta1.MsgDeleteLinkedEntity) | [MsgDeleteLinkedEntityResponse](#fury.iid.v1beta1.MsgDeleteLinkedEntityResponse) |  |
| AddAccordedRight | [MsgAddAccordedRight](#fury.iid.v1beta1.MsgAddAccordedRight) | [MsgAddAccordedRightResponse](#fury.iid.v1beta1.MsgAddAccordedRightResponse) | Add / Delete Accorded Right |
| DeleteAccordedRight | [MsgDeleteAccordedRight](#fury.iid.v1beta1.MsgDeleteAccordedRight) | [MsgDeleteAccordedRightResponse](#fury.iid.v1beta1.MsgDeleteAccordedRightResponse) |  |
| AddIidContext | [MsgAddIidContext](#fury.iid.v1beta1.MsgAddIidContext) | [MsgAddIidContextResponse](#fury.iid.v1beta1.MsgAddIidContextResponse) | Add / Delete Context |
| DeactivateIID | [MsgDeactivateIID](#fury.iid.v1beta1.MsgDeactivateIID) | [MsgDeactivateIIDResponse](#fury.iid.v1beta1.MsgDeactivateIIDResponse) |  |
| DeleteIidContext | [MsgDeleteIidContext](#fury.iid.v1beta1.MsgDeleteIidContext) | [MsgDeleteIidContextResponse](#fury.iid.v1beta1.MsgDeleteIidContextResponse) |  |

 



<a name="fury/entity/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/entity/v1beta1/tx.proto



<a name="fury.entity.v1beta1.MsgCreateEntity"></a>

### MsgCreateEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_type | [string](#string) |  | An Entity Type as defined by the implementer |
| entity_status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| controller | [string](#string) | repeated | the list of controller DIDs |
| context | [fury.iid.v1beta1.Context](#fury.iid.v1beta1.Context) | repeated | JSON-LD contexts |
| verification | [fury.iid.v1beta1.Verification](#fury.iid.v1beta1.Verification) | repeated | Verification Methods and Verification Relationships |
| service | [fury.iid.v1beta1.Service](#fury.iid.v1beta1.Service) | repeated | Service endpoints |
| accorded_right | [fury.iid.v1beta1.AccordedRight](#fury.iid.v1beta1.AccordedRight) | repeated | Legal or Electronic Rights and associated Object Capabilities |
| linked_resource | [fury.iid.v1beta1.LinkedResource](#fury.iid.v1beta1.LinkedResource) | repeated | Digital resources associated with the Subject |
| linked_entity | [fury.iid.v1beta1.LinkedEntity](#fury.iid.v1beta1.LinkedEntity) | repeated | DID of a linked Entity and its relationship with the Subject |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| relayer_node | [string](#string) |  | Address of the operator through which the Entity was created |
| credentials | [string](#string) | repeated | Content ID or Hash of public Verifiable Credentials associated with the subject |
| owner_did | [string](#string) |  | Owner of the Entity NFT | The ownersdid used to sign this transaction. |
| owner_address | [string](#string) |  | The ownersdid address used to sign this transaction. |
| data | [bytes](#bytes) |  | Extention data |
| alsoKnownAs | [string](#string) |  |  |
| linked_claim | [fury.iid.v1beta1.LinkedClaim](#fury.iid.v1beta1.LinkedClaim) | repeated | Digital claims associated with the Subject |






<a name="fury.entity.v1beta1.MsgCreateEntityAccount"></a>

### MsgCreateEntityAccount
create a module account for an entity, account details will be added as a
linkedEntity on entity iid doc where linkedEntity id is didfragment: did#name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | entity id (did) to create account for |
| name | [string](#string) |  | name of account |
| owner_address | [string](#string) |  | The owner_address used to sign this transaction. |






<a name="fury.entity.v1beta1.MsgCreateEntityAccountResponse"></a>

### MsgCreateEntityAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account address that was created for specific entity and account name |






<a name="fury.entity.v1beta1.MsgCreateEntityResponse"></a>

### MsgCreateEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_id | [string](#string) |  |  |
| entity_type | [string](#string) |  |  |
| entity_status | [int32](#int32) |  |  |






<a name="fury.entity.v1beta1.MsgGrantEntityAccountAuthz"></a>

### MsgGrantEntityAccountAuthz
Create a authz grant from entity account (as grantor) to recipient in msg as
grantee for the specific authorization


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | entity id (did) to create account for |
| name | [string](#string) |  | name of account |
| grantee_address | [string](#string) |  | the grantee address that will be able to execute the authz authorization |
| grant | [Grant](#fury.entity.v1beta1.Grant) |  | grant to be Authorized in authz grant |
| owner_address | [string](#string) |  | the owner_address used to sign this transaction. |






<a name="fury.entity.v1beta1.MsgGrantEntityAccountAuthzResponse"></a>

### MsgGrantEntityAccountAuthzResponse







<a name="fury.entity.v1beta1.MsgTransferEntity"></a>

### MsgTransferEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| owner_did | [string](#string) |  | The owner_did used to sign this transaction. |
| owner_address | [string](#string) |  | The owner_address used to sign this transaction. |
| recipient_did | [string](#string) |  |  |






<a name="fury.entity.v1beta1.MsgTransferEntityResponse"></a>

### MsgTransferEntityResponse







<a name="fury.entity.v1beta1.MsgUpdateEntity"></a>

### MsgUpdateEntity
Updates the entity with all the fields, so if field empty will be updated
with default go type, aka never null


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id of entity to be updated |
| entity_status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| credentials | [string](#string) | repeated | Content ID or Hash of public Verifiable Credentials associated with the subject |
| controller_did | [string](#string) |  | The controllerDid used to sign this transaction. |
| controller_address | [string](#string) |  | The controllerAddress used to sign this transaction. |






<a name="fury.entity.v1beta1.MsgUpdateEntityResponse"></a>

### MsgUpdateEntityResponse







<a name="fury.entity.v1beta1.MsgUpdateEntityVerified"></a>

### MsgUpdateEntityVerified
Only relayer nodes can update entity field &#39;entityVerified&#39;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id of entity to be updated |
| entity_verified | [bool](#bool) |  | Whether entity is verified or not based on credentials |
| relayer_node_did | [string](#string) |  | The relayer node&#39;s did used to sign this transaction. |
| relayer_node_address | [string](#string) |  | The relayer node&#39;s address used to sign this transaction. |






<a name="fury.entity.v1beta1.MsgUpdateEntityVerifiedResponse"></a>

### MsgUpdateEntityVerifiedResponse






 

 

 


<a name="fury.entity.v1beta1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateEntity | [MsgCreateEntity](#fury.entity.v1beta1.MsgCreateEntity) | [MsgCreateEntityResponse](#fury.entity.v1beta1.MsgCreateEntityResponse) | CreateEntity defines a method for creating a entity. |
| UpdateEntity | [MsgUpdateEntity](#fury.entity.v1beta1.MsgUpdateEntity) | [MsgUpdateEntityResponse](#fury.entity.v1beta1.MsgUpdateEntityResponse) | UpdateEntity defines a method for updating a entity |
| UpdateEntityVerified | [MsgUpdateEntityVerified](#fury.entity.v1beta1.MsgUpdateEntityVerified) | [MsgUpdateEntityVerifiedResponse](#fury.entity.v1beta1.MsgUpdateEntityVerifiedResponse) | UpdateEntityVerified defines a method for updating if an entity is verified |
| TransferEntity | [MsgTransferEntity](#fury.entity.v1beta1.MsgTransferEntity) | [MsgTransferEntityResponse](#fury.entity.v1beta1.MsgTransferEntityResponse) | Transfers an entity and its nft to the recipient |
| CreateEntityAccount | [MsgCreateEntityAccount](#fury.entity.v1beta1.MsgCreateEntityAccount) | [MsgCreateEntityAccountResponse](#fury.entity.v1beta1.MsgCreateEntityAccountResponse) | Create a module account for an entity, |
| GrantEntityAccountAuthz | [MsgGrantEntityAccountAuthz](#fury.entity.v1beta1.MsgGrantEntityAccountAuthz) | [MsgGrantEntityAccountAuthzResponse](#fury.entity.v1beta1.MsgGrantEntityAccountAuthzResponse) | Create a authz grant from entity account |

 



<a name="fury/iid/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/iid/v1beta1/event.proto



<a name="fury.iid.v1beta1.IidDocumentCreatedEvent"></a>

### IidDocumentCreatedEvent
IidDocumentCreatedEvent is triggered when a new IidDocument is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#fury.iid.v1beta1.IidDocument) |  |  |






<a name="fury.iid.v1beta1.IidDocumentUpdatedEvent"></a>

### IidDocumentUpdatedEvent
DidDocumentUpdatedEvent is an event triggered on a DID document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#fury.iid.v1beta1.IidDocument) |  |  |





 

 

 

 



<a name="fury/iid/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/iid/v1beta1/genesis.proto



<a name="fury.iid.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the did module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iid_docs | [IidDocument](#fury.iid.v1beta1.IidDocument) | repeated |  |





 

 

 

 



<a name="fury/iid/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/iid/v1beta1/query.proto



<a name="fury.iid.v1beta1.QueryIidDocumentRequest"></a>

### QueryIidDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | did id of iid document querying |






<a name="fury.iid.v1beta1.QueryIidDocumentResponse"></a>

### QueryIidDocumentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#fury.iid.v1beta1.IidDocument) |  |  |






<a name="fury.iid.v1beta1.QueryIidDocumentsRequest"></a>

### QueryIidDocumentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="fury.iid.v1beta1.QueryIidDocumentsResponse"></a>

### QueryIidDocumentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocuments | [IidDocument](#fury.iid.v1beta1.IidDocument) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 

 

 


<a name="fury.iid.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IidDocuments | [QueryIidDocumentsRequest](#fury.iid.v1beta1.QueryIidDocumentsRequest) | [QueryIidDocumentsResponse](#fury.iid.v1beta1.QueryIidDocumentsResponse) | IidDocuments queries all iid documents that match the given status. |
| IidDocument | [QueryIidDocumentRequest](#fury.iid.v1beta1.QueryIidDocumentRequest) | [QueryIidDocumentResponse](#fury.iid.v1beta1.QueryIidDocumentResponse) | IidDocument queries a iid documents with an id. |

 



<a name="fury/token/v1beta1/token.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/token.proto



<a name="fury.token.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fury1155_contract_code | [uint64](#uint64) |  |  |






<a name="fury.token.v1beta1.Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  | generated on token intiation through MsgSetupMinter |
| class | [string](#string) |  | class is the token protocol entity DID (validated) |
| name | [string](#string) |  | name is the token name, which must be unique (namespace) |
| description | [string](#string) |  | description is any arbitrary description |
| image | [string](#string) |  | image is the image url for the token |
| type | [string](#string) |  | type is the token type (eg fury1155) |
| cap | [string](#string) |  | cap is the maximum number of tokens with this name that can be minted, 0 is unlimited |
| supply | [string](#string) |  | how much has already been minted for this Token type, aka the supply |
| paused | [bool](#bool) |  | stop allowance of token minter temporarily |
| stopped | [bool](#bool) |  | stop allowance of token minter permanently |
| retired | [TokensRetired](#fury.token.v1beta1.TokensRetired) | repeated | tokens that has been retired for this Token with specific name and contract address |
| cancelled | [TokensCancelled](#fury.token.v1beta1.TokensCancelled) | repeated | tokens that has been cancelled for this Token with specific name and contract address |






<a name="fury.token.v1beta1.TokenData"></a>

### TokenData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | media type value should always be &#34;application/json&#34;

credential link ***.ipfs |
| encrypted | [bool](#bool) |  |  |
| proof | [string](#string) |  |  |
| type | [string](#string) |  |  |
| id | [string](#string) |  | did of entity to map token to |






<a name="fury.token.v1beta1.TokenProperties"></a>

### TokenProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| name | [string](#string) |  | name is the token name, which is same as Token name |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| tokenData | [TokenData](#fury.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |






<a name="fury.token.v1beta1.TokensCancelled"></a>

### TokensCancelled



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="fury.token.v1beta1.TokensRetired"></a>

### TokensRetired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| jurisdiction | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |





 

 

 

 



<a name="fury/token/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/authz.proto



<a name="fury.token.v1beta1.MintAuthorization"></a>

### MintAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| constraints | [MintConstraints](#fury.token.v1beta1.MintConstraints) | repeated |  |






<a name="fury.token.v1beta1.MintConstraints"></a>

### MintConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| name | [string](#string) |  | name is the token name, which must be unique (namespace), will be verified against Token name provided on msgCreateToken |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| tokenData | [TokenData](#fury.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |





 

 

 

 



<a name="fury/token/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/tx.proto



<a name="fury.token.v1beta1.MintBatch"></a>

### MintBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the token name, which must be unique (namespace), will be verified against Token name provided on msgCreateToken |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| amount | [string](#string) |  | amount is the number of tokens to mint |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| token_data | [TokenData](#fury.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |






<a name="fury.token.v1beta1.MsgCancelToken"></a>

### MsgCancelToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#fury.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |






<a name="fury.token.v1beta1.MsgCancelTokenResponse"></a>

### MsgCancelTokenResponse







<a name="fury.token.v1beta1.MsgCreateToken"></a>

### MsgCreateToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| class | [string](#string) |  | class is the token protocol entity DID (validated) |
| name | [string](#string) |  | name is the token name, which must be unique (namespace) |
| description | [string](#string) |  | description is any arbitrary description |
| image | [string](#string) |  | image is the image url for the token |
| token_type | [string](#string) |  | type is the token type (eg fury1155) |
| cap | [string](#string) |  | cap is the maximum number of tokens with this name that can be minted, 0 is unlimited |






<a name="fury.token.v1beta1.MsgCreateTokenResponse"></a>

### MsgCreateTokenResponse







<a name="fury.token.v1beta1.MsgMintToken"></a>

### MsgMintToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |
| owner | [string](#string) |  | address of owner to mint for |
| mint_batch | [MintBatch](#fury.token.v1beta1.MintBatch) | repeated |  |






<a name="fury.token.v1beta1.MsgMintTokenResponse"></a>

### MsgMintTokenResponse







<a name="fury.token.v1beta1.MsgPauseToken"></a>

### MsgPauseToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |
| paused | [bool](#bool) |  | pause or unpause Token Minting allowance |






<a name="fury.token.v1beta1.MsgPauseTokenResponse"></a>

### MsgPauseTokenResponse







<a name="fury.token.v1beta1.MsgRetireToken"></a>

### MsgRetireToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#fury.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| jurisdiction | [string](#string) |  | jurisdiction is the jurisdiction of the token owner. A jurisdiction has the format: &lt;country-code&gt;[-&lt;sub-national-code&gt;[ &lt;postal-code&gt;]] The country-code must be 2 alphabetic characters, the sub-national-code can be 1-3 alphanumeric characters, and the postal-code can be up to 64 alphanumeric characters. Only the country-code is required, while the sub-national-code and postal-code are optional and can be added for increased precision. See the valid format for this below. |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |






<a name="fury.token.v1beta1.MsgRetireTokenResponse"></a>

### MsgRetireTokenResponse







<a name="fury.token.v1beta1.MsgStopToken"></a>

### MsgStopToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |






<a name="fury.token.v1beta1.MsgStopTokenResponse"></a>

### MsgStopTokenResponse







<a name="fury.token.v1beta1.MsgTransferToken"></a>

### MsgTransferToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| recipient | [string](#string) |  | address of receiver |
| tokens | [TokenBatch](#fury.token.v1beta1.TokenBatch) | repeated | all tokens must be in same smart contract |






<a name="fury.token.v1beta1.MsgTransferTokenResponse"></a>

### MsgTransferTokenResponse







<a name="fury.token.v1beta1.TokenBatch"></a>

### TokenBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id that identifies the token |
| amount | [string](#string) |  | amount is the number of tokens to transfer |





 

 

 


<a name="fury.token.v1beta1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateToken | [MsgCreateToken](#fury.token.v1beta1.MsgCreateToken) | [MsgCreateTokenResponse](#fury.token.v1beta1.MsgCreateTokenResponse) |  |
| MintToken | [MsgMintToken](#fury.token.v1beta1.MsgMintToken) | [MsgMintTokenResponse](#fury.token.v1beta1.MsgMintTokenResponse) |  |
| TransferToken | [MsgTransferToken](#fury.token.v1beta1.MsgTransferToken) | [MsgTransferTokenResponse](#fury.token.v1beta1.MsgTransferTokenResponse) |  |
| RetireToken | [MsgRetireToken](#fury.token.v1beta1.MsgRetireToken) | [MsgRetireTokenResponse](#fury.token.v1beta1.MsgRetireTokenResponse) |  |
| CancelToken | [MsgCancelToken](#fury.token.v1beta1.MsgCancelToken) | [MsgCancelTokenResponse](#fury.token.v1beta1.MsgCancelTokenResponse) |  |
| PauseToken | [MsgPauseToken](#fury.token.v1beta1.MsgPauseToken) | [MsgPauseTokenResponse](#fury.token.v1beta1.MsgPauseTokenResponse) |  |
| StopToken | [MsgStopToken](#fury.token.v1beta1.MsgStopToken) | [MsgStopTokenResponse](#fury.token.v1beta1.MsgStopTokenResponse) |  |

 



<a name="fury/token/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/event.proto



<a name="fury.token.v1beta1.TokenCancelledEvent"></a>

### TokenCancelledEvent
TokenCancelledEvent is an event triggered on a Token cancel execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#fury.token.v1beta1.TokenBatch) | repeated |  |






<a name="fury.token.v1beta1.TokenCreatedEvent"></a>

### TokenCreatedEvent
TokenCreatedEvent is an event triggered on a Token creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#fury.token.v1beta1.Token) |  |  |






<a name="fury.token.v1beta1.TokenMintedEvent"></a>

### TokenMintedEvent
TokenMintedEvent is an event triggered on a Token mint execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | the contract address of token contract being initialized |
| minter | [string](#string) |  | the token minter |
| owner | [string](#string) |  | the new tokens owner |
| amount | [string](#string) |  |  |
| tokenProperties | [TokenProperties](#fury.token.v1beta1.TokenProperties) |  |  |






<a name="fury.token.v1beta1.TokenPausedEvent"></a>

### TokenPausedEvent
TokenPausedEvent is an event triggered on a Token pause/unpause execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | the minter address |
| contract_address | [string](#string) |  |  |
| paused | [bool](#bool) |  |  |






<a name="fury.token.v1beta1.TokenRetiredEvent"></a>

### TokenRetiredEvent
TokenRetiredEvent is an event triggered on a Token retire execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#fury.token.v1beta1.TokenBatch) | repeated |  |






<a name="fury.token.v1beta1.TokenStoppedEvent"></a>

### TokenStoppedEvent
TokenStoppedEvent is an event triggered on a Token stopped execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | the minter address |
| contract_address | [string](#string) |  |  |
| stopped | [bool](#bool) |  |  |






<a name="fury.token.v1beta1.TokenTransferredEvent"></a>

### TokenTransferredEvent
TokenTransferedEvent is an event triggered on a Token transfer execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the old token owner |
| recipient | [string](#string) |  | the new tokens owner |
| tokens | [TokenBatch](#fury.token.v1beta1.TokenBatch) | repeated |  |






<a name="fury.token.v1beta1.TokenUpdatedEvent"></a>

### TokenUpdatedEvent
TokenUpdatedEvent is an event triggered on a Token update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#fury.token.v1beta1.Token) |  |  |





 

 

 

 



<a name="fury/token/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/genesis.proto



<a name="fury.token.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#fury.token.v1beta1.Params) |  |  |
| tokens | [Token](#fury.token.v1beta1.Token) | repeated |  |
| token_properties | [TokenProperties](#fury.token.v1beta1.TokenProperties) | repeated |  |





 

 

 

 



<a name="fury/token/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/proposal.proto



<a name="fury.token.v1beta1.SetTokenContractCodes"></a>

### SetTokenContractCodes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fury1155_contract_code | [uint64](#uint64) |  |  |





 

 

 

 



<a name="fury/token/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## fury/token/v1beta1/query.proto



<a name="fury.token.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="fury.token.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#fury.token.v1beta1.Params) |  | params holds all the parameters of this module. |






<a name="fury.token.v1beta1.QueryTokenDocRequest"></a>

### QueryTokenDocRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | minter address to get Token Doc for |
| contract_address | [string](#string) |  |  |






<a name="fury.token.v1beta1.QueryTokenDocResponse"></a>

### QueryTokenDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokenDoc | [Token](#fury.token.v1beta1.Token) |  |  |






<a name="fury.token.v1beta1.QueryTokenListRequest"></a>

### QueryTokenListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |
| minter | [string](#string) |  | minter address to get list for |






<a name="fury.token.v1beta1.QueryTokenListResponse"></a>

### QueryTokenListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |
| tokenDocs | [Token](#fury.token.v1beta1.Token) | repeated |  |






<a name="fury.token.v1beta1.QueryTokenMetadataRequest"></a>

### QueryTokenMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="fury.token.v1beta1.QueryTokenMetadataResponse"></a>

### QueryTokenMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| decimals | [string](#string) |  |  |
| image | [string](#string) |  |  |
| index | [string](#string) |  |  |
| properties | [TokenMetadataProperties](#fury.token.v1beta1.TokenMetadataProperties) |  |  |






<a name="fury.token.v1beta1.TokenMetadataProperties"></a>

### TokenMetadataProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [string](#string) |  |  |
| collection | [string](#string) |  |  |
| cap | [string](#string) |  |  |
| linkedResources | [TokenData](#fury.token.v1beta1.TokenData) | repeated |  |





 

 

 


<a name="fury.token.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#fury.token.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#fury.token.v1beta1.QueryParamsResponse) |  |
| TokenList | [QueryTokenListRequest](#fury.token.v1beta1.QueryTokenListRequest) | [QueryTokenListResponse](#fury.token.v1beta1.QueryTokenListResponse) |  |
| TokenDoc | [QueryTokenDocRequest](#fury.token.v1beta1.QueryTokenDocRequest) | [QueryTokenDocResponse](#fury.token.v1beta1.QueryTokenDocResponse) |  |
| TokenMetadata | [QueryTokenMetadataRequest](#fury.token.v1beta1.QueryTokenMetadataRequest) | [QueryTokenMetadataResponse](#fury.token.v1beta1.QueryTokenMetadataResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

