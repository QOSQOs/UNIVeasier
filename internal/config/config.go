package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/QOSQOs/UNIVeasier/internal/common"
)

type DBConfig struct {
	DRIVER   string `json:"driver"`
	NAME     string `json:"name"`
	USERNAME string `json:"username"`
	PASSWORD string `json:"password"`
	PROTOCOL string `json:"protocol"`
	HOST     string `json:"host"`
	PORT     string `json:"port"`
}

type Config struct {
	DB *DBConfig
}

func ReadConfiguration(filename string) (*Config, error) {
	_, err := os.Stat(filename)
	if err != nil {
		common.Log.Infow("Cannot find the configuration file", "filename", filename, "info", err.Error())
		return nil, err
	}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		common.Log.Infow("Configuration file read failed", "filename", filename, "info", err.Error())
		return nil, err
	}

	var config Config
	err = json.Unmarshal(buf, &config)
	if err != nil {
		common.Log.Infow("Configuration Unmarshal failed", "buffer", string(buf), "info", err.Error())
		return nil, err
	}

	return &config, nil
}
