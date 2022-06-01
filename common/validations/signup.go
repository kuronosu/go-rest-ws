package validations

import "io"

var signupValidator *Validator

func SignupValidable() Validable {
	if signupValidator == nil {
		signupValidator = NewValidator()
		signupValidator.AddValidator("email", IsRequiredValidator, IsEmailValidator)
		signupValidator.AddValidator("password", IsRequiredValidator, IsStringValidator)
	}
	return signupValidator
}

type SignupValidationResults struct {
	*ValidationResults
}

func (vr *ValidationResults) Email() string {
	return vr.Data()["email"].(string)
}

func (vr *ValidationResults) Password() string {
	return vr.Data()["email"].(string)
}

func ValidateSignup(data map[string]any) *SignupValidationResults {
	return &SignupValidationResults{SignupValidable().Validate(data)}
}

func ValidateSignupBody(body io.Reader) (*SignupValidationResults, error) {
	res, err := SignupValidable().ValidateBody(body)
	if err != nil {
		return nil, err
	}
	return &SignupValidationResults{res}, nil
}
