package se

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/spf13/cobra"
)

func init() {
	SetCmd.AddCommand(ConfigCenterCmd)
	SetCmd.AddCommand(RegistryCmd)
	SetCmd.AddCommand(ConfigCmd)
	SetCmd.SetUsageTemplate(common.GetUsageTmpl("set"))
	SetCmd.SetHelpTemplate(common.GetHelpTmpl())
}

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the resource",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
