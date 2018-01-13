package avoucher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAvoucher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Avoucher Suite")
}
