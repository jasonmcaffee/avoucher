package validation_plugins_test

import (
	. "avoucher"
	. "avoucher/interfaces"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

//GoNumericTypeInstances contains instances for every Numeric type defined in Go.
type GoNumericTypeInstances struct{
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
	Int int
	Int8 int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint uint
	Uint8 uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	Float32 float32
	Float64 float64

	Complex64 complex64
	Complex128 complex128

	Byte byte
	Rune rune
}

//GetAllNumericTypeInstances returns a slice of instance values for each property/field defined on the GoNumericTypeInstances struct.
//NOTE: NewGetAllNumericTypeInstances instantiates the AllNumericTypeInstances slice.
func (g *GoNumericTypeInstances) GetAllNumericTypeInstances() []interface{}{
	return g.AllNumericTypeInstances
}

//NewGoNumericTypeInstances contructs a new instance of GoNumericTypeInstances.
//Initializes the AllNumericTypeInstances slice by calling BuildSliceOfAllFieldValuesInObject for performance optimization reasons.
func NewGoNumericTypeInstances() *GoNumericTypeInstances{
	goNumericTypeInstances := &GoNumericTypeInstances{}
	goNumericTypeInstances.AllNumericTypeInstances = BuildSliceOfAllFieldValuesInObject(goNumericTypeInstances)
	return goNumericTypeInstances
}

//GoTypeInstances is used to retrieve instances of every primitive type in Go.
//Useful during testing, but may have other applications.
type GoTypeInstances struct{
	AllTypeInstances []interface{}
	*GoNumericTypeInstances
}

//GetAllTypeInstances returns a slice containing an instance of every primitive type in Go.
//Useful when validating that objectToValidate interface{} is of a given type.
func (g *GoTypeInstances) GetAllTypeInstances() []interface{}{
	return g.AllTypeInstances
}

//NewGoTypeInstances creates a new instance of GoTypeInstances
//initializes slice of AllTypeInstances by appending Numeric type instances, etc together.
func NewGoTypeInstances() *GoTypeInstances{
	gti := &GoTypeInstances{
		GoNumericTypeInstances: NewGoNumericTypeInstances(),
	}
	gti.AllTypeInstances = append(gti.AllTypeInstances, gti.GetAllNumericTypeInstances()...)
	return gti
}

//helper to reflectively iterate over each Field in a given object, retrieve its value, and return slice of all field values.
func BuildSliceOfAllFieldValuesInObject(objToReflect interface{})[]interface{}{
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


var _ = Describe("Validate Type", func() {

	goTypeInstances := NewGoTypeInstances()
	gti := goTypeInstances

	//helper function to assert that when a schema type is set, all other types should be invalid
	//calls IsValid on all type instances defined in GoTypeInstances.GetAllTypeInstances
	//returns true when IsValid returns false for all types other than the schema's Type
	allOtherTypesAreInvalid := func(schema Schema) bool{
		reflectedSchemaType := schema.GetTypeReflectedType()
		matchCount := 0
		for _, t := range goTypeInstances.GetAllTypeInstances() {
			//don't check IsValid() when the t is the same as the schema type (should be valid in this case, but invalid in all other cases)
			reflectedTType := reflect.TypeOf(t)
			if reflectedTType == reflectedSchemaType{
				matchCount++
				continue
			}
			if matchCount > 1 || schema.Validate(t).IsValid(){
				return false
			}
		}
		return true
	}

	Describe("Numeric Type Validation", func(){
		//ints
		It("should provide Int function which helps strictly validate that objectToValidate type is int", func() {
			s := Avoucher().Int()
			Expect(s.Validate(gti.Int).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s)).To(Equal(true))
		})

		It("should provide Int8 function which helps strictly validate that objectToValidate type is int8", func() {
			s := Avoucher().Int8()
			Expect(s.Validate(gti.Int8).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s)).To(Equal(true))
		})

		It("should provide Int16 function which helps strictly validate that objectToValidate type is int16", func() {
			s := Avoucher().Int16()
			Expect(s.Validate(gti.Int16).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s)).To(Equal(true))
		})

		It("should provide Int32 function which helps strictly validate that objectToValidate type is int32", func() {
			s := Avoucher().Int32()
			Expect(s.Validate(gti.Int32).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s)).To(Equal(true))
		})

		It("should provide Int64 function which helps strictly validate that objectToValidate type is int64", func() {
			s := Avoucher().Int64()
			Expect(s.Validate(gti.Int64).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s)).To(Equal(true))
		})

		//uints
		It("should provide Uint function which helps strictly validate that objectToValidate type is uint", func() {
			s := Avoucher().Uint()
			Expect(s.Validate(gti.Uint).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s)).To(Equal(true))
		})

	})

	Describe("Misc Type Validation", func(){
		It("should validate when no kind is set", func() {
			schema := Avoucher()
			validationResult := schema.Validate(123)
			Expect(validationResult.IsValid()).To(Equal(true))
		})

		It("should validate string types", func() {
			schema := Avoucher()
			validationResult := schema.SetType("").Validate("some string")
			Expect(validationResult.IsValid()).To(Equal(true))

			validationResult = schema.Validate(123)
			Expect(validationResult.IsValid()).To(Equal(false))
			fmt.Println(validationResult.GetMessage())
		})
	})



	Describe("Custom Type Validation", func(){
		It("should validate custom struct types", func() {
			type Person struct {
				Name string
			}
			type Animal struct {
				Species string
			}
			schema := Avoucher()
			validationResult := schema.SetType(Person{}).Validate(Person{Name: "Jason"})
			Expect(validationResult.IsValid()).To(Equal(true))

			validationResult = schema.Validate(Animal{Species: "Lion"})
			Expect(validationResult.IsValid()).To(Equal(false))
			fmt.Println(validationResult.GetMessage())
		})

		It("should validate pointers to custom struct types", func() {
			type Person struct {
				Name string
			}
			type Animal struct {
				Species string
			}
			schema := Avoucher()
			validationResult := schema.SetType(&Person{}).Validate(&Person{Name: "Jason"})
			Expect(validationResult.IsValid()).To(Equal(true))

			validationResult = schema.SetType(Person{}).Validate(&Person{Name: "Jason"})
			Expect(validationResult.IsValid()).To(Equal(false))
			fmt.Println(validationResult.GetMessage())

			validationResult = schema.Validate(&Animal{Species: "Lion"})
			Expect(validationResult.IsValid()).To(Equal(false))
			fmt.Println(validationResult.GetMessage())
		})
	})



})
