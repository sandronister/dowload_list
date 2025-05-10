package midiahandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sandronister/download_list/internal/dto"
)

func (h *Handler) GetMidia(c echo.Context) error {
	request := new(dto.Request)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := h.usecase.PostUrl(*request); err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, "URL posted successfully")
}
