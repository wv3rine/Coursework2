package reqvalidator

import (
	"log"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("nonZalgoText", validateText); err != nil {
		log.Fatalf("Couldn't register text validator, err=%v", err)
	}
}

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		log.Println(err.Error())
		return errors.Wrap(err, "parsing error")
	}

	if err := validate.StructCtx(c.Context(), request); err != nil {
		log.Println(err.Error())
		return errors.Wrap(err, "validation error")
	}

	return nil
}

func Validate(request interface{}) error {
	if err := validate.Struct(request); err != nil {
		log.Println(err.Error())
		return errors.Wrap(err, "validation error")
	}

	return nil
}

func ReadQueryRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.QueryParser(request); err != nil {
		return err
	}
	return validate.StructCtx(c.Context(), request)
}

func validateText(fl validator.FieldLevel) bool {
	text := fl.Field().String()

	for len(text) > 0 {
		runeValue, size := utf8.DecodeRuneInString(text)

		if unicode.Is(unicode.Mn, runeValue) {
			return false
		}

		text = text[size:]
	}

	return true
}
