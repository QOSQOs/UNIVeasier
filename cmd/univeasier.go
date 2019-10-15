package main

import (
	"github.com/QOSQOs/UNIVeasier/pkg/api"
)

func main() {
	/*
		filename:= flag.String("config", "", "References to configuration file.")
		flag.Parse()
		s.Initialize(*filename)
	*/

	configPath := `C:\Users\i865255\go\src\github.com\QOSQOs\UNIVeasier\internal\config\config.json`

	s := api.Server{}
	s.Initialize(configPath)
	s.Run()
}
