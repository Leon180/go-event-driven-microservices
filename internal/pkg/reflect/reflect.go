package reflect

import (
	"fmt"
	"reflect"
)

func GetInstance[T any]() T {
	typ := GetType[T]()
	return getInstanceFromType(typ).(T)
}

func GetInstanceByType(typ reflect.Type) any {
	return getInstanceFromType(typ)
}

func GetType[T any]() reflect.Type {
	res := reflect.TypeOf((*T)(nil)).Elem()
	return res
}

func IsPointer[T any]() bool {
	t := reflect.TypeOf((*T)(nil)).Elem()
	return t.Kind() == reflect.Ptr
}

func GetAnysTypeName(input any) string {
	if input == nil {
		return ""
	}
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}
	return fmt.Sprintf("*%s", t.Elem().Name())
}

func GetAnysFullTypeName(input any) string {
	if input == nil {
		return ""
	}
	t := reflect.TypeOf(input)
	return t.String()
}

func GetTypeName[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}
	return fmt.Sprintf("*%s", t.Elem().Name())
}

func GetTypeNameFromType(typ reflect.Type) string {
	if typ.Kind() != reflect.Ptr {
		return typ.Name()
	}
	return fmt.Sprintf("*%s", typ.Elem().Name())
}

func getInstanceFromType(typ reflect.Type) any {
	if typ.Kind() == reflect.Ptr {
		res := reflect.New(typ.Elem()).Interface()
		return res
	}
	return reflect.Zero(typ).Interface()
}

func GetAnysType(value any) reflect.Type {
	if reflect.ValueOf(value).Kind() == reflect.Pointer {
		return reflect.TypeOf(reflect.ValueOf(value).Elem().Interface())
	}
	return reflect.TypeOf(value)
}
