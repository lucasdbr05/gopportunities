package main

import (
	"github.com/lucasdbr05/gopportunities/config"
	"github.com/lucasdbr05/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("MAIN")

	err := config.Init()

	if err != nil {
		logger.Errorf("Oh SHIT! Config initialization error: %v", err)
		return
	}

	router.Init()
}
