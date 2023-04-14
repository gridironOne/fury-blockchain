#!/usr/bin/env bash

# Must be run from root path inside fury-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

NEW_DID="$RANDOM"
FULL_DID="did:earth:pandora-4:$NEW_DID"

echo "Creating DID..."
DID=$(yes $PASSWORD | fury tx iid create-iid "$NEW_DID" "pandora-4" --from miguel --from miguel --chain-id pandora-4 --fees 5000ufury -y | jq .)
echo $DID

#echo "Adding 2 contexts.."
#CONTEXT1=$(yes $PASSWORD | fury tx iid add-iid-context "$NEW_DID" "fury" "https://w3id.org/fury/NS/" --from miguel --from miguel --chain-id pandora-4 --fees 5000ufury -y --output json | jq .)
#echo $CONTEXT1

#CONTEXT2=$(yes $PASSWORD | fury tx iid add-iid-context "$NEW_DID" "iid" "https://w3id.org/fury/NS/"  --from miguel --from miguel --chain-id pandora-4 --fees 5000ufury -y --output json | jq .)
#echo $CONTEXT2

echo "Adding metadata..."
META3=$(yes $PASSWORD | fury tx iid update-iid-meta "$NEW_DID" '{"versionID":"1","deactivated":false,"entityType":"nft","startDate":null,"endDate":null,"status":1,"stage":"yes","relayerNode":"yes","verifiableCredential":"yes","credentials":[]}'  --from miguel --from miguel --chain-id pandora-4 --fees 5000ufury -y --output json)
echo $META3

echo "Querying DID..."
echo $FULL_DID
QUERY_DID=$(fury query iid iid "$FULL_DID" --chain-id pandora-4 --output json | jq .)

echo $QUERY_DID

echo "Changing metadata..."
META3=$(yes $PASSWORD | fury tx iid update-iid-meta "$NEW_DID" '{"versionID":"2","deactivated":false,"entityType":"stove","startDate":null,"endDate":null,"status":1,"stage":"yes","relayerNode":"yes","verifiableCredential":"yes","credentials":[]}'  --from miguel --from miguel --chain-id pandora-4 --fees 5000ufury -y --output json)
echo "Querying IID METADATA"
QUERY_DID=$(fury query iid metadata "$FULL_DID" --chain-id pandora-4 --output json | jq .)
echo "Deactivating IID"
DEAC=$Fnft(yes $PASSWORD | fury tx iid deactivate-iid "$NEW_DID" "true"  --from miguel --from miguel --chain-id pandora-4 --fees 5000ufury -y --output json)
echo "Querying IID METADATA"
QUERY_DID=$(fury query iid metadata "$FULL_DID" --chain-id pandora-4 --output json | jq .)
echo $QUERY_DID
