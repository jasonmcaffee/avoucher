package avoucher

import "reflect"

func Validate(schema Schema, objToValidate interface{}) bool{
	isValid := validateKind(schema, objToValidate)
	if !isValid {
		return isValid
	}

	return isValid
}

func validateKind(schema Schema, objToValidate interface{}) bool{
	schemaType := reflect.TypeOf(schema.Kind)
	objType := reflect.TypeOf(objToValidate)
	isSameKind := schemaType == objType
	return isSameKind
}
