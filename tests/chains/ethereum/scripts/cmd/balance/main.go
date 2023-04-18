package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/datachainlab/yui-relayer-build/tests/chains/ethereum/scripts/cmd/helper"
	"github.com/spf13/cobra"
)

const (
	relayer = 0
)

var rootCmd = &cobra.Command{
	Use:   "wallet",
	Short: "wallet command",
	Long:  "wallet command walletIndex",
	Run: func(cmd *cobra.Command, args []string) {
		pathFile := args[0]
		walletIndex, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		balanceA, balanceB := balanceOf(pathFile, walletIndex)
		fmt.Printf("%d,%d\n", balanceA, balanceB)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func balanceOf(pathFile string, index int64) (*big.Int, *big.Int) {
	chainA, chainB, err := helper.InitializeChains(pathFile)
	if err != nil {
		log.Println("InitializeChains Error: ", err)
		os.Exit(1)
	}
	ctx := context.Background()
	baseDenom := strings.ToLower(chainA.ContractConfig.GetSimpleTokenAddress().String())
	bankA, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, uint32(index)).From, baseDenom)
	if err != nil {
		log.Println("BalanceOf Error: ", err)
		os.Exit(1)
	}
	chanB := chainB.GetChannel()
	expectedDenom := fmt.Sprintf("%v/%v/%v", chanB.PortID, chanB.ID, baseDenom)
	bankB, err := chainB.ICS20Bank.BalanceOf(chainB.CallOpts(ctx, relayer), chainB.CallOpts(ctx, uint32(index)).From, expectedDenom)
	if err != nil {
		log.Println("BalanceOf Error: ", err)
		os.Exit(1)
	}
	return bankA, bankB
}
