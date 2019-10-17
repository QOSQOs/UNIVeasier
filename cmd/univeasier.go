package main

import (
	"os"

	"github.com/QOSQOs/UNIVeasier/pkg/api"
)

func main() {
	/*
		filename:= flag.String("config", "", "References to configuration file.")
		flag.Parse()
		s.Initialize(*filename)
	*/

	configPath := os.Getenv("CONFIG")

	s := api.Server{}
	s.Initialize(configPath)
	s.Run()
}
