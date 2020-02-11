package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetAlongs 沿線情報を返します
func GetAlongs() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}
