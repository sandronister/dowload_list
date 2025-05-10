package midiausecase

import (
	"github.com/sandronister/download_list/config"
	internalType "github.com/sandronister/download_list/internal/types"
	typelogger "github.com/sandronister/download_list/pkg/logger/types"
	"github.com/sandronister/download_list/pkg/system_memory_data/types"
)

type Usecase struct {
	broker types.IBroker
	env    *config.Enviroment
	logger typelogger.ILogger
}

func NewMidiaUseCase(broker types.IBroker, env *config.Enviroment, logger typelogger.ILogger) (*Usecase, error) {
	if broker == nil {
		return nil, internalType.ErrNilBroker
	}

	if env == nil {
		return nil, internalType.ErrNilEnviroment
	}

	if logger == nil {
		return nil, internalType.ErrNilLogger
	}

	return &Usecase{
		broker: broker,
		env:    env,
		logger: logger,
	}, nil
}
