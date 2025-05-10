package main

import (
	"fmt"
	"sync"

	"github.com/sandronister/download_list/config"
	"github.com/sandronister/download_list/internal/di"
	"github.com/sandronister/download_list/pkg/logger/factory"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	env, err := config.LoadEnviroment()

	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		return
	}

	logger, err := factory.NewLogger(env.LogPattern)

	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}

	midiaService, err := di.NewMidiaService(env, logger)

	if err != nil {
		logger.Error("DI", fmt.Sprintf("Error creating midia service: %v", err))
		return
	}

	go midiaService.WebServer()

	fmt.Println("Server started on port", env.WebPort)

	wg.Wait()

}
