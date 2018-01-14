package validation_plugins_test

import (
	. "avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate Type", func() {

	It("should validate when no keys are set", func() {
		schema := Avoucher()
		validationResult := schema.Validate(123)
		Expect(validationResult.IsValid()).To(Equal(true))
	})
})
