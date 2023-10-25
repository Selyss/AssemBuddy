package chtsht

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
	if err != nil {
		log.Fatalf("Error while getting user config dir: %s", err)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error while getting user home dir: %s", err)
	}

	viper.SetConfigName("cued")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(config)

	// I hope im not searching home recursively
	viper.AddConfigPath(home)

	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	topics := viper.GetStringSlice("topics")
	return topics, nil
}
