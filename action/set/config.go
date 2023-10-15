package se

import (
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	kvData          string
	setRegistry     bool
	setConfigCenter bool
)

func init() {
	ConfigCmd.PersistentFlags().StringVar(&kvData, "data", "{}", "Configuration map")
	ConfigCmd.PersistentFlags().BoolVar(&setRegistry, "registry", false, "If set registry conf")
	ConfigCmd.PersistentFlags().BoolVar(&setConfigCenter, "config-center", false, "If set configuration center conf")
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := common.ParseDictArg(kvData)
		if err != nil {
			common.Log("", err)
		}
		configType := seata.NORMAL_CONFIG
		if setRegistry {
			configType = seata.REGISTRY_CONF
		} else if setConfigCenter {
			configType = seata.CONFIG_CENTER_CONF
		}
		common.Log(seata.SetConfiguration(data, configType))
		kvData = "{}"
		setRegistry = false
		setConfigCenter = false
		configType = seata.NORMAL_CONFIG
	},
}
