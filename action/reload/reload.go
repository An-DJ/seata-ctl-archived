package reload

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

func init() {
	ReloadCmd.SetUsageTemplate(common.GetUsageTmpl("reload"))
	ReloadCmd.SetHelpTemplate(common.GetHelpTmpl())
}

var ReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.ReloadConfiguration()
	},
}
