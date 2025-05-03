package main

import (
	"homework8/internal/adapters/adrepo"
	"homework8/internal/app"
	"homework8/internal/ports/httpgin"
)

func main() {
	ser := httpgin.NewHTTPServer(":8080", app.NewApp(adrepo.New()))
	if err := ser.Listen(); err != nil {
		panic(err)
	}
}
