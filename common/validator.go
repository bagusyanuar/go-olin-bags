package common

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ValidationMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func TranslateError(err error, req any) []*ValidationMessage {
	var results []*ValidationMessage
	validationError := err.(validator.ValidationErrors)
	for _, ve := range validationError {
		result := &ValidationMessage{
			Field:   getJSONField(ve.Field(), req),
			Message: customMessage(ve.Tag()),
		}
		results = append(results, result)
	}
	return results
}

func getJSONField(field string, req any) string {
	f, _ := reflect.TypeOf(req).Elem().FieldByName(field)
	value, _ := f.Tag.Lookup("json")
	return value
}

func customMessage(tag string) string {
	switch tag {
	case "required":
		return "Field is required"
	case "len":
		return "Field must be ... character"
	case "uuid4":
		return "Field must be uuid string type"
	case "e164":
		return "Field must be formatted phone number +[code][subscriber number]"
	}
	return "unknown error"
}
