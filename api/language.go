package api

import (
	"net/http"

	"github.com/Coniglio/company-map/model"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetLanguages 言語情報を返します
func GetLanguages() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		languages, err := model.GetLanguages(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, languages)
	}
}
