package action

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	quitCmd.SetUsageTemplate(common.GetUsageTmpl("quit"))
	quitCmd.SetHelpTemplate(common.GetHelpTmpl())
}

var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "Quit the session",
	Run: func(cmd *cobra.Command, args []string) {
		println("Quit the session")
		os.Exit(0)
	},
}
