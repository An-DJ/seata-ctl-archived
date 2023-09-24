package get

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/spf13/cobra"
)

func init() {
	GetCmd.AddCommand(StatusCmd)
	GetCmd.AddCommand(RegistryCmd)
	GetCmd.AddCommand(ConfigCenterCmd)
	GetCmd.AddCommand(ConfigCmd)

	GetCmd.SetUsageTemplate(common.GetUsageTmpl("get"))
	GetCmd.SetHelpTemplate(common.GetHelpTmpl())
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the resource",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
