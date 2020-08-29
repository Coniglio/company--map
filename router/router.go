package router

import (
	"os"
	"github.com/Coniglio/company-map/api"
	"github.com/Coniglio/company-map/db"
	"github.com/Coniglio/company-map/handler"
	mw "github.com/Coniglio/company-map/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init 初期化
func Init() *echo.Echo {
	e := echo.New()

	fp, err := os.OpenFile("/var/log/echo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fp,
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type", "Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization"},
	}))
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	e.Use(mw.TransactionHandler(db.Init()))

	v1 := e.Group("/api/v1")
	{
		v1.GET("/companymaps", api.GetCompanyMaps())
		v1.GET("/languages", api.GetLanguages())
		v1.GET("/alongs", api.GetAlongs())
		v1.GET("/generousWelfares", api.GetGenerousWelfares())
		v1.GET("/displays", api.GetDisplayCompanies())
	}

	return e
}
