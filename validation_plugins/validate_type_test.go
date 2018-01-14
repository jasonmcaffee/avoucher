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
