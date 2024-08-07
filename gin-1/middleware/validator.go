package middleware

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UserPassword(field validator.FieldLevel) bool {
	if match, _ := regexp.MatchString(`^[a-z][A-Z]\d$`, field.Field().String()); match {
		return true
	}
	return false
}
