// Copyright © 2018 ALP.Lab GmbH <office@alp-lab.at>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile, Host, outDir string
	Port                  int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alp",
	Short: "ALP.Lab Command Line Tool",
	Long:  `alp is a command line tool for developers to use services of ALP.Lab individually.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$HOME/.alp/alp.json", "config file")
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "host address of the alplab cloud service")
	rootCmd.PersistentFlags().IntP("port", "P", 443, "port")
	rootCmd.PersistentFlags().StringP("certificate", "C", "", "SSL/TLS certificate to use")
	rootCmd.PersistentFlags().StringP("output", "O", ".", "output directory")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("certificate", rootCmd.PersistentFlags().Lookup("certificate"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag if present
		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigName("alp")          // name of config file (without extension)
		viper.AddConfigPath(home + "/.alp") // call multiple times to add many search paths
		viper.AddConfigPath(".")            //
	}
	// Read in environment variables that match, but do nothing if not read
	viper.SetEnvPrefix("ALP")
	viper.AutomaticEnv()

	// Find and read the config files, but do nothing if not read
	viper.ReadInConfig()

}
