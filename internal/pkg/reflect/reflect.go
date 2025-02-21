package reflect

import (
	"fmt"
	"reflect"
)

func GetInstance[T any]() T {
	return getInstanceFromType(GetType[T]()).(T)
}

func GetType[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func GetTypeName[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}
	return fmt.Sprintf("*%s", t.Elem().Name())
}

func getInstanceFromType(typ reflect.Type) interface{} {
	if typ.Kind() == reflect.Ptr {
		return reflect.New(typ.Elem()).Interface()
	}
	return reflect.Zero(typ).Interface()
}
