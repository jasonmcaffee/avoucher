package avoucher_test

import (
	. "avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Avoucher", func() {
	Describe("Keys Validation", func(){

		It("should validate specified keys exist", func() {
			type Person struct{
				Name string
			}

			schema := NewSchema()
			schema.SetKeys(map[string]*Schema{
				"Name" : NewSchema().SetKind(""),
			})

			validationResult := schema.Validate(Person{})
			Expect(validationResult.IsValid).To(Equal(true))
		})

	})
})