package commons

import (
	"cmp"
	"reflect"
	"strings"
)

const NilPosition = -1

// InArrayByField проверка, присутствует значение в интерфейсе по полю
// Пример:
//
//	type A struct {
//		R struct {
//			Val []string
//		}
//	}
//
// a := A{R: struct{Val []string}{Val: []string{"ROLE"}}}
// if ok, i := utils.InArrayByField("ROLE", &a.R, "val"); ok { fmt.Println("OK", i) }
func InArrayByField(needle interface{}, itr interface{}, field string) (exists bool, index int) {
	exists = false
	index = 0
	val := reflect.ValueOf(itr)
	switch val.Kind() {
	case reflect.Ptr:
		v := val.Elem()
		for i := 0; i < v.Len(); i++ {
			typeOfT := v.Index(i).Type()
			itr := v.Index(i)
			if typeOfT.Kind() == reflect.Interface {
				typeOfT = v.Index(i).Elem().Type()
				itr = v.Index(i).Elem()
			}
			for n := 0; n < itr.NumField(); n++ {
				if typeOfT.Field(n).Anonymous {
					for j := 0; j < itr.Field(n).NumField(); j++ {
						if strings.ToLower(itr.Field(n).Type().Field(j).Name) == field &&
							reflect.DeepEqual(needle, itr.Field(n).Field(j).Interface()) {
							exists = true
							index = j
							return
						}
					}
				}
				switch itr.Field(n).Kind() {
				case reflect.Slice, reflect.Array:
					if strings.ToLower(typeOfT.Field(n).Name) != field {
						break
					}
					for j := 0; j < itr.Field(n).Len(); j++ {
						if reflect.DeepEqual(needle, itr.Field(n).Index(j).Interface()) == true {
							exists = true
							index = j
							return
						}
					}
				default:
					if strings.ToLower(typeOfT.Field(n).Name) == field &&
						reflect.DeepEqual(needle, itr.Field(n).Interface()) {
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
// if index := InArray(3, []uint8{1,2,3}); index != NilPosition { fmt.Println("OK") }
func InArray[S ~[]E, E cmp.Ordered](val E, array S) int {
	for i, v := range array {
		if v == val {
			return i
		}
	}

	return NilPosition
}
