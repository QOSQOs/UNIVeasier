package main

import (
	"os"

	"github.com/QOSQOs/UNIVeasier/pkg/api"
)

func main() {
	configPath := os.Getenv("CONFIG")

	s := api.Server{}
	s.Initialize(configPath)
	s.Run()
}
