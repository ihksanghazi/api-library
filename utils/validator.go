package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validation(v *validator.Validate, request interface{}) []string {
	var validationErrors []string
	if err := v.Struct(request); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("Validation error in field '%s' with value '%s' : %s", e.StructField(), e.Value(), e.Tag()))
		}
	}
	return validationErrors
}
