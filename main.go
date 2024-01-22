package main

import (
	"github.com/matfigueiredo/go-opportunities/config"
	"github.com/matfigueiredo/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
