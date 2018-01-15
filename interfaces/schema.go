package interfaces

import "reflect"

//Schema defines criteria used during validation
//Numeric Types - https://golang.org/ref/spec#Numeric_types
//uint8       the set of all unsigned  8-bit integers (0 to 255)
//uint16      the set of all unsigned 16-bit integers (0 to 65535)
//uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
//uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
//
//int8        the set of all signed  8-bit integers (-128 to 127)
//int16       the set of all signed 16-bit integers (-32768 to 32767)
//int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
//int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
//
//float32     the set of all IEEE-754 32-bit floating-point numbers
//float64     the set of all IEEE-754 64-bit floating-point numbers
//
//complex64   the set of all complex numbers with float32 real and imaginary parts
//complex128  the set of all complex numbers with float64 real and imaginary parts
//
//byte        alias for uint8
//rune        alias for int32
//uint     		either 32 or 64 bits
//int      		same size as uint
type Schema interface {
	//SetType shorcuts
	Type(interface{}) Schema

	//Numeric Types
	Uint() Schema
	Uint8() Schema
	Uint16() Schema
	Uint32() Schema
	Uint64() Schema
	Int() Schema
	Int8() Schema
	Int16() Schema
	Int32() Schema
	Int64() Schema
	Float32() Schema
	Float64() Schema
	Complex64() Schema
	Complex128() Schema
	Byte() Schema
	Rune() Schema

	UintPointer() Schema
	Uint8Pointer() Schema
	Uint16Pointer() Schema
	Uint32Pointer() Schema
	Uint64Pointer() Schema
	IntPointer() Schema
	Int8Pointer() Schema
	Int16Pointer() Schema
	Int32Pointer() Schema
	Int64Pointer() Schema
	Float32Pointer() Schema
	Float64Pointer() Schema
	Complex64Pointer() Schema
	Complex128Pointer() Schema
	BytePointer() Schema
	RunePointer() Schema

	//Other types
	String() Schema
	StringPointer() Schema

	Slice() Schema

	//SetType is used to confirm that validated objects are of the same Type
	SetType(interface{}) Schema
	GetType() interface{}
	//avoid having to use reflection to see if the interface has a value.
	GetIsTypeSet() bool
	//performance optimization. reflect only once when SetType is called.
	GetTypeReflectedType() reflect.Type

	SetKind(reflect.Kind) Schema
	GetKind() reflect.Kind
	GetIsKindSet() bool

	//Keys are used to confirm that validated objects contain properties that conform to Schemas
	Keys(map[string]Schema) Schema
	SetKeys(map[string]Schema) Schema
	GetKeys()map[string]Schema

	//Validate is used to confirm whether validated objects meet criteria set in the schema.
	Validate(objToValidate interface{}) ValidationResult
}
