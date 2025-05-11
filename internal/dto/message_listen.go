package dto

type MessageListen struct {
	Url     string `json:"url"`
	Audio   bool   `json:"audio"`
	Quality bool   `json:"quality"`
}

func (m *MessageListen) GetKind() string {
	if m.Audio {
		return "A"
	}
	return "V"
}

func (m *MessageListen) GetQuality() string {
	if m.Quality {
		return "S"
	}
	return "N"
}
