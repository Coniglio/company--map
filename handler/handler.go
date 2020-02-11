package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// CustomHTTPErrorHandler Jsonエラーハンドラー
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Error()
	}

	c.JSON(code, map[string]interface{}{
		"statusCode": code,
		"message":    msg,
	})
}
