package validation_plugins

import (
	. "avoucher/interfaces"
	. "avoucher/models"
	"fmt"
)

//ValidateTypePlugin extends ValidationPlugin by embedding base validationPlugin struct
//This allows us to get free funcs defined in validationPlugin, which are required by the ValidationPlugin interface.
type ValidateTypePlugin struct {
	//Get func implementations by embedding
	ValidationPlugin
}

//NewValidateTypePlugin instantiates a pointer to the validationPlugin struct, which implements the ValidationPlugin interface.
func NewValidateTypePlugin() ValidationPlugin{
	validationTypePlugin := &ValidateTypePlugin{}
	validationTypePlugin.ValidationPlugin = NewValidationPlugin("ValidateType", validationTypePlugin.ValidateType)
	return validationTypePlugin
}

//ValidateType is the validation function which determines if the schema.Type is the same as the objToValidate's Type.
//ie. reflect.TypeOf(schema.GetType()) == reflect.TypeOf(objToValidate)
func (v *ValidateTypePlugin) ValidateType(schema Schema, reflectedObjectToValidate ReflectedObjectToValidate) ValidationResult {
	//use embedded CreateValidationResult to set IsValid = true, TestName = v.Name
	validationResult := v.CreateValidationResult()

	//consider valid to be true when Kind is not set
	if schema.GetIsTypeSet() == false {
		return validationResult
	}

	//get the reflected types, so reflect isn't called in each validation func
	schemaType := schema.GetTypeReflectedType()
	objType := reflectedObjectToValidate.GetReflectedType()

	//determine if they are the same types
	isSameKind := schemaType == objType
	if !isSameKind {
		validationResult.SetIsValid(false)
		validationResult.SetMessage(fmt.Sprintf("Schema type %v did not match obj type %v", schemaType, objType))
	}
	return validationResult
}
