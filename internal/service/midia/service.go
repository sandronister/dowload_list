package midia

import (
	"github.com/sandronister/download_list/config"
	"github.com/sandronister/download_list/internal/infra/web"
	"github.com/sandronister/download_list/internal/types"
	brokertypes "github.com/sandronister/download_list/pkg/system_memory_data/types"
)

type Service struct {
	server        *web.EchoServer
	broker        brokertypes.IBroker
	logger        types.ILog
	varEnviroment *config.Enviroment
}

func NewMidiaService(
	server *web.EchoServer,
	broker brokertypes.IBroker,
	logger types.ILog,
	varEnviroment *config.Enviroment,
) *Service {
	return &Service{
		server:        server,
		broker:        broker,
		logger:        logger,
		varEnviroment: varEnviroment,
	}
}
