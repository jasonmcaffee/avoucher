package models

import "reflect"


//GoTypeInstances is used to retrieve instances of every primitive type in Go.
//Useful during testing, but may have other applications.
type GoTypeInstances struct {
	AllTypeInstances []interface{}
	*GoNumericTypeInstances
}

//GetAllTypeInstances returns a slice containing an instance of every primitive type in Go.
//Useful when validating that objectToValidate interface{} is of a given type.
func (g *GoTypeInstances) GetAllTypeInstances() []interface{} {
	return g.AllTypeInstances
}

//NewGoTypeInstances creates a new instance of GoTypeInstances
//initializes slice of AllTypeInstances by appending Numeric type instances, etc together.
func NewGoTypeInstances() *GoTypeInstances {
	gti := &GoTypeInstances{
		GoNumericTypeInstances: NewGoNumericTypeInstances(),
	}
	gti.AllTypeInstances = append(gti.AllTypeInstances, gti.GetAllNumericTypeInstances()...)
	return gti
}


//GoNumericTypeInstances contains instances for every Numeric type defined in Go.
type GoNumericTypeInstances struct {
	//cache of all property values on this object (built when NewGoNumericTypeInstances is called)
	AllNumericTypeInstances []interface{}
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
	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint   uint
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	Float32 float32
	Float64 float64

	Complex64  complex64
	Complex128 complex128

	Byte byte
	Rune rune
}

//GetAllNumericTypeInstances returns a slice of instance values for each property/field defined on the GoNumericTypeInstances struct.
//NOTE: NewGetAllNumericTypeInstances instantiates the AllNumericTypeInstances slice.
func (g *GoNumericTypeInstances) GetAllNumericTypeInstances() []interface{} {
	return g.AllNumericTypeInstances
}

//NewGoNumericTypeInstances contructs a new instance of GoNumericTypeInstances.
//Initializes the AllNumericTypeInstances slice by calling BuildSliceOfAllFieldValuesInObject for performance optimization reasons.
func NewGoNumericTypeInstances() *GoNumericTypeInstances {
	goNumericTypeInstances := &GoNumericTypeInstances{}
	goNumericTypeInstances.AllNumericTypeInstances = BuildSliceOfAllFieldValuesInObject(goNumericTypeInstances)
	return goNumericTypeInstances
}

//helper to reflectively iterate over each Field in a given object, retrieve its value, and return slice of all field values.
func BuildSliceOfAllFieldValuesInObject(objToReflect interface{}) []interface{} {
	result := []interface{}{}
	reflectedObjectValue := reflect.ValueOf(objToReflect)
	reflectedObjectIndirect := reflect.Indirect(reflectedObjectValue)
	for i := 0; i < reflectedObjectIndirect.NumField(); i++ {
		field := reflectedObjectIndirect.Field(i)
		fieldAsInterface := field.Interface()
		result = append(result, fieldAsInterface)
	}
	return result
}
