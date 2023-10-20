package main

import (
	"os"
	"path"
)

func getHomeConfig() (string, error) {

	config, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(config, ".cued.json"), nil
}

func getUserConfig() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(config, "cued", "cued.json"), nil
}

func GetConfig() (string, error) {
	config, err := getUserConfig()
	if err != nil {
		return "", err
	}

	config, err = getHomeConfig()
	if err != nil {
		return "", err
	}

}
