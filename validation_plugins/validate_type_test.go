package validation_plugins_test

import (
	. "avoucher"
	. "avoucher/interfaces"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

type GoNumericTypeInstances struct{
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

	Uint uint
}

func (g *GoNumericTypeInstances) GetAllNumericTypeInstances() []interface{}{
	return []interface{}{
		g.Int,
		g.Int8,
		g.Uint,
	}
}

type GoTypeInstances struct{
	*GoNumericTypeInstances
}

func (g *GoTypeInstances) GetAllTypeInstances() []interface{}{
	return g.GetAllNumericTypeInstances()
}

func NewGoTypeInstances() *GoTypeInstances{
	goTypeInstances := &GoTypeInstances{
		&GoNumericTypeInstances{},
	}
	return goTypeInstances
}


var _ = Describe("Validate Type", func() {

	goTypeInstances := NewGoTypeInstances()
	gti := goTypeInstances

	//helper function to assert that when a schema type is set, all other types should be invalid
	//calls IsValid on all type instances defined in GoTypeInstances.GetAllTypeInstances
	//returns true when IsValid returns false for all types other than the schema's Type
	otherTypesAreInvalid := func(schema Schema) bool{
		reflectedSchemaType := schema.GetTypeReflectedType()
		for _, t := range goTypeInstances.GetAllTypeInstances() {
			//don't check IsValid() when the t is the same as the schema type (should be valid in this case, but invalid in all other cases)
			reflectedTType := reflect.TypeOf(t)
			if reflectedTType == reflectedSchemaType{
				continue
			}
			if schema.Validate(t).IsValid(){
				return false
			}
		}
		return true
	}

	Describe("Numeric Type Validation", func(){
		It("should provide TypeInt function which strictly validates type is int", func() {
			s := Avoucher().Int()
			Expect(s.Validate(gti.Int).IsValid()).To(Equal(true))
			Expect(otherTypesAreInvalid(s)).To(Equal(true))
		})

		It("should provide TypeInt function which strictly validates type is int", func() {
			s := Avoucher().Uint()
			Expect(s.Validate(gti.Uint).IsValid()).To(Equal(true))
			Expect(otherTypesAreInvalid(s)).To(Equal(true))
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
