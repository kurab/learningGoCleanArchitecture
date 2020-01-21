package validation

import "github.com/go-playground/validator/v10"

type customValidator struct {
    Validator *validator.Validate
}

type CustomValidator interface {
    Validate(i interface{}) (err error)
}

func NewCustomValidator(v *validator.Validate) CustomValidator {
    return &customValidator{Validator: v}
}

func (cv *customValidator) Validate(i interface{}) (err error) {
    return cv.Validator.Struct(i)
}
