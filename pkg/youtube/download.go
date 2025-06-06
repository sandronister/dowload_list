package youtube

import (
	"fmt"

	"github.com/sandronister/download_list/pkg/youtube/types"
)

func Download(entity *types.Input) error {

	err := getPath(entity)

	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	if entity.IsAudio() {
		return getAudio(entity)
	}

	if entity.IsQuality() {
		return getQuality(entity)
	}

	return getRegular(entity)

}
