package avoucher

import . "avoucher/models"
import . "avoucher/interfaces"
import . "avoucher/validation_plugins"

//DefaultValidationPlugins is a slice of validations to be performed by default.
var defaultValidationPlugins []ValidationPlugin = []ValidationPlugin{
	ValidateType,
}

func Avoucher() Schema {
	validator := NewValidator(defaultValidationPlugins)
	return NewSchema(validator)
}
