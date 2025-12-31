package utils

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(s)

	if err != nil {
		// explicit casting of err
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)

		}
	}
	return errors
}
