package validation_plugins

import (
	. "avoucher/interfaces"
	. "avoucher/models"
	"reflect"
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

	//func getField(v *Vertex, field string) int {
	//	r := reflect.ValueOf(v)
	//	f := reflect.Indirect(r).FieldByName(field)
	//	return int(f.Int())
	//}
	reflectedObjectValue := reflect.ValueOf(reflectedObjectToValidate.GetObjectToValidate())
	reflectedObjectIndirect := reflect.Indirect(reflectedObjectValue)
	schemaKeys := schema.GetKeys()
	for keyName, schema := range schemaKeys {
		field := reflectedObjectIndirect.FieldByName(keyName)
		//fieldValue := reflect.ValueOf(field)
		if !field.IsValid() {
			validationResult.SetIsValid(false)
			return validationResult
		}
		fieldAsInterface := field.Interface()
		validationResult = schema.Validate(fieldAsInterface)
		if !validationResult.IsValid(){
			return validationResult
		}
	}

	return validationResult
}
