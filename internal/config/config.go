package config

import (
	"github-io/internal/constant"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct{
	Auth struct {
		Token string `yaml:"token"`
	} `yaml:"auth"`
}

func InsertToken(token string) error {
	path := filepath.Join(constant.CONFIG_PATH, "/main.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	config.Auth.Token = token

	newData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, newData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetToken() (string, error) {
	path := filepath.Join(constant.CONFIG_PATH, "/main.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	return config.Auth.Token, nil
}