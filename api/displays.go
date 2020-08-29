package api

import (
	"net/http"

	"github.com/Coniglio/company-map/model"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetDisplayCompanies 表示する企業を返します
func GetDisplayCompanies() echo.HandlerFunc {
	return func(c echo.Context) error {
		languages := c.QueryParam("languages")
		alongs := c.QueryParam("alongs")
		generousWelfares := c.QueryParam("generousWelfares")

		tx := c.Get("Tx").(*gorp.Transaction)

		displays, err := model.GetDisplayCompanies(
			tx,
			languages,
			alongs,
			generousWelfares,
		)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, displays)
	}
}
