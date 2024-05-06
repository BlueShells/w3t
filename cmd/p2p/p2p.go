package p2p

import (
	"github.com/spf13/cobra"
)

// P2pCmd represents the p2p command
var P2pCmd = &cobra.Command{
	Use:   "p2p",
	Short: "p2p utils",
}

func init() {
	P2pCmd.AddCommand(genNodeIdCmd)
}
