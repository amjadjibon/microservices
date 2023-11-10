package error_parser

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

func ParseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
			// loop through each of those and convert them one by one
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
			out = append(out, parseMarshallingError(*typedError))
		default:
			if err.Error() == "EOF" {
				out = append(out, "Body is empty")
				continue
			}
			out = append(out, err.Error())
		}
	}
	return out
}

func parseFieldError(e validator.FieldError) string {
	field := e.Field()
	tag := strings.Split(e.Tag(), "|")[0]

	switch tag {
	case "required":
		return fmt.Sprintf("The field '%s' is required", field)
	case "required_without":
		return fmt.Sprintf("The field '%s' is required if '%s' is not supplied", field, e.Param())
	case "lt", "ltfield", "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		comparison := "less than"
		if strings.HasPrefix(tag, "gt") {
			comparison = "greater than"
		}
		return fmt.Sprintf("The field '%s' must be %s '%s'", field, comparison, param)
	default:
		english := en.New()
		translator := ut.New(english, english)
		translatorInstance, found := translator.GetTranslator("en")
		if !found {
			return fmt.Errorf("%v", e).Error()
		}
		return e.Translate(translatorInstance)
	}
}

func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}
