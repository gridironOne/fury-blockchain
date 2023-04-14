#!/usr/bin/env bash

# For development purposes this script assumes you already ran run_with_all_data for app.toml changes and 1s block time

echo "Exporting app state to genesis file..."
fury export >genesis.json

echo "Backing up existing genesis file..."
cp "$HOME"/.fury/config/genesis.json "$HOME"/.fury/config/genesis.json.backup

echo "Moving new genesis file to $HOME/.fury/config/genesis.json..."
mv genesis.json "$HOME"/.fury/config/genesis.json

fury tendermint unsafe-reset-all
fury validate-genesis

fury start --pruning "nothing"
