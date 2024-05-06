/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package p2p

import (
	"encoding/hex"
	"errors"
	"fmt"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	p2pcrypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// btcCmd represents the btc command
var genNodeIdCmd = &cobra.Command{
	Use:   "gen-node-id",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			os.Exit(0)
		}
		run(cmd, args)
	},
}

func init() {
	genNodeIdCmd.Flags().StringP("private-key", "i", "", "private node key")
	genNodeIdCmd.Flags().StringP("public-key", "u", "", "public node key")
}

func run(cmd *cobra.Command, args []string) {

	publicKey, _ := cmd.Flags().GetString("public-key")
	privateKey, _ := cmd.Flags().GetString("private-key")
	var publicBz []byte
	if len(publicKey) != 0 {
		if strings.HasPrefix(publicKey, "0x") {
			publicKey = strings.TrimPrefix(publicKey, "0x")
		}
		if !strings.HasPrefix(publicKey, "04") {
			publicKey = "04" + publicKey
		}
		decoded, err := hex.DecodeString(publicKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding publicKey %s: %s", publicKey, err.Error())
		}
		publicBz = decoded
	} else if len(privateKey) != 0 {
		privKey, err := ethcrypto.HexToECDSA(privateKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error convert privateKey %s: %s", privateKey, err.Error())
		}
		pubkeybytes := ethcrypto.CompressPubkey(&privKey.PublicKey)
		publicBz = pubkeybytes
	} else {
		fmt.Fprintf(os.Stderr, "--private-key | --public-key at least one needs to be specified")
	}

	peerId, err := GetPeerIDFromSecp256PubKey(publicBz)
	if err != nil {
		fmt.Errorf("failed to convert peer id")
	}
	fmt.Println(peerId)
}

func GetPeerIDFromSecp256PubKey(pk []byte) (peer.ID, error) {
	if len(pk) == 0 {
		return "", errors.New("empty public key raw bytes")
	}
	ppk, err := p2pcrypto.UnmarshalSecp256k1PublicKey(pk)
	if err != nil {
		return "", fmt.Errorf("fail to convert pubkey to the crypto pubkey used in libp2p: %w", err)
	}

	return peer.IDFromPublicKey(ppk)
}
