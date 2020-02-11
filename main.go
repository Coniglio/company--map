package main

import (
	"github.com/Coniglio/company-map/route"
)

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":80"))
}
