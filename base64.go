package commons

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

// Base64Unmarshal сериализация JSON объекта в структуру
// Пример:
// type A struct {
//	B string `json:"b"`
// }
// var a A
// if err := Base64Unmarshal("{"b":"hello"}", &a); nil != err { fmt.Println(err) }
func Base64Unmarshal(s string, itr interface{}) error {
	b, _ := base64.StdEncoding.DecodeString(s)
	err := json.Unmarshal(b, &itr)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
	}
	return err
}

// MapToStructureUnmarshal сериализация карты в структуру
// для примера смотрите документацию https://pkg.go.dev/github.com/mitchellh/mapstructure
func MapToStructureUnmarshal(in interface{}, itr interface{}) interface{} {
	_ = mapstructure.Decode(in, &itr)
	return itr
}
