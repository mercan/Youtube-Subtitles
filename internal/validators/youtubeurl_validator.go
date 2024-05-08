package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func youtubeURLValidation(fl validator.FieldLevel) bool {
	url := fl.Field().String()

	if url == "" {
		return true
	}

	re := regexp.MustCompile("^((http|https)\\:\\/\\/)?(www\\.youtube\\.com|youtube\\.com|youtu\\.?be)\\/((watch\\?v=)?([a-zA-Z0-9_]{11}))(&.*)*$")

	return re.MatchString(url)
}
