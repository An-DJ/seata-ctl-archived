package get

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	confKeys string
)

func init() {
	ConfigCmd.PersistentFlags().StringVar(&confKeys, "key", "[]", "Configuration key")
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		params, err := common.ParseArrayArg(confKeys)
		if err != nil {
			common.Log("", err)
		}
		common.Log(seata.GetConfigurations(params))
		confKeys = "[]"
	},
}
