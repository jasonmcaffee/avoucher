package interfaces

type Schema interface {
	GetIsTypeSet() bool
	SetType(interface{}) Schema
	GetType() interface{}

	SetKeys(map[string]Schema) Schema
	GetKeys()map[string]Schema

	Validate(objToValidate interface{}) ValidationResult
}
