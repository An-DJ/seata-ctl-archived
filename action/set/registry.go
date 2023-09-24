package se

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	registryConfKey   string
	registryConfValue string
)

func init() {
	RegistryCmd.PersistentFlags().StringVar(&registryConfKey, "key", "", "Configuration key")
	RegistryCmd.PersistentFlags().StringVar(&registryConfValue, "value", "", "Configuration Value")
}

var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Set the registry configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.SetRegistryConfiguration(registryConfKey, registryConfValue)
		registryConfKey = ""
		registryConfValue = ""
	},
}
