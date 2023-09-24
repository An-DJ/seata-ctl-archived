package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of seata-ctl",
	Long:  `All software has versions. This is seata-ctl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0")
	},
}