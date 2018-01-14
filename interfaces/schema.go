package interfaces

import "reflect"

//Schema defines criteria used during validation
type Schema interface {
	//SetType is used to confirm that validated objects are of the same Type
	SetType(interface{}) Schema
	GetType() interface{}
	//avoid having to use reflection to see if the interface has a value.
	GetIsTypeSet() bool
	//performance optimization. reflect only once when SetType is called.
	GetTypeReflectedType() reflect.Type

	//Keys are used to confirm that validated objects contain properties that conform to Schemas
	SetKeys(map[string]Schema) Schema
	GetKeys()map[string]Schema

	//Validate is used to confirm whether validated objects meet criteria set in the schema.
	Validate(objToValidate interface{}) ValidationResult
}
