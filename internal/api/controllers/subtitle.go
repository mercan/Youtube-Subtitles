package controllers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"github.com/mercan/Go-Youtube-Subtitles/internal/helpers"
	"github.com/mercan/Go-Youtube-Subtitles/internal/utils"
	"github.com/mercan/Go-Youtube-Subtitles/internal/validators"
)

type Video struct {
	ID        string             `json:"id"`
	Subtitles []helpers.Subtitle `json:"subtitles"`
}

type SubtitlesRequest struct {
	Url  string `query:"url" validate:"required,youtubeURL"`
	Text string `query:"text" validate:"required"`
}

func GetSubtitlesHandler(c echo.Context) error {
	// Parse and validate request
	var request SubtitlesRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request"))
	}

	if err := validators.ValidateStruct(request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
	}

	// Extract video ID from URL
	videoId, err := extractVideoID(request.Url)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
	}

	// Download subtitles
	if err := helpers.DownloadSubtitles(videoId); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
	}

	// Get subtitles for the video
	subtitles, err := helpers.GetSubtitles(videoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
	}

	// Filter subtitles based on the provided text
	filteredSubtitles := helpers.FilterSubtitles(videoId, request.Text, subtitles)

	return c.JSON(http.StatusOK, utils.NewSuccessResponse(Video{ID: videoId, Subtitles: filteredSubtitles}))
}

// extractVideoID extracts the video ID from a YouTube URL.
func extractVideoID(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if parsedURL.Host == "www.youtube.com" || parsedURL.Host == "youtube.com" {
		return parsedURL.Query().Get("v"), nil
	}

	// Handle https://youtu.be/<id> format
	return parsedURL.Path[1:], nil
}
