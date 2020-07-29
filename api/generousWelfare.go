package api

import (
	"net/http"

	"github.com/Coniglio/company-map/model"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetGenerousWelfares 福利厚生情報を返します
func GetGenerousWelfares() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		generousWelfares, err := model.GetGenerousWelfare(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, generousWelfares)
	}
}
