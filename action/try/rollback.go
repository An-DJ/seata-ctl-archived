package try

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	rollbackXID string
)

func init() {
	RollbackCmd.SetUsageTemplate(common.GetUsageTmpl("rollback"))
	RollbackCmd.SetHelpTemplate(common.GetHelpTmpl())
	RollbackCmd.PersistentFlags().StringVar(&rollbackXID, "xid", "", "rollback a txn with xid")
}

var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback a txn",
	Run: func(cmd *cobra.Command, args []string) {
		seata.RollbackTxn(rollbackXID)
		rollbackXID = ""
	},
}
