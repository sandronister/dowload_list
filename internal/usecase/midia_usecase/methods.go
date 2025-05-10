package midiausecase

import (
	"encoding/json"

	"github.com/sandronister/download_list/internal/dto"
	"github.com/sandronister/download_list/pkg/system_memory_data/types"
)

func (m *Usecase) PostUrl(request dto.Request) error {

	jsonReqquest, err := json.Marshal(request)

	if err != nil {
		m.logger.Error("Error marshaling request: ", err)
		return err
	}

	message := &types.Message{
		Topic: m.env.BrokerTopic,
		Value: jsonReqquest,
	}

	return m.broker.Publish(message)
}
