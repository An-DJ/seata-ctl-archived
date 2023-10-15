package cmd

import (
	"fmt"
	"github.com/seata/seata-ctl/action"
	"github.com/seata/seata-ctl/action/common"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "seata-ctl",
		Short: "seata-ctl is a CLI tool for Seata",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	credential := seata.GetAuth()
	rootCmd.PersistentFlags().StringVar(&credential.ServerIp, "ip", "127.0.0.1", "Seata Server IP")
	rootCmd.PersistentFlags().IntVar(&credential.ServerPort, "port", 7091, "Seata Server Admin Port")
	rootCmd.PersistentFlags().StringVar(&credential.Username, "username", "seata", "Username")
	rootCmd.PersistentFlags().StringVar(&credential.Password, "password", "seata", "Password")
	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "seata-ctl",
		Short: "seata-ctl is a CLI tool for Seata",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	})
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
		HiddenDefaultCmd:    true,
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" || arg == "version" {
			os.Exit(0)
		}
	}
	address := seata.GetAuth().GetAddress()
	err := seata.GetAuth().Login()
	if err != nil {
		fmt.Println("login failed!")
		os.Exit(1)
	}

	for {
		printPrompt(address)
		err = common.ReadArgs(os.Stdin)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err = action.Execute(); err != nil {
			fmt.Println(err)
			os.Args = []string{}
		}
	}
}

func printPrompt(address string) {
	fmt.Printf("%s > ", address)
}
