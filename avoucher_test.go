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
		})

		It("should validate int", func(){

		})

	})
})
