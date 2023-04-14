#!/usr/bin/env bash

# Must be run from root path inside fury-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

# Create multi-sig key using Miguel, Francesco, and Shaun, but with a 2 key
# threshold, meaning that only two people have to sign any transactions.
fury keys delete mfs-multisig -y  # remove just in case it already exists
fury keys add mfs-multisig \
  --multisig miguel,francesco,shaun \
  --multisig-threshold 2

MULTISIG_ADDR=$(yes $PASSWORD | fury keys show mfs-multisig -a)
MIGUEL_ADDR=$(yes $PASSWORD | fury keys show miguel -a)
SHAUN_ADDR=$(yes $PASSWORD | fury keys show shaun -a)

# Send tokens to the multi-sig address
fury_tx bank send miguel "$MULTISIG_ADDR" 1000000ufury
# Check balance of the multi-sig address
fury_q bank balances "$MULTISIG_ADDR"

# Send some tokens back to Miguel
#
# ...part 1: generate transaction
fury tx bank send "$MULTISIG_ADDR" "$MIGUEL_ADDR" 123ufury \
  --generate-only \
  --gas-prices="$GAS_PRICES" \
  --chain-id="$CHAIN_ID" \
  > tx.json
#
# ...part 2: miguel signs tx.json and produces tx-signed-miguel.json
fury tx sign tx.json \
  --from "$MIGUEL_ADDR" \
  --multisig "$MULTISIG_ADDR" \
  --sign-mode amino-json \
  --chain-id="$CHAIN_ID" \
  >> tx-signed-miguel.json
#
# ...part 2: shaun signs tx.json and produces tx-signed-shaun.json
fury tx sign tx.json \
  --from "$SHAUN_ADDR" \
  --multisig "$MULTISIG_ADDR" \
  --sign-mode amino-json \
  --chain-id="$CHAIN_ID" \
  >> tx-signed-shaun.json
#
# ...part 3: join signatures
fury tx multisign tx.json mfs-multisig tx-signed-miguel.json tx-signed-shaun.json \
  --from mfs-multisig \
  --chain-id="$CHAIN_ID" \
  > tx_ms.json
#
# ...part 4: broadcast
fury_tx broadcast ./tx_ms.json

# Check Miguel's balance
fury_q bank balances "$MIGUEL_ADDR"
# Check balance of the multi-sig address
fury_q bank balances "$MULTISIG_ADDR"

# Cleanup
fury keys delete mfs-multisig -y
rm tx.json tx-signed-miguel.json tx-signed-shaun.json tx_msg.json