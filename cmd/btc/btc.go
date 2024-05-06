package btc

import (
	"github.com/spf13/cobra"
)

// BtcCmd represents the btc command
var BtcCmd = &cobra.Command{
	Use:   "btc",
	Short: "bitcoin utils",
}

func init() {
	BtcCmd.AddCommand(hdkeyCmd)
}
