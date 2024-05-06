/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package eth

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// btcCmd represents the btc command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			os.Exit(0)
		}
		balance(cmd, args)
	},
}

func init() {
	balanceCmd.Flags().StringP("rpc", "r", "http://127.0.0.1:8545", "eth rpc")
	balanceCmd.Flags().StringP("address", "a", "", "the address you want to query, e.x 0x0000000000000000000000000000000000000000")
}

func balance(cmd *cobra.Command, args []string) {
	rpc, _ := cmd.Flags().GetString("rpc")
	address, _ := cmd.Flags().GetString("address")
	fmt.Println(rpc, address)
}
