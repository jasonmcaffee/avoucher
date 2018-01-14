package models

import (
	"reflect"
	. "avoucher/interfaces"
)

type reflectedObjectToValidate struct{
	ObjectToValidate interface{}
	ReflectedType reflect.Type
}

func NewReflectedObjectToValidate(objectToValidate interface{}) ReflectedObjectToValidate{
	return &reflectedObjectToValidate{
		ObjectToValidate: objectToValidate,
		ReflectedType: reflect.TypeOf(objectToValidate),
	}
}

func (r *reflectedObjectToValidate) GetObjectToValidate() interface{}{
	return r.ObjectToValidate
}

func (r *reflectedObjectToValidate) GetReflectedType() reflect.Type{
	return r.ReflectedType
}