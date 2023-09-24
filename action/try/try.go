package try

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

func init() {
	TryCmd.SetUsageTemplate(common.GetUsageTmpl("try"))
	TryCmd.SetHelpTemplate(common.GetHelpTmpl())
}

var TryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try if this node is ready",
	Run: func(cmd *cobra.Command, args []string) {
		seata.TryTxn()
	},
}
