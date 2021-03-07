package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/famartinrh/uninterface/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	_ "gocloud.dev/runtimevar/filevar"
)

var cfgFile string
var appConfig *config.Config
var docstoresConfig map[string]*config.Docstore = make(map[string]*config.Docstore)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uninterface [options]",
	Short: "Universal interface for middleware and cloud services",
	// Args:         cobra.MaximumNArgs(2),
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initializeConfig(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// if len(args) == 0 {
		// 	return cmd.Help()
		// }
		return runServer(appConfig)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file, yaml")

	rootCmd.PersistentFlags().IntP("verbosity", "v", 0, "number for the log level verbosity")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))

}

func initializeConfig(cmd *cobra.Command) error {
	if cfgFile == "" {
		cfgFileEnv := os.Getenv("CONFIG_FILE")
		if cfgFileEnv == "" {
			return errors.New("config file not provided")
		} else {
			cfgFile = cfgFileEnv
		}
	}
	appConfig = &config.Config{}
	filebytes, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("The provided config file does not exist")
		}
		return err
	}
	err = yaml.Unmarshal(filebytes, appConfig)
	if err != nil {
		return errors.New("Error parsing config file " + err.Error())
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && viper.GetInt("verbosity") > 5 {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	return nil
}
