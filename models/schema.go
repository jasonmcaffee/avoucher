package models

import . "avoucher/interfaces"
type (
	schema struct {
		//the type of object we are validating.
		Type      interface{}
		IsTypeSet bool

		//properties on the struct to be validated.
		Keys map[string]Schema
		Validator Validator
	}
)

func NewSchema(validator Validator) Schema {
	return &schema{
		Validator : validator,
	}
}

func (s *schema) SetType(t interface{}) Schema {
	s.Type = t
	s.IsTypeSet = true
	return s
}

func (s *schema) GetType() interface{}{
	return s.Type
}

func (s *schema) GetIsTypeSet() bool{
	return s.IsTypeSet
}

//TODO: if Keys is already set, iterate over keys and overwrite appropriate Keys
func (s *schema) SetKeys(keys map[string]Schema) Schema {
	s.Keys = keys
	return s
}

func (s *schema) GetKeys() map[string]Schema{
	return s.Keys
}

func (s *schema) Validate(objToValidate interface{}) ValidationResult {
	validationResult := s.Validator.Validate(s, objToValidate)
	return validationResult
}