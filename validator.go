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

func (v *ValidationResult) GetMessage() string{
	message := ""
	if v.Message != nil{
		message = *v.Message
	}
	return message
}

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
//Validation is stopped when
func Validate(schema *Schema, objToValidate interface{}) ValidationResult{
	validationResult := ValidateKind(schema, objToValidate)
	return validationResult
}

func ValidateUsingValidationPlugins(schema *Schema, objToValidate interface{}, validationPlugins []ValidationPlugin) ValidationResult{
 return ValidationResult{}
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
