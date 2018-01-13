package avoucher

func Validate(schema Schema, objToValidate interface{}) bool{
	isValid := validateKind(schema, objToValidate)
	if !isValid {
		return isValid
	}

	return isValid
}

func validateKind(schema Schema, objToValidate interface{}) bool{
	return true
}
