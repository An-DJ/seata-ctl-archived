package cmd

import (
	"bufio"
	"fmt"
	"github.com/seata/seata-ctl/action"
	"github.com/seata/seata-ctl/seata"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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
		if !getArgs() {
			continue
		}
		if err := action.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func printPrompt(address string) {
	fmt.Printf("%s > ", address)
}

func getArgs() bool {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return false
	}
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	args := strings.Split(input, " ")

	os.Args = []string{""}
	for _, arg := range args {
		if arg != "" {
			os.Args = append(os.Args, arg)
		}
	}
	return true
}
