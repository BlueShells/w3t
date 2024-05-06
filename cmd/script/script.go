package script

import (
	"github.com/spf13/cobra"
)

// cosmosCmd represents the cosmos command
var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "script utils",
}

func init() {
	scriptCmd.AddCommand(balanceCmd)
}
