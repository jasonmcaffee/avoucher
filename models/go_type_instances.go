package models

import (
	"reflect"
)


//GoTypeInstances is used to retrieve instances of every primitive type in Go.
//Useful during testing, but may have other applications.
type GoTypeInstances struct {
	AllTypeInstances []reflect.Value
	*GoNumericTypeInstances
}

//GetAllTypeInstances returns a slice containing an instance of every primitive type in Go.
//Useful when validating that objectToValidate interface{} is of a given type.
func (g *GoTypeInstances) GetAllTypeInstances() []reflect.Value {
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
	AllNumericTypeInstances []reflect.Value
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

	IntPointer   *int
	Int8Pointer  *int8
	Int16Pointer *int16
	Int32Pointer *int32
	Int64Pointer *int64

	UintPointer   *uint
	Uint8Pointer  *uint8
	Uint16Pointer *uint16
	Uint32Pointer *uint32
	Uint64Pointer *uint64

	Float32Pointer *float32
	Float64Pointer *float64

	Complex64Pointer  *complex64
	Complex128Pointer *complex128

	BytePointer *byte
	RunePointer *rune
}

//GetAllNumericTypeInstances returns a slice of instance values for each property/field defined on the GoNumericTypeInstances struct.
//NOTE: NewGetAllNumericTypeInstances instantiates the AllNumericTypeInstances slice.
func (g *GoNumericTypeInstances) GetAllNumericTypeInstances() []reflect.Value {
	return g.AllNumericTypeInstances
}

//NewGoNumericTypeInstances contructs a new instance of GoNumericTypeInstances.
//Initializes the AllNumericTypeInstances slice by calling BuildSliceOfAllFieldValuesInObject for performance optimization reasons.
func NewGoNumericTypeInstances() *GoNumericTypeInstances {
	//g := GoNumericTypeInstances{}
	gnti := &GoNumericTypeInstances{}
	gnti.AllNumericTypeInstances = BuildSliceOfAllFieldValuesInObject(gnti, []string{"AllNumericTypeInstances"})
	return gnti
}

func BuildSliceOfAllFieldValuesInObject(objToReflect interface{}, ignoreFieldsWithNames []string) []reflect.Value {
	result := []reflect.Value{}
	reflectedObjectValue := reflect.ValueOf(objToReflect)
	val := reflect.Indirect(reflectedObjectValue)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		//fmt.Println("field type: ", field.Type(), " kind: ", field.Kind(), " name: ", fieldName)
		//if field.Kind() == reflect.Ptr{
		//	fmt.Println("field is pointer")
		//}

		if contains(ignoreFieldsWithNames, fieldName){
			continue
		}

		result = append(result, field)

	}
	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
