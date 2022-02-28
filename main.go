package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mercan/Go-Youtube-Subtitles/router"
)

func main() {
	// Echo Instance
	e := echo.New()

	// Routes
	router.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
