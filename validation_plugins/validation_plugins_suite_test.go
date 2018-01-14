package validation_plugins_test

import (
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"

"testing"
)

func TestValidationPlugins(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validation Plugins Suite")
}
