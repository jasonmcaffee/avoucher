package models

import (
	"reflect"
	. "avoucher/interfaces"
)

type reflectedSchema struct{
	Schema Schema
	ReflectedSchemaGetType reflect.Type
}

func NewReflectedSchema(schema Schema) ReflectedSchema{
	return &reflectedSchema{
		Schema: schema,
		ReflectedSchemaGetType: reflect.TypeOf(schema.GetType()),
	}
}

func (r *reflectedSchema) GetSchema() Schema{
	return r.Schema
}

func (r *reflectedSchema) GetReflectedSchemaGetType() reflect.Type{
	return r.ReflectedSchemaGetType
}