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

	Describe("Struct types", func(){
		It("should validate presence of each key", func(){
			type Person struct{
				Name string
				Age int
			}
			person := Person{Name:"Jason", Age: 38}

			schema := Avoucher().Keys(map[string]Schema{
				"Name" : Avoucher().String(),
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

			schema := Avoucher().Keys(map[string]Schema{
				"Name" : Avoucher().String(),
				"Age" : Avoucher().Int(),
			})

			personMissingAge := PersonMissingAge{Name:"John"}
			validationResult := schema.Validate(personMissingAge)
			Expect(validationResult.IsValid()).To(Equal(false))
		})
	})


	Describe("Map types", func(){

		It("should consider map valid when it contains all keys", func(){
			objectToValidate := map[string]interface{}{
				"Name": "Jason",
				"Age" : 38,
			}

			personSchema := Avoucher().Keys(map[string]Schema{
				"Name": Avoucher().String(),
				"Age" : Avoucher().Int(),
			})

			Expect(personSchema.Validate(objectToValidate).IsValid()).To(Equal(true))
		})

		It("should consider map valid when it contains all keys", func(){
			objectToValidate := map[string]interface{}{
				"Name": "Jason",
			}

			personSchema := Avoucher().Keys(map[string]Schema{
				"Name": Avoucher().String(),
				"Age" : Avoucher().Int(),
			})

			Expect(personSchema.Validate(objectToValidate).IsValid()).To(Equal(false))
		})

	})

	Describe("Examples", func(){
		It("should show example of validating several objects, and fail when struct is missing a Key", func(){
			//setup 2 structs which both define the same Name and Age properties
			type Person struct{
				Name string
				Age int
			}
			type Dog struct{
				Name string
				Age int
			}
			//setup struct that does not define the Name property
			type Rock struct{
				Age int
			}

			//define a schema which requires that Name and Age properties to be of type string and int, respectively
			animalSchema := Avoucher().Keys(map[string]Schema{
				"Name": Avoucher().String(),
				"Age": Avoucher().Int(),
			})

			//instantiate the objects we wish to validate
			person := Person{Name:"Jason"}
			dog := Dog{Name:"Rex"}
			cow := map[string]interface{}{
				"Name": "Mooo",
				"Age": 10,
			}
			rock := Rock{Age:100000}

			//demonstrate that person, dog, and cow all meet the requirements of the animalSchema
			Expect(animalSchema.Validate(person).IsValid()).To(Equal(true))
			Expect(animalSchema.Validate(dog).IsValid()).To(Equal(true))
			Expect(animalSchema.Validate(cow).IsValid()).To(Equal(true))

			//Fail because the Rock struct does not define a Name property
			Expect(animalSchema.Validate(rock).IsValid()).To(Equal(false))

		})

		It("should show example of nested schemas", func(){
			type Name struct{
				First string
				Middle string
				Last string
			}
			type Person struct{
				Name Name
				Age int
			}
			nameSchema := Avoucher().Keys(map[string]Schema{
				"First": Avoucher().String(),
				"Middle": Avoucher().String(),
				"Last": Avoucher().String(),
			})
			ageSchema := Avoucher().Int()

			personSchema := Avoucher().Keys(map[string]Schema{
				"Name": nameSchema,
				"Age": ageSchema,
			})

			person := Person{
				Age: 32,
				Name: Name{
					First: "Jason",
					Middle: "Lee",
					Last: "McAffee",
				},
			}

			Expect(personSchema.Validate(person).IsValid()).To(Equal(true))
		})
	})




})
