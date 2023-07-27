package conf

import (
	"deviceservice/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var C Config

type Ffs struct {
	Addr         string `mapstructure:"addr"`
	ProductKey   string `mapstructure:"productKey"`
	DeviceName   string `mapstructure:"DeviceName"`
	DeviceSecret string `mapstructure:"DeviceSecret"`
}
type Landleaf struct {
	Addr string `mapstructure:"addr"`
	Tsl  struct {
		Property []string
		Event    []string
		Service  []string
	} `mapstructure:"tsl"`
}
type Gateway struct {
	ProductKey   string `mapstructure:"productKey"`
	DeviceName   string `mapstructure:"deviceName"`
	DeviceSecret string `mapstructure:"deviceSecret"`
}

type Config struct {
	Gateway  `mapstructure:"Gateway"`
	Ffs      `mapstructure:"ffs"`
	Landleaf `mapstructure:"landleaf"`
	//MySQL struct {
	//	Database    string   `mapstructure:"database"`
	//	Host        string   `mapstructure:"host"`
	//	Port        string   `mapstructure:"port"`
	//	User        string   `mapstructure:"user"`
	//	Password    string   `mapstructure:"password"`
	//	Automigrate bool     `mapstructure:"automigrate"`
	//	MaxIdle     int      `mapstructure:"max_idle"`
	//	MaxOpen     int      `mapstructure:"max_open"`
	//	Timezone    string   `mapstructure:"timezone"`
	//} `mapstructure:"mysql"`
	//
	//Redis struct {
	//	Network   string `mapstructure:"network"`
	//	Host      string `mapstructure:"host"`
	//	Password  string `mapstructure:"password"`
	//	Database  int    `mapstructure:"database"`
	//	URL       string `mapstructure:"url"`
	//	MaxIdle   int    `mapstructure:"max_idle"`
	//	MaxActive int    `mapstructure:"max_active"`
	//} `mapstructure:"redis"`
}

func GetConf() {
	viper.SetConfigName("deviceservice")
	viper.AddConfigPath("$HOME/go/src/deviceservice/internal/conf/")
	viper.AddConfigPath("/etc/deviceservice/")
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal("error loading config file")
		os.Exit(utils.ConfigFileErr)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.WithError(err).Fatal("unmarshal config error")
		os.Exit(utils.ConfigFileErr)
	}
}
