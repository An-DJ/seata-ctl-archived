package try

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/spf13/cobra"
)

func init() {
	TryCmd.AddCommand(BeginCmd)
	TryCmd.AddCommand(CommitCmd)
	TryCmd.AddCommand(RollbackCmd)
	TryCmd.SetUsageTemplate(common.GetUsageTmpl("try"))
	TryCmd.SetHelpTemplate(common.GetHelpTmpl())
}

var TryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try example transactions",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
