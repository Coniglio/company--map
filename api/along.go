package api

import (
	"net/http"

	"github.com/Coniglio/company-map/model"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetAlongs 沿線情報を返します
func GetAlongs() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		alongs, err := model.GetAlongs(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, alongs)
	}
}
