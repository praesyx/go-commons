package commons

import (
	"reflect"
	"strings"
)

// InArrayByField проверка, присутствует значение в интерфейсе по полю
// Пример:
// type A struct {
//	R struct {
//		Val []string
//	}
//}
// a := A{R: struct{Val []string}{Val: []string{"ROLE"}}}
// if ok := utils.InArrayByField("ROLE", &a.R, "val"); ok { fmt.Println("OK") }
func InArrayByField(needle interface{}, itr interface{}, field string) (exists bool) {
	exists = false
	val := reflect.ValueOf(itr)
	switch val.Kind() {
	case reflect.Ptr:
		v := val.Elem()
		for i := 0; i < v.Len(); i++ {
			typeOfT := v.Index(i).Type()
			for n := 0; n < v.Index(i).NumField(); n++ {
				if typeOfT.Field(n).Anonymous {
					for j := 0; j < v.Index(i).Field(n).NumField(); j++ {
						if strings.ToLower(v.Index(i).Field(n).Type().Field(j).Name) == field &&
							reflect.DeepEqual(needle, v.Index(i).Field(n).Field(j).Interface()) {
							exists = true
							return
						}
					}
				}
				switch v.Index(i).Field(n).Kind() {
				case reflect.Slice, reflect.Array:
					if strings.ToLower(typeOfT.Field(n).Name) != field {
						break
					}
					for j := 0; j < v.Index(i).Field(n).Len(); j++ {
						if reflect.DeepEqual(needle, v.Index(i).Field(n).Index(j).Interface()) == true {
							exists = true
							return
						}
					}
				default:
					if strings.ToLower(typeOfT.Field(n).Name) == field &&
						reflect.DeepEqual(needle, v.Index(i).Field(n).Interface()) {
						exists = true
						return
					}
				}
			}
		}
	}
	return
}

// InArray проверка, присутствует значение в срезе или массиве.
// Пример:
// if ok := InArray(3, uint8{1,2,3}); ok { fmt.Println("OK") }
func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	switch reflect.TypeOf(array).Kind() {
	case reflect.Ptr:
		v := reflect.Indirect(reflect.ValueOf(array).Elem())
		for i := 0; i < v.Len(); i++ {
			if ok := reflect.DeepEqual(val, v.Index(i).Interface()); ok {
				exists = ok
				return
			}
		}
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if ok := reflect.DeepEqual(val, s.Index(i).Interface()); ok {
				exists = ok
				return
			}
		}
	}

	return
}
