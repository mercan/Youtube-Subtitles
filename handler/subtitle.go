package handler

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/mercan/Go-Youtube-Subtitles/utils"

	"github.com/labstack/echo/v4"
)

type Video struct {
	ID        string           `json:"id"`
	Subtitles []utils.Subtitle `json:"subtitles"`
}

type SuccessResponse struct {
	Status int   `json:"status"`
	Data   Video `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GetSubtitle(c echo.Context) error {
	videoURL := c.QueryParam("url")
	text := c.QueryParam("text")

	if videoURL == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Video ID is required",
		})
	}

	if text == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Text is required",
		})
	}

	match, _ := regexp.MatchString("^((http|https)\\:\\/\\/)?(www\\.youtube\\.com|youtube\\.com|youtu\\.?be)\\/((watch\\?v=)?([a-zA-Z0-9_]{11}))(&.*)*$", videoURL)
	if !match {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid video id",
		})
	}

	u, err := url.Parse(videoURL)
	var videoId string

	if u.Host == "www.youtube.com" || u.Host == "youtube.com" {
		videoId = u.Query().Get("v") // Get the video id from the url
	} else {
		videoId = u.Path[1:] // remove first slash
	}

	if err := utils.DownloadSubtitle(videoId); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "The video has no subtitles",
		})
	}

	subtitles, err := utils.GetSubtitles(videoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	filteredSubtitles := utils.FilterSubtitles(videoId, subtitles, text)
	video := Video{
		ID:        videoId,
		Subtitles: filteredSubtitles,
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Status: http.StatusOK,
		Data:   video,
	})
}
