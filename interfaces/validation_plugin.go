package interfaces

//ValidationPlugin func signature for validation.
type ValidationPlugin func(schema Schema, objToValidate interface{}) ValidationResult
