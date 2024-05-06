package eth

import (
	"github.com/spf13/cobra"
)

// EthCmd represents the eth command
var EthCmd = &cobra.Command{
	Use:   "eth",
	Short: "ethereum utils",
}

func init() {
	EthCmd.AddCommand(hdkeyCmd)
}
