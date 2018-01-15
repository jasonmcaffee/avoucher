package validation_plugins_test

import (
	. "avoucher"
	. "avoucher/interfaces"
	. "avoucher/models"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Validate Type", func() {

	goTypeInstances := NewGoTypeInstances()
	gti := goTypeInstances

	//helper function to assert that when a schema type is set, all other types should be invalid
	//calls IsValid on all type instances defined in GoTypeInstances.GetAllTypeInstances
	//returns true when IsValid returns false for all types other than the schema's Type
	//maxMatches - some types are aliases for other types, and therefore will be considered the same kind. e.g. byte == uint8, byte == byte, so 2 matches are allowed
	allOtherTypesAreInvalid := func(schema Schema, maxMatches int) bool{
		reflectedSchemaType := schema.GetTypeReflectedType()
		isSchemaTypePointer := reflectedSchemaType.Kind() == reflect.Ptr
		matchCount := 0
		allTypeInstances := gti.GetAllTypeInstances()
		//fmt.Println("allTypeInstances len", len(allTypeInstances))
		for _, t := range allTypeInstances {
			//don't check IsValid() when the t is the same as the schema type (should be valid in this case, but invalid in all other cases)
			reflectedTType := t.Type()
			//fmt.Println("reflected type is ", reflectedTType)
			isTPointer := reflectedTType.Kind() == reflect.Ptr

			if reflectedTType == reflectedSchemaType && isSchemaTypePointer == isTPointer{
				//fmt.Println(fmt.Sprintf("schema type %v matches type %v", reflectedSchemaType, reflectedTType))
				//fmt.Println(fmt.Sprintf("schema type is pointer: %v  type is pointer: %v", isSchemaTypePointer, isTPointer))
				matchCount++
				if matchCount > maxMatches {
					return false
				}
				continue
			}
			if schema.Validate(t).IsValid(){
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
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Int8 function which helps strictly validate that objectToValidate type is int8", func() {
			s := Avoucher().Int8()
			Expect(s.Validate(gti.Int8).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Int16 function which helps strictly validate that objectToValidate type is int16", func() {
			s := Avoucher().Int16()
			Expect(s.Validate(gti.Int16).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Int32 function which helps strictly validate that objectToValidate type is int32", func() {
			s := Avoucher().Int32()
			Expect(s.Validate(gti.Int32).IsValid()).To(Equal(true))
			Expect(s.Validate(gti.Rune).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 2)).To(Equal(true)) //rune == int32 so 2 matches allowed.
		})

		It("should provide Int64 function which helps strictly validate that objectToValidate type is int64", func() {
			s := Avoucher().Int64()
			Expect(s.Validate(gti.Int64).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		//uints
		It("should provide Uint function which helps strictly validate that objectToValidate type is uint", func() {
			s := Avoucher().Uint()
			Expect(s.Validate(gti.Uint).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Uint8 function which helps strictly validate that objectToValidate type is uint8", func() {
			s := Avoucher().Uint8()
			Expect(s.Validate(gti.Uint8).IsValid()).To(Equal(true))
			Expect(s.Validate(gti.Byte).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 2)).To(Equal(true))
		})

		It("should provide Uint16 function which helps strictly validate that objectToValidate type is uint16", func() {
			s := Avoucher().Uint16()
			Expect(s.Validate(gti.Uint16).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Uint32 function which helps strictly validate that objectToValidate type is uint32", func() {
			s := Avoucher().Uint32()
			Expect(s.Validate(gti.Uint32).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Uint64 function which helps strictly validate that objectToValidate type is uint64", func() {
			s := Avoucher().Uint64()
			Expect(s.Validate(gti.Uint64).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		//other numeric
		It("should provide Float32 function which helps strictly validate that objectToValidate type is float32", func() {
			s := Avoucher().Float32()
			Expect(s.Validate(gti.Float32).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Float64 function which helps strictly validate that objectToValidate type is float64", func() {
			s := Avoucher().Float64()
			Expect(s.Validate(gti.Float64).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Complex64 function which helps strictly validate that objectToValidate type is complex64", func() {
			s := Avoucher().Complex64()
			Expect(s.Validate(gti.Complex64).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Complex128 function which helps strictly validate that objectToValidate type is complex128", func() {
			s := Avoucher().Complex128()
			Expect(s.Validate(gti.Complex128).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 1)).To(Equal(true))
		})

		It("should provide Byte function which helps strictly validate that objectToValidate type is byte", func() {
			s := Avoucher().Byte()
			Expect(s.Validate(gti.Byte).IsValid()).To(Equal(true))
			Expect(s.Validate(gti.Uint8).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 2)).To(Equal(true)) //byte == uint8 and byte == byte
		})

		It("should provide Rune function which helps strictly validate that objectToValidate type is rune", func() {
			s := Avoucher().Rune()
			Expect(s.Validate(gti.Rune).IsValid()).To(Equal(true))
			Expect(s.Validate(gti.Int32).IsValid()).To(Equal(true))
			Expect(allOtherTypesAreInvalid(s, 2)).To(Equal(true)) //rune == int32 and rune == rune
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

		It("should consider pointers of type to be valid", func(){
			type Person struct {
				Name string
			}
			s := Avoucher().Type(&Person{})
			objToValidate := &Person{Name:"Jason"}
			Expect(s.Validate(objToValidate).IsValid()).To(Equal(true))
		})

		It("should not consider a type and a pointer of a type to be valid", func(){
			type Person struct {
				Name string
			}
			s := Avoucher().Type(Person{})
			objToValidate := &Person{Name: "Jason"}
			Expect(s.Validate(objToValidate).IsValid()).To(Equal(false))
		})

		It("should not consider a pointer of a type and a type to be valid", func(){
			type Person struct {
				Name string
			}
			s := Avoucher().Type(&Person{})
			objToValidate := Person{Name: "Jason"}
			Expect(s.Validate(objToValidate).IsValid()).To(Equal(false))
		})

		It("should not consider different types to be valid", func(){
			type Person struct {
				Name string
			}
			type Animal struct {
				Species string
			}
			s := Avoucher().Type(Person{})
			objToValidate := Animal{Species:"Dog"}
			Expect(s.Validate(objToValidate).IsValid()).To(Equal(false))
		})

		It("should not consider pointers of different types to be valid", func(){
			type Person struct {
				Name string
			}
			type Animal struct {
				Species string
			}
			s := Avoucher().Type(&Person{})
			objToValidate := &Animal{Species:"Dog"}
			Expect(s.Validate(objToValidate).IsValid()).To(Equal(false))
		})

	})



})
