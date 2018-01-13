package avoucher

type(
	Schema struct{
		//the type of object we are validating. todo: allow for AnyKind, so that we only validate keys.
		Kind interface{}
		isKindSet bool

		//properties on the struct to be validated.
		Keys map[string]*Schema
	}
)

func NewSchema() *Schema {
	return &Schema{}
}

func (s *Schema) SetKind(kind interface{}) *Schema{
	s.Kind = kind
	s.isKindSet = true
	return s
}

//TODO: if Keys is already set, iterate over keys and overwrite appropriate Keys
func (s *Schema) SetKeys(keys map[string]*Schema) *Schema{
	s.Keys = keys
	return s
}

func (s *Schema) Validate(objToValidate interface{}) (bool, error){
	isValid, err := Validate(s, objToValidate)
	return isValid, err
}


