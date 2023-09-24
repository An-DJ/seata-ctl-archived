package get

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	registryConfKey string
)

func init() {
	RegistryCmd.PersistentFlags().StringVar(&registryConfKey, "key", "", "Configuration key")
}

var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Get the registry configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.GetRegistryConfigurations(registryConfKey)
		registryConfKey = ""
	},
}
