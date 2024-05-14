package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var youtubeURLRegex = regexp.MustCompile(`^(?:https?://)?(?:www\.)?(?:youtube\.com/watch\?v=|youtu\.be/)([a-zA-Z0-9_-]{11})$`)

func youtubeURLValidation(fl validator.FieldLevel) bool {
	url := fl.Field().String()

	if url == "" {
		return true
	}

	return youtubeURLRegex.MatchString(url)
}
