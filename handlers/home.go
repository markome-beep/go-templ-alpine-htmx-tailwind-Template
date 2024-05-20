package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markome-beep/go-templ-alpine-htmx-tailwind-Template/views/layouts"
)

func Home(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, layouts.Base())
}
