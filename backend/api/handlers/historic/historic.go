package handlers

import (
	"net/http"

	providers "github.com/f1gopher/f1gopherlib/internal/providers"
	"github.com/labstack/echo"
)

func HandleHistoric(c echo.Context) error {
	var result []any
	for _, r := range providers.RaceHistory() {
		result = append(result, r)
	}

	c.JSON(http.StatusOK, result)

	return nil
}
