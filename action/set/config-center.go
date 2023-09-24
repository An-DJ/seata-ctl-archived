package se

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	configCenterKey string
	configCenterValue string
)

func init() {
	ConfigCenterCmd.PersistentFlags().StringVar(&configCenterKey, "key", "", "Configuration key")
	ConfigCenterCmd.PersistentFlags().StringVar(&configCenterValue, "value", "", "Configuration value")
}

var ConfigCenterCmd = &cobra.Command{
	Use:   "config-center",
	Short: "Set the config-center configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.SetConfigCenterConfiguration(configCenterKey, configCenterValue)
		configCenterKey = ""
		configCenterValue = ""
	},
}
