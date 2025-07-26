package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) []string {
	var errors []string

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			tag := err.Tag()
			param := err.Param()

			switch tag {
			case "required":
				errors = append(errors, field+" is required")
			case "min":
				errors = append(errors, field+" must be at least "+param+" characters")
			case "max":
				errors = append(errors, field+" must be at most "+param+" characters")
			case "oneof":
				errors = append(errors, field+" must be one of: "+param)
			default:
				errors = append(errors, field+" is invalid")
			}
		}
	}

	return errors
}

func IsEmptyValue(v interface{}) bool {
	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return val.IsNil()
	}
	return false
}
