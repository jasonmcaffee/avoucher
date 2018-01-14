package interfaces

//ValidationPlugin func signature for validation.
type ValidationFunc func(schema Schema, objToValidate interface{}) ValidationResult

type ValidationPlugin interface {
	GetValidationFunc() ValidationFunc
	GetName() string
}