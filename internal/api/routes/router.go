package routes

import (
	"github.com/mercan/Go-Youtube-Subtitles/internal/api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(c *echo.Echo) {
	c.GET("/subtitles", controllers.GetSubtitlesHandler)
}
