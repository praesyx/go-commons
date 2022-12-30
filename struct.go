package commons

import (
	"reflect"
)

func StructFieldHasEmptyValue(itr interface{}) (exists bool, name string) {
	exists = false
	name = ""
	val := reflect.ValueOf(itr)
	switch val.Kind() {
	case reflect.Struct:
		for n := 0; n < val.NumField(); n++ {
			v := val.Field(n)
			if isZeroValue(v) {
				exists = true
				name = val.Type().Field(n).Name
				return
			}
		}
	}
	return
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZeroValue(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZeroValue(v.Field(i))
		}
		return z
	}

	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}
