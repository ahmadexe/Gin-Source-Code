package validations

import "github.com/go-playground/validator/v10"


func IsAdult(field validator.FieldLevel) bool {
	var age int8 = int8(field.Field().Int())
	return age >= 18
}