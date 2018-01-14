package interfaces

import "reflect"

type ReflectedSchema interface {
	GetSchema() Schema
	GetReflectedSchemaGetType() reflect.Type
}
