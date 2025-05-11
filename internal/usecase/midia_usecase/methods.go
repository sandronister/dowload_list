package midiausecase

import (
	"encoding/json"

	"github.com/sandronister/download_list/internal/dto"
	"github.com/sandronister/download_list/pkg/system_memory_data/types"
)

func (m *Usecase) PostUrl(request dto.Request) []error {
	var errors []error

	for _, url := range request.Urls {
		urlMessage := dto.MessageListen{
			Url:     url,
			Audio:   request.Audio,
			Quality: request.Quality,
		}

		jsonReqquest, err := json.Marshal(urlMessage)

		if err != nil {
			m.logger.Error("Error marshaling request: ", err)
			errors = append(errors, err)
			continue
		}

		message := &types.Message{
			Topic: m.env.BrokerTopic,
			Value: jsonReqquest,
		}

		err = m.broker.Publish(message)

		if err != nil {
			m.logger.Error("Error publishing message: ", err)
			errors = append(errors, err)
			continue
		}
		m.logger.Info("Message published successfully")
	}

	return errors
}
