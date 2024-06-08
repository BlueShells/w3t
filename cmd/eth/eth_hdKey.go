/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package eth

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"w3t/cmd/eth/hdwallet"
)

// btcCmd represents the btc command
var hdkeyCmd = &cobra.Command{
	Use:   "hdkey",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			os.Exit(0)
		}
		hdkey(cmd, args)
	},
}

func init() {
	hdkeyCmd.Flags().IntP("number", "N", 10, "set number of keys to generate")
	hdkeyCmd.Flags().StringP("mnemonic", "m", "", "optional list of words to re-generate a root key")
	hdkeyCmd.Flags().StringP("hd-path", "H", "", "HD Path")
	hdkeyCmd.Flags().BoolP("print-pubkey", "p", false, "print pubkey or not")
}

func hdkey(cmd *cobra.Command, args []string) {
	number, _ := cmd.Flags().GetInt("number")
	mnemonic, _ := cmd.Flags().GetString("mnemonic")
	hdPath, _ := cmd.Flags().GetString("hd-path")
	printPubkey, _ := cmd.Flags().GetBool("print-pubkey")

	if hdPath != "" {
		addr, privKey, pubKey := genWallet(mnemonic, hdPath)
		fmt.Printf("%-8s %-2s\n", "hdPath:", hdPath)
		fmt.Printf("%-8s %-2s\n", "address:", addr)
		fmt.Printf("%-8s %-2s\n", "privKey:", privKey)
		fmt.Printf("%-8s %-2s\n", "pubKey:", pubKey)
	} else {
		for i := 0; i <= number; i++ {
			hdPath = "m/44'/60'/0'/0/" + strconv.Itoa(i)
			addr, privKey, pubKey := genWallet(mnemonic, hdPath)
			if printPubkey {
				fmt.Printf("%-18s %s %s %s\n", hdPath, addr, privKey, pubKey)
			} else {
				fmt.Printf("%-18s %s %s\n", hdPath, addr, privKey)
			}
		}
	}
}

func genWallet(mnemonic string, hdPath string) (string, string, string) {

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath(hdPath)
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}
	return account.Address.Hex(), privateKey, publicKey
}
