package models

import . "avoucher/interfaces"

//validationResult is an implementation of the ValidationResult interface
type validationResult struct {
	Valid       	bool
	Message       *string
	ExpectedValue interface{}
	ActualValue   interface{}
	TestName 			string
}

//NewValidationResult instantiates an object which implements the ValidationResult interface
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

//IsValid returns a boolean indicating whether validation succeeded or not.
func (v *validationResult) IsValid() bool{
	return v.Valid
}

//SetIsValid sets whether validation succeeded or not.
func (v *validationResult) SetIsValid(isValid bool){
	v.Valid = isValid
}

//SetTestType is used to set the name of the test that produced the ValidationResult
func (v *validationResult) SetTestName(testType string){
	v.TestName = testType
}

//GetTestType returns the name of the test that produced the ValidationResult
func (v *validationResult) GetTestName() string{
	return v.TestName
}

//SetExpectedValue sets the value that the test expected
func (v *validationResult) SetExpectedValue(expectedValue interface{}){
	v.ExpectedValue = expectedValue
}

//GetExpectedValue returns the value that the test expected
func (v *validationResult) GetExpectedValue() interface{}{
	return v.ExpectedValue
}

//SetActualValue sets the value the test encountered in the validated object
func (v *validationResult) SetActualValue(actualValue interface{}){
	v.ActualValue = actualValue
}

//GetActualValue returns the value the test encountered in the validated object
func (v *validationResult) GetActualValue() interface{}{
	return v.ActualValue
}
