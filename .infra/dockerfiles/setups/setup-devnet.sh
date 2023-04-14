#!/usr/bin/env bash
PASSWORD=“12345678”
# fury init local --chain-id devnet-1
/app/fury init local --chain-id devnet-1
# yes ‘y’ | fury keys delete miguel --force
# yes ‘y’ | fury keys delete francesco --force
# yes ‘y’ | fury keys delete shaun --force
# yes ‘y’ | fury keys delete fee --force
# yes ‘y’ | fury keys delete fee2 --force
# yes ‘y’ | fury keys delete fee3 --force
# yes ‘y’ | fury keys delete fee4 --force
# yes ‘y’ | fury keys delete fee5 --force
# yes ‘y’ | fury keys delete reserveOut --force
# yes $PASSWORD | fury keys add miguel
# yes $PASSWORD | fury keys add francesco
# yes $PASSWORD | fury keys add shaun
# yes $PASSWORD | fury keys add fee
# yes $PASSWORD | fury keys add fee2
# yes $PASSWORD | fury keys add fee3
# yes $PASSWORD | fury keys add fee4
# yes $PASSWORD | fury keys add fee5
# yes $PASSWORD | fury keys add reserveOut
yes 12345678 | /app/fury keys add dev-main
# Note: important to add ‘miguel’ as a genesis-account since this is the chain’s validator
yes 12345678 | /app/fury add-genesis-account $(/app/fury keys show dev-main -a) 100000000000000000ufury

# yes $PASSWORD | fury add-genesis-account “$(fury keys show miguel -a)” 1000000000000ufury,1000000000000res,1000000000000rez,1000000000000uxgbp
# yes $PASSWORD | fury add-genesis-account “$(fury keys show francesco -a)” 1000000000000ufury,1000000000000res,1000000000000rez
# yes $PASSWORD | fury add-genesis-account “$(fury keys show shaun -a)” 1000000000000ufury,1000000000000res,1000000000000rez
# Add pubkey-based genesis accounts
# MIGUEL_ADDR=“fury107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx”    # address from did:fury:4XJLBfGtWSGKSz4BeRxdun’s pubkey
# FRANCESCO_ADDR=“fury1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y” # address from did:fury:UKzkhVSHc3qEFva5EY2XHt’s pubkey
# SHAUN_ADDR=“fury1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un”     # address from did:fury:U4tSpzzv91HHqWW1YmFkHJ’s pubkey
# yes $PASSWORD | fury add-genesis-account “$MIGUEL_ADDR” 1000000000000ufury,1000000000000res,1000000000000rez
# yes $PASSWORD | fury add-genesis-account “$FRANCESCO_ADDR” 1000000000000ufury,1000000000000res,1000000000000rez
# yes $PASSWORD | fury add-genesis-account “$SHAUN_ADDR” 1000000000000ufury,1000000000000res,1000000000000rez
# yes $PASSWORD | fury add-genesis-account “fury19h3lqj50uhzdrv8mkafnp55nqmz4ghc2sd3m48” 1000000000000ufury,1000000000000res,1000000000000rez
# Add fury did
yes 12345678 | /app/fury gentx dev-main 10000000000ufury --chain-id devnet-1
HOME=/root
/app/fury collect-gentxs

# FURY_DID=“did:fury:U4tSpzzv91HHqWW1YmFkHJ”
# export FROM=\”fury_did\“: \“\”
# export TO=“\”fury_did\“: \“$FURY_DID\“”
sed -i 's/"fury_did" : ""/"fury_did" : ""/;s/"bond_denom" : "stake"/"bond_denom" : "ufury"/;s/"mint_denom" : "stake"/"mint_denom" : "ufury"/;s/stake/ufury/;s/"Reserved_bond_tokens" : "\[\]"/"Reserved_bond_tokens" : "\[\]"/;s/"minimum-gas-prices" : ""/"minimum-gas-prices" : "0.025ufury"/;s/"enable" : "false"/"enable" : "true"/;s/"swagger" : "false"/"swagger" : "true"/;' $HOME/.fury/config/genesis.json
MAX_VOTING_PERIOD="30s"  # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.fury/config/genesis.json
# sed -i '/fury_did/c\   \"i xo_did\" : \"did:fury:aaaaaaaaaaa\",' $HOME/.fury/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
# export STAKING_TOKEN=“ufury”
# export FROM=“\”bond_denom\“: \“stake\“”
# export TO=“\”bond_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/bond_denom/c\   \"bond_denom\" : \"ufury\",' $HOME/.fury/config/genesis.json

# export FROM=“\”mint_denom\“: \“stake\“”
# export TO=“\”mint_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/mint_denom/c\   \"mint_denom\" : \"ufury\",' $HOME/.fury/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
# export FEE_TOKEN=“ufury”
# export FROM=“\”stake\“”
# export TO=“\”$FEE_TOKEN\“”
# sed 's/stake/ufury' $HOME/.fury/config/genesis.json

# Set reserved bond tokens
# export RESERVED_BOND_TOKENS=“”  # example: ” \“abc\“, \“def\“, \“ghi\” ”
# export FROM=“\”reserved_bond_tokens\“: \[\]”
# export TO=“\”reserved_bond_tokens\“: \[$RESERVED_BOND_TOKENS\]”
# sed -i '/reserved_bond_tokens/c\   \"Reserved_bond_tokens\" : \"[]\",' $HOME/.fury/config/genesis.json

# Set min-gas-prices (using fee token)
# export FROM=“minimum-gas-prices = \“\”
# export TO=“minimum-gas-prices = \“0.025$FEE_TOKEN\“”
# sed -i '/minimum-gas-prices/c\   \"minimum-gas-prices\" : \"0.025ufury\",' $HOME/.fury/config/genesis.json
# TODO: config missing from new version (REF: https://github.com/cosmos/cosmos-sdk/issues/8529)
#fury config chain-id devnet-1
#fury config output jsonW
#fury config indent true
#fury config trust-node true
# fury gentx miguel 1000000ufury --chain-id devnet-1
/app/fury validate-genesis
# Enable REST API (assumed to be at line 104 of app.toml)
# export FROM=“enable = false”
# TO=“enable = true”
# sed -i '/enable/c\   enable = true' $HOME/.fury/config/genesis.json
# Enable Swagger docs (assumed to be at line 107 of app.toml)
# export FROM=“swagger = false”
# export TO=“swagger = true”
# sed -i '/swagger/c\   swagger = true' $HOME/.fury/config/genesis.json
# Uncomment the below to broadcast node RPC endpoint
#FROM=“laddr = \“tcp:\/\/127.0.0.1:26657\“”
#TO=“laddr = \“tcp:\/\/0.0.0.0:26657\“”
#sed -i “s/$FROM/$TO/” “$HOME”/.fury/config/config.toml
# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i ‘s/timeout_commit = “5s”/timeout_commit = “1s”/g’ “$HOME”/.fury/config/config.toml
#sed -i ‘s/timeout_propose = “3s”/timeout_propose = “1s”/g’ “$HOME”/.fury/config/config.toml
# fury start --pruning “nothing”