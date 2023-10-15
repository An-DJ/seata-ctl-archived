package try

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	commitXID string
)

func init() {
	CommitCmd.SetUsageTemplate(common.GetUsageTmpl("commit"))
	CommitCmd.SetHelpTemplate(common.GetHelpTmpl())
	CommitCmd.PersistentFlags().StringVar(&commitXID, "xid", "", "commit a txn with xid")
}

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit a txn",
	Run: func(cmd *cobra.Command, args []string) {
		seata.CommitTxn(commitXID)
		commitXID = ""
	},
}
