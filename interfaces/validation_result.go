package interfaces

type ValidationResult interface {
	IsValid() bool
	SetIsValid(isValid bool)
	GetMessage() string
	SetMessage(message string)
}
