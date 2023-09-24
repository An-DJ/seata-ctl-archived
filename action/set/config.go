package se

import (
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
)

var (
	confKey   string
	confValue string
)

func init() {
	ConfigCmd.PersistentFlags().StringVar(&confKey, "key", "", "Configuration key")
	ConfigCmd.PersistentFlags().StringVar(&confValue, "value", "", "Configuration Value")
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		seata.SetConfiguration(confKey, confValue)
		confKey = ""
		confValue = ""
	},
}
