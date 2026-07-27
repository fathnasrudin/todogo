package validator

import (
	"reflect"
	"strings"

	pgValidator "github.com/go-playground/validator/v10"
)


func NewValidator()pgValidator.Validate {
	 validator := pgValidator.New();

	// Tell validator to use the "json" tag name if it exists
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return *validator
}