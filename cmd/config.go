package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Topics []string
}

func GetConfig() ([]string, error) {

	config, err := os.UserConfigDir()
	home, err := os.UserHomeDir()

	viper.SetConfigName("cued")
	viper.SetConfigType("yml")
	viper.AddConfigPath(config)
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return viper.GetStringSlice("topics"), nil

}
