package main

import (
	"embed"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Noisey96/budget-tracker/templates"
)

//go:embed static
var static embed.FS

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", HomeHandler)

	// Post-Router Middleware
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		Filesystem: http.FS(static),
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} | ${uri} | ${status}\n",
	}))
	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

// Handler
func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Hello("John"))
}