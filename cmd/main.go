package main

import (
	"github.com/gofor-little/env"
	"github.com/labstack/echo/v4"
	"github.com/markome-beep/go-templ-alpine-htmx-tailwind-Template/handlers"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	app.GET("/", handlers.Home)
	port := env.Get("PORT", ":3000")

	app.Logger.Fatal(app.Start(port))
}
