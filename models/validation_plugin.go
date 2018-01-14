package models

import . "avoucher/interfaces"

//validationPlugin implements the ValidationPlugin interface
type validationPlugin struct{
	ValidationFunc ValidationFunc
	Name string
}

//NewValidationPlugin returns a struct pointer which implements ValidationPlugin
func NewValidationPlugin(name string, validationFunc ValidationFunc) ValidationPlugin{
	return &validationPlugin{
		Name: name,
		ValidationFunc: validationFunc,
	}
}

//GetValidationFunc returns the ValidationFunc which is used to validate an object against a schema.
func (v *validationPlugin) GetValidationFunc() ValidationFunc{
	return v.ValidationFunc
}

//SetValidationFunc sets the ValidationFunc to be used to validate an object against a schema.
func (v *validationPlugin) SetValidationFunc(validationFunc ValidationFunc){
	v.ValidationFunc = validationFunc
}

func (v *validationPlugin) SetName(pluginName string){
	v.Name = pluginName
}

func (v *validationPlugin) GetName() string{
	return v.Name
}

func (v *validationPlugin) CreateValidationResult() ValidationResult{
	validationResult := NewValidationResult()
	validationResult.SetIsValid(true)
	validationResult.SetTestName(v.Name)
	return validationResult
}