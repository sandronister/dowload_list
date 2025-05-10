package di

import (
	"github.com/sandronister/download_list/config"
	midiahandler "github.com/sandronister/download_list/internal/infra/handler/midia_handler"
	"github.com/sandronister/download_list/internal/infra/web"
	midiausecase "github.com/sandronister/download_list/internal/usecase/midia_usecase"
	typelogger "github.com/sandronister/download_list/pkg/logger/types"
	"github.com/sandronister/download_list/pkg/system_memory_data/types"
)

func NewWebServer(broker types.IBroker, env *config.Enviroment, logger typelogger.ILogger) *web.EchoServer {
	usecase, err := midiausecase.NewMidiaUseCase(broker, env, logger)

	if err != nil {
		logger.Fatal("Failed to create usecase")
		return nil
	}

	handler := midiahandler.NewMidiaHandler(usecase)
	if handler == nil {
		logger.Fatal("Failed to create handler")
		return nil
	}

	server := web.NewServer(env.WebPort)
	if server == nil {
		logger.Fatal("Failed to create server")
		return nil
	}

	server.AddRoute("POST", "/midia", handler.GetMidia)

	return server
}
