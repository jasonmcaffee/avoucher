package avoucher

type(
	Schema struct{
		Kind interface{}
	}
)

func NewSchema() Schema {
	return Schema{}
}

func (s Schema) SetKind(kind interface{}) Schema{
	s.Kind = kind
	return s
}

func (s Schema) Validate(objToValidate interface{}) bool{
	isValid := Validate(s, objToValidate)
	return isValid
}


