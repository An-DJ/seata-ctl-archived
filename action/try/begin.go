package try

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	timeout int
)

func init() {
	BeginCmd.SetUsageTemplate(common.GetUsageTmpl("begin"))
	BeginCmd.SetHelpTemplate(common.GetHelpTmpl())
	BeginCmd.PersistentFlags().IntVar(&timeout, "timeout", 3000, "begin a txn with timeout")
}

var BeginCmd = &cobra.Command{
	Use:   "begin",
	Short: "begin a txn",
	Run: func(cmd *cobra.Command, args []string) {
		seata.BeginTxn(timeout)
		timeout = 3000
	},
}
