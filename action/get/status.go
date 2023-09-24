package get

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status",
	Long:  `Get the status`,
	Run: func(cmd *cobra.Command, args []string) {
		seata.GetStatus()
	},
}
