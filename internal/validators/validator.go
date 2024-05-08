package validators

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func init() {
	err := validate.RegisterValidation("youtubeURL", youtubeURLValidation)
	if err != nil {
		panic(err)
	}

}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
