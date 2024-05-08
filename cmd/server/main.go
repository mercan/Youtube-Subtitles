package main

import (
	"github.com/labstack/echo/v4"
	router "github.com/mercan/Go-Youtube-Subtitles/internal/api/routes"
)

func main() {
	// Echo Instance
	e := echo.New()

	// Routes
	router.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
