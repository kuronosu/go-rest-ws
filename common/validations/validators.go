package validations

import (
	"errors"
	"net/mail"
	"reflect"
)

var (
	_STRING = reflect.TypeOf("")
	_INT    = reflect.TypeOf(0)
	_FLOAT  = reflect.TypeOf(0.0)
)

func IsRequiredValidator(data interface{}) error {
	if data == nil {
		return errors.New("is required")
	}
	return nil
}

func IsTypeValidator(data interface{}, typ reflect.Type) error {
	if data == nil {
		return nil
	}
	if reflect.TypeOf(data) != typ {
		return errors.New("is not a " + typ.Name())
	}
	return nil
}

func IsStringValidator(data interface{}) error {
	return IsTypeValidator(data, _STRING)
}

func IsIntValidator(data interface{}) error {
	return IsTypeValidator(data, _INT)
}

func IsFloat64Validator(data interface{}) error {
	return IsTypeValidator(data, _FLOAT)
}

func IsEmailValidator(data interface{}) error {
	err := IsStringValidator(data)
	if err != nil {
		return err
	}
	_, err = mail.ParseAddress(data.(string))
	return err
}
