package main

import (
	"time"
	"errors"

	"github.com/spf13/viper"
)

const (
	ConfigPath string = "/home/mr-papi/SoftwareCode/Work/world-datascience-institute/AuctionHub-server/"
	ConfigFileName string = "AuctionHub-server.config"
	ConfigFileType string = "toml"
)

type Configuration struct {
    Web3Provider string `mapstructure:"Web3Provider"` 
    ContractAddress string `mapstructure:"ContractAddress"` 
	DbDriver   string `mapstructure:"DbDriver"`
	DbHost     string `mapstructure:"DbHost"`
	DbPort     int32  `mapstructure:"DbPort"`
	DbName     string `mapstructure:"DbName"`
	ServerHost string `mapstructure:"ServerHost"`
	ServerPort string `mapstructure:"ServerPort"`
	ServerBaseTimeout time.Duration `mapstructure:"ServerBaseTimeout"`
}

var Config Configuration

func InitConfig() (err error) {
	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	viper.AddConfigPath(ConfigPath)

	err = viper.ReadInConfig()
	if err != nil {
		panic(errors.New("failed to load config file"))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(errors.New("failed to save config"))
	}

	return nil
}
