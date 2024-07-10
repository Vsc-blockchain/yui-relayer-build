#!/bin/sh
set -eux
LCP_BIN=${LCP_BIN:-lcp}
ADDRESSES_DIR=${ADDRESSES_DIR:-../vsg-bridge/addresses_test}
ENCLAVE_DIR=${ENCLAVE_DIR:-../lcp/bin}

MRENCLAVE=$(${LCP_BIN} enclave metadata --enclave=${ENCLAVE_DIR}/enclave.signed.so | jq -r .mrenclave)
IBC_ADDRESS=`cat $ADDRESSES_DIR/IBCHandler`
LC_ADDRESS=`cat $ADDRESSES_DIR/LCPClient`
TMCHAINID="vsc_420045-1"
ETHCHAINID=31337

mkdir -p config/fixed
jq -n -f config/tm.json.tpl --argjson ETHCHAINID ${ETHCHAINID} --arg TMCHAINID ${TMCHAINID} --arg MRENCLAVE ${MRENCLAVE} --arg LC_ADDRESS $LC_ADDRESS > config/fixed/ibc-0.json
jq -n -f config/eth.json.tpl --argjson ETHCHAINID ${ETHCHAINID} --arg MRENCLAVE ${MRENCLAVE} --arg IBC_ADDRESS ${IBC_ADDRESS} --arg LC_ADDRESS $LC_ADDRESS > config/fixed/ibc-1.json