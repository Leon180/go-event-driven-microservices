package types

import (
	"reflect"
)

type TypeMaker interface {
	RegistType(name string, typ reflect.Type)
	GetTypeInstance(name string) (any, error)
	GetTypeInstancePointer(name string) (any, error)
}
