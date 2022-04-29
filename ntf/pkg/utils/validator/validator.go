package validator

import (
	"fmt"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (

	// CustomValidator struct to hold validator
	CustomValidator struct {
		Validator *validator.Validate
	}
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	id := id.New()
	uni = ut.New(id, id)

	trans, _ = uni.GetTranslator("id")

	validate = validator.New()
	id_translations.RegisterDefaultTranslations(validate, trans)
}

func (cv *CustomValidator) Validate(i interface{}) (err error) {
	err = validate.Struct(i)
	if err != nil {
		err = ValidatorError(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return
}

func ValidatorError(err error) error {
	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		return fmt.Errorf(e.Translate(trans))
	}

	return err
}
