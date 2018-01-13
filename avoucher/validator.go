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
	schemaTypeKind := reflect.TypeOf(schema.Kind).Kind()
	objTypeKind := reflect.TypeOf(objToValidate).Kind()
	isSameKind := schemaTypeKind == objTypeKind
	return isSameKind
}
