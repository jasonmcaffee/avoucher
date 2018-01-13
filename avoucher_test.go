package avoucher_test

import (
	. "avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Avoucher", func() {

	Describe("Kind Validation", func(){
		It("should validate string", func(){
			schema := NewSchema()
			isValid, err := schema.SetKind("").Validate("some string")
			Expect(isValid).To(Equal(true))
			Expect(err).To(BeNil())

			isValid, err = schema.Validate(123)
			Expect(isValid).To(Equal(false))
			Expect(err).To(BeNil())
		})

		It("should validate int", func(){
			schema := NewSchema()
			isValid, err := schema.SetKind(123).Validate(43)
			Expect(isValid).To(Equal(true))
			Expect(err).To(BeNil())

			isValid, err = schema.Validate("asdf")
			Expect(isValid).To(Equal(false))
			Expect(err).To(BeNil())
		})

		It("should validate custom structs", func(){
			type Person struct{
				Name string
			}
			type Animal struct{
				Species string
			}
			schema := NewSchema()
			isValid, err := schema.SetKind(Person{Name:"Jason"}).Validate(Person{Name:"Jason"})
			Expect(isValid).To(Equal(true))
			Expect(err).To(BeNil())

			isValid, err = schema.Validate(Animal{Species:"Lion"})
			Expect(isValid).To(Equal(false))
			Expect(err).To(BeNil())
		})

		It("should validate pointers to custom structs", func(){
			type Person struct{
				Name string
			}
			type Animal struct{
				Species string
			}
			schema := NewSchema()
			isValid, err := schema.SetKind(&Person{}).Validate(&Person{Name:"Jason"})
			Expect(isValid).To(Equal(true))
			Expect(err).To(BeNil())

			isValid, err = schema.SetKind(Person{}).Validate(&Person{Name:"Jason"})
			Expect(isValid).To(Equal(false))
			Expect(err).To(BeNil())

			isValid, err = schema.Validate(&Animal{Species:"Lion"})
			Expect(isValid).To(Equal(false))
			Expect(err).To(BeNil())
		})

	})
})
