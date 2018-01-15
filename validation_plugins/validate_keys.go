package validation_plugins

import (
	. "avoucher/interfaces"
	. "avoucher/models"
	"reflect"
	"fmt"
)

//ValidateTypePlugin extends ValidationPlugin by embedding base validationPlugin struct
//This allows us to get free funcs defined in validationPlugin, which are required by the ValidationPlugin interface.
type ValidateKeysPlugin struct {
	//Get func implementations defined in validationPlugin by embedding
	ValidationPlugin
}

//NewValidateKeysPlugin instantiates a pointer to the validationPlugin struct, which implements the ValidationPlugin interface.
func NewValidateKeysPlugin() ValidationPlugin{
	validationKeysPlugin := &ValidateKeysPlugin{}
	validationKeysPlugin.ValidationPlugin = NewValidationPlugin("ValidateKeys", validationKeysPlugin.ValidateKeys)
	return validationKeysPlugin
}

//ValidateKeys iterates over each Key (property name + schema) defined in the schema.
//Ensures the property name exists on the objectToValidate.
//Ensures the property value meets Schema criteria.
//Stops validating when first invalid property is encountered.
//TODO: provide option to continue validating when invalid property is encountered.
//may be useful https://gist.github.com/hvoecking/10772475#file-translate-go-L191
func (v *ValidateKeysPlugin) ValidateKeys(schema Schema, reflectedObjectToValidate ReflectedObjectToValidate) ValidationResult {
	//use embedded CreateValidationResult to set IsValid = true, TestName = v.Name
	validationResult := v.CreateValidationResult()

	if len(schema.GetKeys()) <= 0{
		return validationResult
	}

	//get reflection objects needed in order to access Fields
	reflectedObjectValue := reflectedObjectToValidate.GetReflectedValue()
	//reflectedObjectIndirect := reflect.Indirect(reflectedObjectValue)

	fmt.Println("Validating keys on type: ", reflectedObjectValue.Type(),  " kind: ", reflectedObjectValue.Kind())
	switch reflectedObjectValue.Kind(){
	case reflect.Struct:
		validationResult = validateStruct(schema, reflectedObjectToValidate, validationResult)
	default:
		validationResult = validateStruct(schema, reflectedObjectToValidate, validationResult)
		//validationResult.SetIsValid(false)
		//validationResult.SetMessage(fmt.Sprintf("The type: %v is not supported by ValidateKeys plugin", reflectedObjectValue.Type()))
	}

	return validationResult
}

func validateStruct(schema Schema, reflectedObjectToValidate ReflectedObjectToValidate, validationResult ValidationResult) ValidationResult{
	//get reflection objects needed in order to access Fields
	reflectedObjectValue := reflectedObjectToValidate.GetReflectedValue()
	reflectedObjectIndirect := reflect.Indirect(reflectedObjectValue)

	//iterate over each Key defined in the schema and ensure it exists on the objectToValidate
	schemaKeys := schema.GetKeys()
	for keyName, schema := range schemaKeys {
		//get the field by name
		field := reflectedObjectIndirect.FieldByName(keyName)

		//IsValid == false when the field doesn't exist
		if !field.IsValid() {
			validationResult.SetIsValid(false)
			return validationResult
		}

		//cast the field as an interface, then pass it to the Schema assigned to the key
		fieldAsInterface := field.Interface()
		validationResult = schema.Validate(fieldAsInterface)
		if !validationResult.IsValid(){
			return validationResult
		}
	}

	return validationResult
}