package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetLanguages 言語情報を返します
func GetLanguages() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
