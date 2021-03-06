package cmd

import (
	"fmt"
	"os"

	"github.com/mikaelm1/pirate/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	VERSION    string
	DOService  service.DOService
	cfgFile    string
	Verbose    bool
	outputType string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pirate",
	Short: "\nPirate is a cli for accessing Digital Ocean's API\n",
	Long: `Pirate is a cli for managing your servers on Digital Ocean. 
Use it to create and delete droplets, manage ssh keys, and manage your balancers,
adding and removing droplets to them as needed.

Complete documentation is available at https://github.com/mikaelm1/Pirate.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	VERSION = version
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pirate.yaml)")
	RootCmd.PersistentFlags().StringVar(&outputType, "output", "text", "The type of output to diplay")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	// viper.SetConfigName(".pirate") // name of config file (without extension)
	// viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
