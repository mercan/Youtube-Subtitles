package router

import (
	"github.com/mercan/Go-Youtube-Subtitles/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(c *echo.Echo) {
	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"Time": [${time_rfc3339}], "Method": "${method}", "Path": "${path}", "Status": ${status}, "Latency": "${latency_human}", UserAgent: ${user_agent}, "Error": ${error}}`,
	})

	c.Use(logger)

	c.GET("/subtitle", handler.GetSubtitle)
}
