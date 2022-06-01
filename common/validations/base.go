package validations

import (
	"encoding/json"
	"io"

	"github.com/kuronosu/go-rest-ws/common"
)

type ValidatorFunc func(interface{}) error

type ValidationErrors map[string][]string

func (errors ValidationErrors) AddError(field string, err error) bool {
	if err == nil || common.IsEmpty(field) {
		return false
	}
	if _, ok := errors[field]; !ok {
		errors[field] = make([]string, 0)
	}
	errors[field] = append(errors[field], err.Error())
	return true
}

type ValidationData map[string]interface{}

type ValidationResults struct {
	Errors ValidationErrors
	data   ValidationData
}

func (vr ValidationResults) IsOk() bool {
	return len(vr.Errors) == 0
}

func (vr ValidationResults) Data() ValidationData {
	return vr.data
}

type Validable interface {
	Validate(map[string]any) *ValidationResults
	ValidateBody(body io.Reader) (*ValidationResults, error)
}

type Validator struct {
	validators map[string][]ValidatorFunc
}

func (validable *Validator) Validate(data map[string]any) *ValidationResults {
	errors := make(ValidationErrors, 0)
	for field, validators := range validable.validators {
		for _, validator := range validators {
			errors.AddError(field, validator(data[field]))
		}
	}
	return &ValidationResults{Errors: errors, data: data}
}

func (validable *Validator) ValidateBody(body io.Reader) (*ValidationResults, error) {
	var x ValidationData
	if err := json.NewDecoder(body).Decode(&x); err != nil {
		return nil, err
	}
	return validable.Validate(x), nil
}

func (validable *Validator) AddValidator(field string, validators ...ValidatorFunc) bool {
	if common.IsEmpty(field) {
		return false
	}
	if _, ok := validable.validators[field]; !ok {
		validable.validators[field] = make([]ValidatorFunc, 0)
	}
	validable.validators[field] = append(validable.validators[field], validators...)
	return true
}

func NewValidator() *Validator {
	return &Validator{validators: make(map[string][]ValidatorFunc)}
}
