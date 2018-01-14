package interfaces

//ValidationResult is returned by validate functions to indicate whether the test passed, the reason why it failed, etc.
type ValidationResult interface {
	IsValid() bool
	SetIsValid(bool)

	//GetMessage returns the message set by a validation plugin. Empty string when IsValid() == true.
	GetMessage() string
	//SetMessage sets the message indicating why a given test failed.
	SetMessage(string)

	//SetExpectedValue is used to indicate what a given validation was expecting. (some value on the schema)
	SetExpectedValue(interface{})
	//GetExpectedValue returns what the test was expecting. (some value on the schema)
	GetExpectedValue() interface{}

	//SetActualValue is used to indicate the value a given test received. (objToValidate)
	SetActualValue(interface{})
	//GetActualValue returns what the test received. (objToValidate)
	GetActualValue() interface{}

	//SetTestName is used to indicate which test created the ValidationResult.
	SetTestName(string)
	//GetTestName returns the name of the test that created the ValidationResult.
	GetTestName() string
}
