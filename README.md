# Avoucher
Validation library for Go.

## Usage Examples
```go
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
```


## Developer Setup

### Glide
From the root directory, run
```
glide install
```

### Running Tests
