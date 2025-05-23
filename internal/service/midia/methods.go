package midia

import (
	"encoding/json"

	"github.com/sandronister/download_list/internal/dto"
	"github.com/sandronister/download_list/pkg/system_memory_data/types"
	youtube "github.com/sandronister/download_list/pkg/youtube"
	midiatype "github.com/sandronister/download_list/pkg/youtube/types"
)

func (m *Service) WebServer() error {
	err := m.server.Start()
	if err != nil {
		m.logger.Error("Midia service", "Error starting server: "+err.Error())
	}

	m.logger.Info("Midia service", "Server started successfully")
	return err
}

func (m *Service) ListenToQueue(message chan<- types.Message) {
	configBroker := &types.ConfigBroker{
		Topic: []string{m.varEnviroment.BrokerTopic},
	}

	err := m.broker.ListenToQueue(configBroker, message)

	if err != nil {
		m.logger.Error("Midia service", "Error listening to queue: "+err.Error())
	}
}

func (m *Service) ReadMessage(message <-chan types.Message) {
	for msg := range message {
		messageListen := new(dto.MessageListen)

		err := json.Unmarshal(msg.Value, &messageListen)

		if err != nil {
			m.logger.Error("Midia service", "Error unmarshalling message: "+err.Error())
		}

		requestDownload := &midiatype.Input{
			Url:     messageListen.Url,
			Path:    m.varEnviroment.DownloadPath,
			Kind:    messageListen.GetKind(),
			Quality: messageListen.GetQuality(),
		}

		err = youtube.Download(requestDownload)
		if err != nil {
			m.logger.Error("Midia service", "Error downloading video: "+err.Error())
		}
	}
}
