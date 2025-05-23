package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ValidateRegister(data RegisterRequest) error {
	return validate.Struct(data)
}

func ValidateLogin(data LoginRequest) error {
	return validate.Struct(data)
}
