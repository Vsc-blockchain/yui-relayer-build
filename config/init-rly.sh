#!/bin/bash

set -eux

#SCRIPT_DIR=$(cd $(dirname $0); pwd)

RLY_BIN=$(which yrly)
RLY_HOME="$HOME/.yui-relayer"
RLY="${RLY_BIN} --debug --home ${RLY_HOME}"
#FIXTURES_DIR=${SCRIPT_DIR}/../fixtures

echo "Generating ${RLY_BIN} configurations..."

# Ensure ${RLY_BIN} is installed
if ! [ -x ${RLY_BIN} ]; then
  echo "Error: ${RLY_BIN} is not installed." >&2
  exit 1
fi

rm -rf ${RLY_HOME} &> /dev/null

${RLY} config init
${RLY} chains add-dir config/fixed

# setup key for tendermint client
#SEED0=$(jq -r '.mnemonic' < ${FIXTURES_DIR}/tendermint/ibc0/key_seed.json)
#echo "Key $(${RLY} tendermint keys restore ibc0 testkey "$SEED0") imported from ibc0 to relayer..."