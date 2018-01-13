package avoucher_test

import (
	. "go-validate/avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

var _ = Describe("Avoucher", func() {
	fmt.Println("running tests")

	It("should validate kind", func(){
		schema := NewSchema()
		isValid := schema.SetKind("").Validate("some string")
		Expect(isValid).To(Equal(true))
	})

})
