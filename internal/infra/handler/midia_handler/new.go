package midiahandler

import midiausecase "github.com/sandronister/download_list/internal/usecase/midia_usecase"

type Handler struct {
	usecase *midiausecase.Usecase
}

func NewMidiaHandler(usecase *midiausecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
