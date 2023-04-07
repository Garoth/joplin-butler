package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	CONFIG_DIR_NAME  = ".joplin-butler"
	CONFIG_FILE_NAME = "config.json"
)

type Config struct {
	AuthToken string `json:"auth_token"`
	APIToken  string `json:"api_token"`
}

func (me *Config) Save() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(usr.HomeDir, CONFIG_DIR_NAME), 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	configPath := filepath.Join(usr.HomeDir, CONFIG_DIR_NAME, CONFIG_FILE_NAME)
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(me)
	if err != nil {
		return err
	}

	fmt.Println("Configuration saved:", configPath)
	return nil
}

func (me *Config) Load() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	configPath := filepath.Join(usr.HomeDir, CONFIG_DIR_NAME, CONFIG_FILE_NAME)
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(me)
	if err != nil {
		return err
	}

	fmt.Println("Configuration loaded:", configPath)
	return nil
}
