package common

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validator[T any](c *fiber.Ctx) (*T, error) {
	body := new(T)

	if err := c.BodyParser(body); err != nil {
		return nil, errors.New("Invalid req")
	}

	var validate = validator.New(validator.WithRequiredStructEnabled())
	if errs := validate.Struct(body); errs != nil {
		err := errs.(validator.ValidationErrors)[0]
		switch err.Tag() {
		case "required":
			return nil, errors.New(fmt.Sprintf("Field '%s' cannot be blank", err.Field()))
		case "email":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be a valid email address", err.Field()))
		case "len":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be exactly %v characters long", err.Field(), err.Param()))
		case "min":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be at least %v characters long", err.Field(), err.Param()))
		case "max":
			return nil, errors.New(fmt.Sprintf("Field '%s' cannot be than %v characters long", err.Field(), err.Param()))
		case "alphanumunicode":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be exactly letters or numbers", err.Field()))
		case "alphaunicode":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be exactly letters", err.Field()))
		case "number":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be a number", err.Field()))
		case "gte":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be greater than or equal %v", err.Field(), err.Param()))
		case "lte":
			return nil, errors.New(fmt.Sprintf("Field '%s' must be less than or equal %v", err.Field(), err.Param()))
		default:
			return nil, errors.New(fmt.Sprintf("Field '%s': '%v' must satisfy '%s' '%v' criteria", err.Field(), err.Value(), err.Tag(), err.Param()))
		}
	}

	return body, nil
}
