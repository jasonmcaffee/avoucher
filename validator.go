package avoucher

import (
	"reflect"
	"fmt"
)

type ValidationResult struct{
	IsValid bool
	Message *string
	ExpectedValue interface{}
	ActualValue interface{}
}

//GetMessage returns an empty string if Message is nil. Returns Message value otherwise.
func (v *ValidationResult) GetMessage() string{
	message := ""
	if v.Message != nil{
		message = *v.Message
	}
	return message
}

//SetMessage sets the Message as pointer to message param
func (v *ValidationResult) SetMessage(message string){
	v.Message = &message
}

//ValidationPlugin func signature for validation.
type ValidationPlugin func(schema *Schema, objToValidate interface{}) ValidationResult

//DefaultValidationPlugins is a slice of validations to be performed by default.
var DefaultValidationPlugins []ValidationPlugin = []ValidationPlugin{
	ValidateKind,
}

//Validate validates by iterating over the defaultValidationPlugins.
//Validation is stopped as soon as a validation fails/returns false
func Validate(schema *Schema, objToValidate interface{}) ValidationResult{
	return ValidateUsingValidationPlugins(schema, objToValidate, DefaultValidationPlugins)
}

func ValidateUsingValidationPlugins(schema *Schema, objToValidate interface{}, validationPlugins []ValidationPlugin) ValidationResult{
	validationResult := ValidationResult{IsValid:true}
	for _, validationPlugin := range validationPlugins{
		validationResult = validationPlugin(schema, objToValidate)
		if !validationResult.IsValid{
			return validationResult
		}
	}
	return validationResult
}

func ValidateKind(schema *Schema, objToValidate interface{}) ValidationResult{
	validationResult := ValidationResult{IsValid:true}
	//consider valid to be true when Kind is not set
	if schema.isKindSet == false{
		return validationResult
	}
	schemaType := reflect.TypeOf(schema.Kind)
	objType := reflect.TypeOf(objToValidate)
	isSameKind := schemaType == objType
	if !isSameKind {
		validationResult.IsValid = false
		validationResult.SetMessage(fmt.Sprintf("Schema kind %v did not match obj type %v", schemaType, objType))
		validationResult.ExpectedValue = schemaType
		validationResult.ActualValue = objType
	}
	return validationResult
}
