package reflect

import (
	"fmt"
	"reflect"
)

func GetInstance[T any]() T {
	typ := GetType[T]()
	return getInstanceFromType(typ).(T)
}

func GetType[T any]() reflect.Type {
	res := reflect.TypeOf((*T)(nil)).Elem()
	return res
}

func IsPointer[T any]() bool {
	t := reflect.TypeOf((*T)(nil)).Elem()
	return t.Kind() == reflect.Ptr
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
		res := reflect.New(typ.Elem()).Interface()
		return res
	}
	return reflect.Zero(typ).Interface()
}
