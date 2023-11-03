package main

import (
	"github.com/jtonynet/gopportunities/config"
	"github.com/jtonynet/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
	}

	router.Initialize()
}
