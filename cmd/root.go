// Copyright © 2016 Owen Butler <owen.butler@gmail.com>
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

package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/olivere/elastic.v1"
)

var PrintHeaders bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ges",
	Short: "command line es cluster stats and health",
	Long:  `command line es cluster stats and health`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&PrintHeaders, "verbose", "v", false, "print column headers")
	cobra.OnInitialize(initES)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ges.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var elasticClient = &elastic.Client{}

func initES() {
	client, err := elastic.NewClient(http.DefaultClient)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting client : %s\n", err)
		return
	}
	elasticClient = client
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// if cfgFile != "" { // enable ability to specify config file via flag
// viper.SetConfigFile(cfgFile)
// }

// viper.SetConfigName(".ges")  // name of config file (without extension)
// viper.AddConfigPath("$HOME") // adding home directory as first search path
// viper.AutomaticEnv()         // read in environment variables that match

// If a config file is found, read it in.
// if err := viper.ReadInConfig(); err == nil {
// fmt.Println("Using config file:", viper.ConfigFileUsed())
// }
// }
