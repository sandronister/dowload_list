package di

import (
	"fmt"

	"github.com/sandronister/download_list/config"
	"github.com/sandronister/download_list/internal/service/midia"
	"github.com/sandronister/download_list/pkg/logger/types"
	"github.com/sandronister/download_list/pkg/system_memory_data/factory"
)

func NewMidiaService(varEnv *config.Enviroment, logger types.ILogger) (*midia.Service, error) {
	broker, err := factory.GetBroker(varEnv.BrokerHost, varEnv.BrokerPort)

	if err != nil {
		logger.Error("DI", fmt.Sprintf("Error creating broker: %v", err))
		return nil, err
	}

	server := NewWebServer(broker, varEnv, logger)

	if server == nil {
		logger.Error("DI", "Error creating server")
		return nil, fmt.Errorf("error creating server")
	}

	return midia.NewMidiaService(
		server,
		broker,
		logger,
		varEnv,
	), nil

}
