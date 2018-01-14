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
	reflectedObjectToValidate := NewReflectedObjectToValidate(objToValidate)
	validationResult := NewValidationResult()
	validationResult.SetIsValid(true)
	validationResult.SetTestName("None")

	for _, validationPlugin := range validationPlugins {
		validationResult = validationPlugin.GetValidationFunc()(schema, reflectedObjectToValidate)
		if !validationResult.IsValid() {
			return validationResult
		}
	}
	return validationResult
}
