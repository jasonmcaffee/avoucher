package validation_plugins

import (
	. "avoucher/interfaces"
	. "avoucher/models"
	"fmt"
)

//ValidateTypePlugin extends ValidationPlugin by embedding base validationPlugin struct
//This allows us to get free funcs defined in validationPlugin, which are required by the ValidationPlugin interface.
type ValidateKindPlugin struct {
	//Get func implementations by embedding
	ValidationPlugin
}

//NewValidateKindPlugin instantiates a pointer to the validationPlugin struct, which implements the ValidationPlugin interface.
func NewValidateKindPlugin() ValidationPlugin{
	validationKindPlugin := &ValidateKindPlugin{}
	validationKindPlugin.ValidationPlugin = NewValidationPlugin("ValidateKind", validationKindPlugin.ValidateKind)
	return validationKindPlugin
}

//ValidateType is the validation function which determines if the schema.Type is the same as the objToValidate's Type.
//ie. reflect.TypeOf(schema.GetType()) == reflect.TypeOf(objToValidate)
func (v *ValidateKindPlugin) ValidateKind(schema Schema, reflectedObjectToValidate ReflectedObjectToValidate) ValidationResult {
	//use embedded CreateValidationResult to set IsValid = true, TestName = v.Name
	validationResult := v.CreateValidationResult()

	//consider valid to be true when Kind is not set
	if schema.GetIsKindSet() == false {
		return validationResult
	}

	//get the reflected types, so reflect isn't called in each validation func
	schemaKind:= schema.GetKind()
	objKind := reflectedObjectToValidate.GetReflectedType().Kind()

	//determine if they are the same types
	isSameKind := schemaKind == objKind
	if !isSameKind {
		validationResult.SetIsValid(false)
		validationResult.SetMessage(fmt.Sprintf("Schema kind %v did not match obj kind %v", schemaKind, objKind))
	}
	return validationResult
}
