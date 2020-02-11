package api

import (
	"net/http"

	"github.com/Coniglio/company-map/model"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetCompanyMaps 企業情マップ報を返します
func GetCompanyMaps() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		companyMaps, err := model.GetCompanyMaps(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, companyMaps)
	}
}
