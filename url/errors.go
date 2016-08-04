package url

import (
	"fmt"
	"reflect"
)

type ArgumentError struct {
	Kind  reflect.Kind
	Value interface{}
}

func NewArgumentError(k reflect.Kind, v interface{}) ArgumentError {
	return ArgumentError{
		Kind:  k,
		Value: v,
	}
}

func (ae ArgumentError) Error() string {
	return fmt.Sprintf("I need my argument of type interface{} to be a struct - you gave me a %s '%#v'", ae.Kind, ae.Value)
}

type FieldError struct {
	Kind reflect.Kind
	Name string
}

func NewFieldError(k reflect.Kind, n string) FieldError {
	return FieldError{
		Kind: k,
		Name: n,
	}
}

func (fe FieldError) Error() string {
	return fmt.Sprintf("I need all the fields in my argument to be strings - you gave me a struct with a field '%s' which is a %s", fe.Name, fe.Kind)
}
