#!/usr/bin/env bash

# Must be run from root path inside fury-blockchain for source to work
rm -rf ~/.fury

source ./scripts/constants.sh
fury init local --chain-id $CHAIN_ID --overwrite

# When incorporating another genesis file
# cp "$HOME"/.fury/config/genesis.json "$HOME"/.fury/config/genesis.json.backup #Backing up
# cp genesis.json "$HOME"/.fury/config/genesis.json #Copy over new genesis in root

# first add and remove a dummy user so PASSWORD can be set in keychain
yes $PASSWORD | fury keys add dummy &>/dev/null
yes $PASSWORD | fury keys delete ${USERS[i]} -y &>/dev/null

for ((i = 0; i < ${#USERS[@]}; ++i)); do
  # delete key if exists
  yes $PASSWORD | fury keys delete ${USERS[i]} -y 2>/dev/null
  # create key with constant mnemonic in /scripts/constants.sh
  (
    echo ${MNEMONICS[i]}
    echo $PASSWORD
  ) | fury keys add ${USERS[i]} --recover
  # add as genesis-account with fees
  yes $PASSWORD | fury add-genesis-account $(fury keys show ${USERS[i]} -a) 1000000000000ufury,1000000000000res,1000000000000rez,1000000000000uxgbp
done

# Add fury did
FURY_DID="did:fury:U4tSpzzv91HHqWW1YmFkHJ"
FROM="\"fury_did\": \"\""
TO="\"fury_did\": \"$FURY_DID\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
STAKING_TOKEN="ufury"
FROM="\"bond_denom\": \"stake\""
TO="\"bond_denom\": \"$STAKING_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json
FROM="\"mint_denom\": \"stake\""
TO="\"mint_denom\": \"$STAKING_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
FEE_TOKEN="ufury"
FROM="\"stake\""
TO="\"$FEE_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set reserved bond tokens
RESERVED_BOND_TOKENS="" # example: " \"abc\", \"def\", \"ghi\" "
FROM="\"reserved_bond_tokens\": \[\]"
TO="\"reserved_bond_tokens\": \[$RESERVED_BOND_TOKENS\]"
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

# Set min-gas-prices (using fee token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025$FEE_TOKEN\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/app.toml

MAX_VOTING_PERIOD="90s" # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json

yes $PASSWORD | fury gentx miguel 1000000ufury --chain-id $CHAIN_ID
fury collect-gentxs
fury validate-genesis

# Enable REST API
FROM="enable = false"
TO="enable = true"
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/app.toml

# Enable Swagger docs
FROM="swagger = false"
TO="swagger = true"
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/app.toml

# Broadcast node RPC endpoint
FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
sed -i -e "s/$FROM/$TO/" "$HOME"/.fury/config/config.toml

# Set timeouts to 1s for shorter block times
sed -i -e "s/timeout_commit = "5s"/timeout_commit = "1s"/g" "$HOME"/.fury/config/config.toml
sed -i -e "s/timeout_propose = "3s"/timeout_propose = "1s"/g" "$HOME"/.fury/config/config.toml

# fury start --pruning "nothing" --log_level "trace" --trace
fury start --pruning "nothing"
