package interfaces

//Schema defines criteria used during validation
type Schema interface {
	//SetType is used to confirm that validated objects are of the same Type
	SetType(interface{}) Schema
	GetType() interface{}
	GetIsTypeSet() bool

	//Keys are used to confirm that validated objects contain properties that conform to Schemas
	SetKeys(map[string]Schema) Schema
	GetKeys()map[string]Schema

	//Validate is used to confirm whether validated objects meet criteria
	Validate(objToValidate interface{}) ValidationResult
}
