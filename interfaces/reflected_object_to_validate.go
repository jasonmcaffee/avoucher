package interfaces

import "reflect"

//ReflectedObjectToValidate caches reflected data so reflection occurs only once instead of in each validation plugin.
type ReflectedObjectToValidate interface {
	//GetObjectToValidate returns the raw value
	GetObjectToValidate() interface{}
	//GetReflectedType returns the value of reflect.TypeOf(GetObjectToValidate())
	GetReflectedType() reflect.Type
	//GetReflectedValue returns the value of reflect.ValueOf(GetObjectToValidate())
	GetReflectedValue() reflect.Value
}