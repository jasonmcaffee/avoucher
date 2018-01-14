package models

import . "avoucher/interfaces"

type validationResult struct {
	Valid       bool
	Message       *string
	ExpectedValue interface{}
	ActualValue   interface{}
}

func NewValidationResult() ValidationResult {
	return &validationResult{}
}

//GetMessage returns an empty string if Message is nil. Returns Message value otherwise.
func (v *validationResult) GetMessage() string {
	message := ""
	if v.Message != nil {
		message = *v.Message
	}
	return message
}

//SetMessage sets the Message as pointer to message param
func (v *validationResult) SetMessage(message string) {
	v.Message = &message
}

func (v *validationResult) IsValid() bool{
	return v.Valid
}

func (v *validationResult) SetIsValid(isValid bool){
	v.Valid = isValid
}