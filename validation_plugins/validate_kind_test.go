package validation_plugins_test

import (
	. "avoucher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Validate Kind", func() {
	It("should provide SetKind function for setting the reflect.Kind the objectToValidate should be", func(){
		s := Avoucher().SetKind(reflect.Slice)
		objectToValidate := []string{}
		Expect(s.Validate(objectToValidate).IsValid()).To(Equal(true))
	})

	It("should provide a Slice function shortcut to setting Kind", func(){
		s := Avoucher().Slice()
		objectsToValidate := []interface{}{
			[]string{},
			&[]string{},
			[]int{},
			&[]int{},
			[]chan int{},
			[]func(){},
		}
		for _, objectToValidate := range objectsToValidate{
			Expect(s.Validate(objectToValidate).IsValid()).To(Equal(true))
		}
	})
})
