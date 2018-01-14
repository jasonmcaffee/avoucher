package models

import (
	. "avoucher/interfaces"
)

type validator struct{
	ValidationPlugins []ValidationPlugin
}

func NewValidator(validationPlugins []ValidationPlugin) Validator{
	return &validator{
		ValidationPlugins: validationPlugins,
	}
}

//Validate validates by iterating over the defaultValidationPlugins.
//Validation is stopped as soon as a validation fails/returns false
func (v *validator) Validate(schema Schema, objToValidate interface{}) ValidationResult {
	return v.ValidateUsingValidationPlugins(schema, objToValidate, v.ValidationPlugins)
}

func (v *validator) ValidateUsingValidationPlugins(schema Schema, objToValidate interface{}, validationPlugins []ValidationPlugin) ValidationResult {
	validationResult := NewValidationResult()
	validationResult.SetIsValid(true)

	for _, validationPlugin := range validationPlugins {
		validationResult = validationPlugin(schema, objToValidate)
		if !validationResult.IsValid() {
			return validationResult
		}
	}
	return validationResult
}

//func ValidateType(schema *Schema, objToValidate interface{}) ValidationResult{
//	validationResult := ValidationResult{IsValid:true}
//	//consider valid to be true when Kind is not set
//	if schema.IsKindSet == false{
//		return validationResult
//	}
//	schemaType := reflect.TypeOf(schema.Kind)
//	objType := reflect.TypeOf(objToValidate)
//	isSameKind := schemaType == objType
//	if !isSameKind {
//		validationResult.IsValid = false
//		validationResult.SetMessage(fmt.Sprintf("Schema kind %v did not match obj type %v", schemaType, objType))
//		validationResult.ExpectedValue = schemaType
//		validationResult.ActualValue = objType
//	}
//	return validationResult
//}
