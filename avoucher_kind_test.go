package avoucher_test

import (
	. "avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

var _ = Describe("Avoucher", func() {

	Describe("Kind Validation", func(){
		It("should validate when no kind is set", func(){
			schema := NewSchema()
			validationResult := schema.Validate(123)
			Expect(validationResult.IsValid).To(Equal(true))
			Expect(validationResult.Message).To(BeNil())
		})

		It("should validate string", func(){
			schema := NewSchema()
			validationResult := schema.SetKind("").Validate("some string")
			Expect(validationResult.IsValid).To(Equal(true))
			Expect(validationResult.Message).To(BeNil())

			validationResult = schema.Validate(123)
			Expect(validationResult.IsValid).To(Equal(false))
			Expect(validationResult.Message).ToNot(BeNil())
			fmt.Println(validationResult.GetMessage())
		})

		It("should validate int", func(){
			schema := NewSchema()
			validationResult := schema.SetKind(123).Validate(43)
			Expect(validationResult.IsValid).To(Equal(true))
			Expect(validationResult.Message).To(BeNil())

			validationResult = schema.Validate("asdf")
			Expect(validationResult.IsValid).To(Equal(false))
			Expect(validationResult.Message).ToNot(BeNil())
			fmt.Println(validationResult.GetMessage())
		})

		It("should validate custom structs", func(){
			type Person struct{
				Name string
			}
			type Animal struct{
				Species string
			}
			schema := NewSchema()
			validationResult := schema.SetKind(Person{}).Validate(Person{Name:"Jason"})
			Expect(validationResult.IsValid).To(Equal(true))
			Expect(validationResult.Message).To(BeNil())

			validationResult = schema.Validate(Animal{Species:"Lion"})
			Expect(validationResult.IsValid).To(Equal(false))
			Expect(validationResult.Message).ToNot(BeNil())
			fmt.Println(validationResult.GetMessage())
		})

		It("should validate pointers to custom structs", func(){
			type Person struct{
				Name string
			}
			type Animal struct{
				Species string
			}
			schema := NewSchema()
			validationResult := schema.SetKind(&Person{}).Validate(&Person{Name:"Jason"})
			Expect(validationResult.IsValid).To(Equal(true))
			Expect(validationResult.Message).To(BeNil())

			validationResult = schema.SetKind(Person{}).Validate(&Person{Name:"Jason"})
			Expect(validationResult.IsValid).To(Equal(false))
			Expect(validationResult.Message).ToNot(BeNil())
			fmt.Println(validationResult.GetMessage())

			validationResult = schema.Validate(&Animal{Species:"Lion"})
			Expect(validationResult.IsValid).To(Equal(false))
			Expect(validationResult.Message).ToNot(BeNil())
			fmt.Println(validationResult.GetMessage())
		})

	})
})
