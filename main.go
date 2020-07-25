package main

import (
	"github.com/Coniglio/company-map/router"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":1080"))
}
