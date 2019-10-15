package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/QOSQOs/UNIVeasier/internal/common"
)

type SERVERConfig struct {
	HOST string `json:"host"`
	PORT string `json:"port"`
}

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
	DB     *DBConfig
	SERVER *SERVERConfig
}

func ReadConfiguration(filepath string) (*Config, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		common.Log.Infow("Cannot find the configuration file", "filepath", filepath, "info", err.Error())
		return nil, err
	}

	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		common.Log.Infow("Configuration file read failed", "filepath", filepath, "info", err.Error())
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
