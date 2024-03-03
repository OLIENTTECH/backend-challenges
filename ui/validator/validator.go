package validator

import (
	"errors"

	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/ja"
	"github.com/labstack/echo/v4"

	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
)

// CustomValidator : CustomValidator for echo
type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func NewCustomValidator() echo.Validator {
	japanese := ja_JP.New()
	uni := ut.New(japanese, japanese)
	trans, _ := uni.GetTranslator("ja")
	validate := validator.New()
	if err := ja.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(err)
	}

	return &CustomValidator{
		validator: validate,
		trans:     trans,
	}
}

// Validate : the implementation of echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			translatedMessage := validationErrs[0].Translate(cv.trans)

			return cerror.New("validetor: "+translatedMessage, cerror.WithInvalidArgumentCode())
		}

		return err
	}

	return nil
}
