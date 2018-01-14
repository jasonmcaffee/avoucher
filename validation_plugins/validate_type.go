package validation_plugins

import (
	. "avoucher/interfaces"
	. "avoucher/models"
	"fmt"
	"reflect"
)

type ValidateTypePlugin struct {
	//Get func implementations by embedding
	ValidationPlugin
}

func NewValidateTypePlugin() ValidationPlugin{
	name := "ValidateType"
	return &ValidateTypePlugin{
		ValidationPlugin: NewValidationPlugin(name, ValidateType),
	}
}

//ValidateType
func ValidateType(schema Schema, objToValidate interface{}) ValidationResult {
	validationResult := NewValidationResult()
	validationResult.SetIsValid(true)
	//consider valid to be true when Kind is not set
	if schema.GetIsTypeSet() == false {
		return validationResult
	}
	schemaType := reflect.TypeOf(schema.GetType())
	objType := reflect.TypeOf(objToValidate)
	isSameKind := schemaType == objType
	if !isSameKind {
		validationResult.SetIsValid(false)
		validationResult.SetMessage(fmt.Sprintf("Schema kind %v did not match obj type %v", schemaType, objType))
	}
	return validationResult
}
