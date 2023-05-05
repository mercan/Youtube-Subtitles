package main

import (
	router "github.com/mercan/Go-Youtube-Subtitles/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo Instance
	e := echo.New()

	// Routes
	router.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
