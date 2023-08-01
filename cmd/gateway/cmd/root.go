/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"deviceservice/internal/conf"
	"deviceservice/internal/service"
	"deviceservice/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deviceservice",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	 Run: func(cmd *cobra.Command, args []string) {
	 	service.RegisterService(conf.C)
	 	service.Start()
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deviceservice.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		b, err := os.ReadFile(cfgFile)
		if err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
		//fmt.Printf("配置文件打印 \n %s",string(b))
		viper.SetConfigType("yaml")
		if err = viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			log.WithError(err).Fatal("error loading config file")
			os.Exit(utils.ConfigFileErr)
		}
	} else {
		viper.SetConfigName("deviceservice")
		viper.AddConfigPath("../../internal/conf/")
		viper.AddConfigPath("/etc/deviceservice/")
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatal("error loading config file")
			os.Exit(utils.ConfigFileErr)
		}
	}

	if err := viper.Unmarshal(&conf.C); err != nil {
		log.WithError(err).Fatal("unmarshal config error")
		os.Exit(utils.ConfigFileErr)
	}
	log.Debug("Read config:",conf.C)
}


func GetConfig() {
	initConfig()
}
