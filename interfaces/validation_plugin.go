package interfaces

//ValidationPlugin func signature for validation.
type ValidationFunc func(schema Schema, objToValidate interface{}) ValidationResult

type ValidationPlugin interface {
	GetValidationFunc() ValidationFunc
	//SetValidationFunc allows for a plugin's validation func to be overridden when needed.
	SetValidationFunc(ValidationFunc)

	GetName() string
	SetName(string)

	//CreateValidationResult creates a new ValidationResult.
	//Convenient for setting TestName, IsValid, etc with default values so new plugins can be created with less boilerplate code.
	CreateValidationResult() ValidationResult
}