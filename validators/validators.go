package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var translator ut.Translator

func init() {
    validate = validator.New()
}

func ValidateStruct(s interface{}) error {
    return validate.Struct(s)
}

func SetCustomErrorMessages() {
    validate.RegisterTranslation("required", translator, registrationFunc("required", "{0} is required"), translationFunc)
    validate.RegisterTranslation("email", translator, registrationFunc("email", "{0} must be a valid email address"), translationFunc)
}

func translationFunc(ut ut.Translator, fe validator.FieldError) string {
    t, _ := ut.T(fe.Tag(), fe.Field())
    return t
}

func registrationFunc(tag, translation string) validator.RegisterTranslationsFunc {
    return func(ut ut.Translator) error {
        return ut.Add(tag, translation, true)
    }
}
