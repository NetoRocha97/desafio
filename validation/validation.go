package validation

import (
   "github.com/go-playground/validator"
   "github.com/labstack/echo"
)

type (
   ValidationUtil struct {
      validator *validator.Validate
   }
)

func NewValidationUtil() echo.Validator {
   return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
   return v.validator.Struct(i)
}