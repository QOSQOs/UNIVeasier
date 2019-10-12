package main

import (
	"encoding/json"
	"io/ioutil"
)

type Database struct {
	DRIVER   string `json:"driver"`
	NAME     string `json:"name"`
	USERNAME string `json:"username"`
	PASSWORD string `json:"password"`
	PROTOCOL string `json:"protocol"`
	HOST     string `json:"host"`
	PORT     string `json:"port"`
}

type Config struct {
	DB *Database `json:"db"`
}

func GetConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorw("Cannot read the config file", "Error", err.Error())
		return nil, err
	}

	log.Info(string(buf))

	var config Config
	err = json.Unmarshal(buf, &config)
	if err != nil {
		log.Errorw("Cannot unmarshal", "Error", err.Error())
		return nil, err
	}

	return &config, nil
}
