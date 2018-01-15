package validation_plugins_test

import (
	. "avoucher"
	. "avoucher/interfaces"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate Keys", func() {

	It("should validate when no keys are set", func() {
		schema := Avoucher()
		validationResult := schema.Validate(123)
		Expect(validationResult.IsValid()).To(Equal(true))
	})

	It("should validate presence of each key", func(){
		type Person struct{
			Name string
			Age int
		}
		person := Person{Name:"Jason", Age: 38}

		schema := Avoucher().SetKeys(map[string]Schema{
			"Name" : Avoucher().SetType(""),
			"Age" : Avoucher().Int(),
		})

		validationResult := schema.Validate(person)
		Expect(validationResult.IsValid()).To(Equal(true))

		type PersonAgeIsWrongType struct{
			Name string
			Age string
		}
		personAgeIsWrongType := PersonAgeIsWrongType{Name:"Jane", Age:"21"}
		validationResult = schema.Validate(personAgeIsWrongType)
		Expect(validationResult.IsValid()).To(Equal(false))

	})

	It("should consider missing key invalid", func(){
		type PersonMissingAge struct{
			Name string
		}

		schema := Avoucher().SetKeys(map[string]Schema{
			"Name" : Avoucher().SetType(""),
			"Age" : Avoucher().Int(),
		})

		personMissingAge := PersonMissingAge{Name:"John"}
		validationResult := schema.Validate(personMissingAge)
		Expect(validationResult.IsValid()).To(Equal(false))
	})

	XIt("should work with maps", func(){
		objectToValidate := map[string]interface{}{
			"Name": "Jason",
			"Age" : 38,
		}

		personSchema := Avoucher().SetKeys(map[string]Schema{
			"Name": Avoucher().Type(""),
			"Age" : Avoucher().Int(),
		})

		Expect(personSchema.Validate(objectToValidate).IsValid()).To(Equal(true))
	})

})
