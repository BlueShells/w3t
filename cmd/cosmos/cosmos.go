package cosmos

import (
	"github.com/spf13/cobra"
)

// cosmosCmd represents the cosmos command
var cosmosCmd = &cobra.Command{
	Use:   "cosmos",
	Short: "cosmos utils",
}

func init() {
	cosmosCmd.AddCommand(hdkeyCmd)
}
