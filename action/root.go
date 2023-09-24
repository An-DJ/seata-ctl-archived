package action

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/action/get"
	"github.com/seata/seata-ctl/action/reload"
	se "github.com/seata/seata-ctl/action/set"
	del "github.com/seata/seata-ctl/action/try"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(quitCmd,
		get.GetCmd, del.TryCmd,
		reload.ReloadCmd, se.SetCmd)
	rootCmd.SetHelpTemplate(common.GetHelpTmplWithOnlyAvailableCmd())
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
		HiddenDefaultCmd:    true,
	}
}

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() error {
	return rootCmd.Execute()
}
