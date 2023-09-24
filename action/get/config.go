package get

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	confKey string
)

func init() {
	ConfigCmd.PersistentFlags().StringVar(&confKey, "key", "", "Configuration key")
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.GetConfigurations(confKey)
		confKey = ""
	},
}
