package avoucher_test

import (
	. "avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

var _ = Describe("Avoucher", func() {
	fmt.Println("running tests")

	Describe("Kind Validation", func(){
		It("should validate string", func(){
			schema := NewSchema()
			isValid := schema.SetKind("").Validate("some string")
			Expect(isValid).To(Equal(true))

			isValid = schema.Validate(123)
			Expect(isValid).To(Equal(false))
		})

		It("should validate int", func(){
			schema := NewSchema()
			isValid := schema.SetKind(123).Validate(43)
			Expect(isValid).To(Equal(true))

			isValid = schema.Validate("asdf")
			Expect(isValid).To(Equal(false))
		})

	})
})
