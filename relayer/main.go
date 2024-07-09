package main

import (
	"log"

	"github.com/datachainlab/ethereum-ibc-relay-chain/pkg/relay/ethereum"
	"github.com/datachainlab/ibc-hd-signer/pkg/hd"
	lcp "github.com/datachainlab/lcp-go/relay"
	rawsigner "github.com/datachainlab/lcp-go/relay/signers/raw"
	lcptm "github.com/datachainlab/lcp-go/relay/tendermint"
	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
	mock "github.com/hyperledger-labs/yui-relayer/provers/mock/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		mock.Module{},
		ethereum.Module{},
		hd.Module{},
		lcp.Module{},
		lcptm.Module{},
		rawsigner.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
