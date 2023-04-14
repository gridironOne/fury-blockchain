#!/usr/bin/env bash

# IT IS RECOMMENDED TO CHANGE THE VOTING PERIOD IN run_with_all_data.sh
# TO 30 seconds FOR FASTER GOVERNANCE

# Must be run from root path inside fury-blockchain for source to work
source ./scripts/constants.sh

wait_chain_start

echo "Query transfer params before param change"
fury_q ibc-transfer params

echo "Submitting param change proposal"
fury_tx gov submit-proposal param-change gov_param_change_proposal.json --from=miguel

echo "Query proposal 1"
fury_q gov proposal 1

echo "Depositing 10000000ufury to reach minimum deposit"
fury_tx gov deposit 1 10000000ufury --from=miguel

echo "Query proposal 1 deposits"
fury_q gov deposits 1

echo "Voting yes for proposal"
fury_tx gov vote 1 yes --from=miguel

echo "Query proposal 1 tally"
fury_q gov tally 1

echo "Waiting for proposal to pass..."
while :; do
  RET=$(fury_q gov proposal 1 2>&1)
  if [[ ($RET == *'PROPOSAL_STATUS_VOTING_PERIOD'*) ]]; then
    sleep 1
  else
    echo "A few more seconds..."
    sleep 6
    break
  fi
done

echo "Query transfer params (expected to be true and false)"
fury_q ibc-transfer params
