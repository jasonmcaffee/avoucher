package avoucher

import "reflect"

func Validate(schema *Schema, objToValidate interface{}) (bool, error){
	isValid, err := validateKind(schema, objToValidate)
	if !isValid || err != nil {
		return isValid, err
	}

	return isValid, err
}

func validateKind(schema *Schema, objToValidate interface{}) (bool, error){
	//consider valid to be true when Kind is not set
	if schema.isKindSet == false{
		return true, nil
	}
	schemaType := reflect.TypeOf(schema.Kind)
	objType := reflect.TypeOf(objToValidate)
	isSameKind := schemaType == objType
	return isSameKind, nil
}
