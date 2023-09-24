package get

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	configCenterKey string
)

func init() {
	ConfigCenterCmd.PersistentFlags().StringVar(&configCenterKey, "key", "", "Configuration key")
}

var ConfigCenterCmd = &cobra.Command{
	Use:   "config-center",
	Short: "Get the config-center configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.GetConfigCenterConfigurations(configCenterKey)
		configCenterKey = ""
	},
}
