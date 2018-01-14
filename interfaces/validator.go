package interfaces

type Validator interface{
	Validate(schema Schema, objToValidate interface{}) ValidationResult
}
