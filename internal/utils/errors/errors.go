package errors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorList(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	out := []ErrorMsg{}

	if errors.As(err, &ve) {
		for _, fe := range ve {
			out = append(out, ErrorMsg{fe.Field(), getErrorMsg(fe)})
		}
	}

	return out
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"

	case "lte":
		return "Should be less than " + fe.Param()

	case "gte":
		return "Should be greater than " + fe.Param()
	}

	return "Unknown error"
}
