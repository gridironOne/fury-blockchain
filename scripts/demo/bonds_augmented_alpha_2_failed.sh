#!/usr/bin/env bash

# Must be run from root path inside fury-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

FEE=$(yes $PASSWORD | fury keys show fee -a)
RESERVE_OUT=$(yes $PASSWORD | fury keys show reserveOut -a)

BOND_DID="did:fury:U7GK8p8rVhJMKhBVRCJJ8c"
#BOND_DID_FULL='{
#  "did":"did:fury:U7GK8p8rVhJMKhBVRCJJ8c",
#  "verifyKey":"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW",
#  "encryptionPublicKey":"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m",
#  "secret":{
#    "seed":"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053",
#    "signKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC",
#    "encryptionPrivateKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC"
#  }
#}'

MIGUEL_ADDR="fury107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"
FRANCESCO_ADDR="fury1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y"
SHAUN_ADDR="fury1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"
MIGUEL_DID_FULL='{
  "did":"did:fury:4XJLBfGtWSGKSz4BeRxdun",
  "verifyKey":"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt",
  "encryptionPublicKey":"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS",
  "secret":{
    "seed":"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180",
    "signKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh",
    "encryptionPrivateKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh"
  }
}'
FRANCESCO_DID="did:fury:UKzkhVSHc3qEFva5EY2XHt"
FRANCESCO_DID_FULL='{
  "did":"did:fury:UKzkhVSHc3qEFva5EY2XHt",
  "verifyKey":"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej",
  "encryptionPublicKey":"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si",
  "secret":{
    "seed":"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de",
    "signKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM",
    "encryptionPrivateKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM"
  }
}'
SHAUN_DID_FULL='{
  "did":"did:fury:U4tSpzzv91HHqWW1YmFkHJ",
  "verifyKey":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
  "encryptionPublicKey":"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i",
  "secret":{
    "seed":"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e",
    "signKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR",
    "encryptionPrivateKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR"
  }
}'

# Ledger DIDs
echo "Ledgering DID 1/3..."
fury_tx did add-did-doc "$MIGUEL_DID_FULL" 
echo "Ledgering DID 2/3..."
fury_tx did add-did-doc "$FRANCESCO_DID_FULL" 
echo "Ledgering DID 3/3..."
fury_tx did add-did-doc "$SHAUN_DID_FULL" 

# d0 := 500.0   // initial raise (reserve)
# p0 := 0.01    // initial price (reserve per token)
# theta := 0.4  // initial allocation (percentage)
# kappa := 3.0  // degrees of polynomial (i.e. x^2, x^4, x^6)

# R0 = 300              // initial reserve (1-theta)*d0
# S0 = 50000            // initial supply
# V0 = 416666666666.667 // invariant

echo "Creating bond..."
fury_tx bonds create-bond \
  --token=abc \
  --name="A B C" \
  --description="Description about A B C" \
  --function-type=augmented_function \
  --function-parameters="d0:500.0,p0:0.01,theta:0.4,kappa:3.0" \
  --reserve-tokens=res \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --reserve-withdrawal-address="$RESERVE_OUT" \
  --max-supply=1000000abc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells \
  --alpha-bond \
  --batch-blocks=1 \
  --outcome-payment="100000" \
  --bond-did="$BOND_DID" \
  --creator-did="$MIGUEL_DID_FULL" \
  --controller-did="$FRANCESCO_DID"
echo "Created bond..."
fury_q bonds bond "$BOND_DID"

echo "Miguel buys 20000abc..."
fury_tx bonds buy 20000abc 100000res "$BOND_DID" "$MIGUEL_DID_FULL" 
echo "Miguel's account..."
fury_q bank balances "$MIGUEL_ADDR"

echo "Francesco buys 20000abc..."
fury_tx bonds buy 20000abc 100000res "$BOND_DID" "$FRANCESCO_DID_FULL" 
echo "Francesco's account..."
fury_q bank balances "$FRANCESCO_ADDR"

echo "Francesco updates the bond state to FAILED"
fury_tx bonds update-bond-state "FAILED" "$BOND_DID" "$FRANCESCO_DID_FULL"

echo "Miguel withdraws share..."
fury_tx bonds withdraw-share "$BOND_DID" "$MIGUEL_DID_FULL" 
echo "Miguel's account..."
fury_q bank balances "$MIGUEL_ADDR"

echo "Francesco withdraws share..."
fury_tx bonds withdraw-share "$BOND_DID" "$FRANCESCO_DID_FULL" 
echo "Francesco's account..."
fury_q bank balances "$FRANCESCO_ADDR"

echo "Bond reserve is now empty and supply is 0..."
fury_q bonds bond "$BOND_DID"
