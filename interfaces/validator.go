package interfaces

type Validator interface{
	Validate(reflectedSchema ReflectedSchema, objToValidate interface{}) ValidationResult
}
