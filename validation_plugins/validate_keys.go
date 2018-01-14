package validation_plugins

import (
	. "avoucher/interfaces"
	. "avoucher/models"
)

//ValidateTypePlugin extends ValidationPlugin by embedding base validationPlugin struct
//This allows us to get free funcs defined in validationPlugin, which are required by the ValidationPlugin interface.
type ValidateKeysPlugin struct {
	//Get func implementations by embedding
	ValidationPlugin
}

//NewValidateKeysPlugin instantiates a pointer to the validationPlugin struct, which implements the ValidationPlugin interface.
func NewValidateKeysPlugin() ValidationPlugin{
	validationKeysPlugin := &ValidateKeysPlugin{}
	validationKeysPlugin.ValidationPlugin = NewValidationPlugin("ValidateKeys", validationKeysPlugin.ValidateKeys)
	return validationKeysPlugin
}

//ValidateKeys
func (v *ValidateKeysPlugin) ValidateKeys(schema Schema, reflectedObjectToValidate ReflectedObjectToValidate) ValidationResult {
	//use embedded CreateValidationResult to set IsValid = true, TestName = v.Name
	validationResult := v.CreateValidationResult()
	return validationResult
}
