package models

import (
	. "avoucher/interfaces"
	"reflect"
)
type (
	schema struct {
		//the type of object we are validating.
		TheType       interface{}
		IsTypeSet     bool
		ReflectedType reflect.Type

		//properties on the struct to be validated.
		TheKeys   map[string]Schema
		Validator Validator
	}
)

func NewSchema(validator Validator) Schema {
	schema := &schema{
		Validator : validator,
	}
	return schema
}

//shortcut for SetType
func (s *schema) Type(t interface{}) Schema{
	return s.SetType(t)
}

//shortcuts for Numeric types ==========================================================================================
func (s *schema) Int() Schema{
	var t int
	return s.SetType(t)
}
func (s *schema) Int8() Schema{
	var t int8
	return s.SetType(t)
}
func (s *schema) Int16() Schema{
	var t int16
	return s.SetType(t)
}
func (s *schema) Int32() Schema{
	var t int32
	return s.SetType(t)
}
func (s *schema) Int64() Schema{
	var t int64
	return s.SetType(t)
}

func (s *schema) Uint() Schema{
	var t uint
	return s.SetType(t)
}
func (s *schema) Uint8() Schema{
	var t uint8
	return s.SetType(t)
}
func (s *schema) Uint16() Schema{
	var t uint16
	return s.SetType(t)
}
func (s *schema) Uint32() Schema{
	var t uint32
	return s.SetType(t)
}
func (s *schema) Uint64() Schema{
	var t uint64
	return s.SetType(t)
}
func (s *schema) Float32() Schema{
	var t float32
	return s.SetType(t)
}
func (s *schema) Float64() Schema{
	var t float64
	return s.SetType(t)
}
func (s *schema) Complex64() Schema{
	var t complex64
	return s.SetType(t)
}
func (s *schema) Complex128() Schema{
	var t complex128
	return s.SetType(t)
}
func (s *schema) Byte() Schema{
	var t byte
	return s.SetType(t)
}
func (s *schema) Rune() Schema{
	var t rune
	return s.SetType(t)
}

func (s *schema) IntPointer() Schema{
	var t *int
	return s.SetType(t)
}
func (s *schema) Int8Pointer() Schema{
	var t *int8
	return s.SetType(t)
}
func (s *schema) Int16Pointer() Schema{
	var t *int16
	return s.SetType(t)
}
func (s *schema) Int32Pointer() Schema{
	var t *int32
	return s.SetType(t)
}
func (s *schema) Int64Pointer() Schema{
	var t *int64
	return s.SetType(t)
}

func (s *schema) UintPointer() Schema{
	var t *uint
	return s.SetType(t)
}
func (s *schema) Uint8Pointer() Schema{
	var t *uint8
	return s.SetType(t)
}
func (s *schema) Uint16Pointer() Schema{
	var t *uint16
	return s.SetType(t)
}
func (s *schema) Uint32Pointer() Schema{
	var t *uint32
	return s.SetType(t)
}
func (s *schema) Uint64Pointer() Schema{
	var t *uint64
	return s.SetType(t)
}
func (s *schema) Float32Pointer() Schema{
	var t *float32
	return s.SetType(t)
}
func (s *schema) Float64Pointer() Schema{
	var t *float64
	return s.SetType(t)
}
func (s *schema) Complex64Pointer() Schema{
	var t *complex64
	return s.SetType(t)
}
func (s *schema) Complex128Pointer() Schema{
	var t *complex128
	return s.SetType(t)
}
func (s *schema) BytePointer() Schema{
	var t *byte
	return s.SetType(t)
}
func (s *schema) RunePointer() Schema{
	var t *rune
	return s.SetType(t)
}

//shortcuts for other types ===========================================================================================
func (s *schema) String() Schema{
	var t string
	return s.SetType(t)
}
func (s *schema) StringPointer() Schema{
	var t string
	return s.SetType(&t)
}


func (s *schema) SetType(t interface{}) Schema {
	s.TheType = t
	s.IsTypeSet = true
	s.ReflectedType = reflect.TypeOf(t)
	return s
}

func (s *schema) GetType() interface{}{
	return s.TheType
}

func (s *schema) GetTypeReflectedType() reflect.Type{
	return s.ReflectedType
}

func (s *schema) GetIsTypeSet() bool{
	return s.IsTypeSet
}

//TODO: if TheKeys is already set, iterate over keys and overwrite appropriate TheKeys
func (s *schema) Keys(keys map[string]Schema) Schema {
	return s.SetKeys(keys)
}
func (s *schema) SetKeys(keys map[string]Schema) Schema {
	s.TheKeys = keys
	return s
}

func (s *schema) GetKeys() map[string]Schema{
	return s.TheKeys
}

func (s *schema) Validate(objToValidate interface{}) ValidationResult {
	validationResult := s.Validator.Validate(s, objToValidate)
	return validationResult
}
