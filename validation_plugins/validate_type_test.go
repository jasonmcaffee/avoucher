package validation_plugins_test

import (
	. "avoucher"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate Type", func() {

	It("should validate when no kind is set", func() {
		schema := Avoucher()
		validationResult := schema.Validate(123)
		Expect(validationResult.IsValid()).To(Equal(true))
	})

	It("should validate string", func() {
		schema := Avoucher()
		validationResult := schema.SetType("").Validate("some string")
		Expect(validationResult.IsValid()).To(Equal(true))

		validationResult = schema.Validate(123)
		Expect(validationResult.IsValid()).To(Equal(false))
		fmt.Println(validationResult.GetMessage())
	})

	It("should validate int", func() {
		schema := Avoucher()
		validationResult := schema.SetType(123).Validate(43)
		Expect(validationResult.IsValid()).To(Equal(true))

		validationResult = schema.Validate("asdf")
		Expect(validationResult.IsValid()).To(Equal(false))
		fmt.Println(validationResult.GetMessage())
	})

	It("should validate custom structs", func() {
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

	It("should validate pointers to custom structs", func() {
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
