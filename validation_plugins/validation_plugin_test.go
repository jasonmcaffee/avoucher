package validation_plugins_test

import (

	. "avoucher/validation_plugins"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validation Plugin", func() {

	It("should allow creation of custom validation plugins", func() {
		plugin := NewValidateTypePlugin()
		Expect(plugin.GetName()).To(Equal("ValidateType"))
	})
})
