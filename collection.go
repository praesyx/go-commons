package commons

import (
	"errors"
	"reflect"
)

type Collection struct {
	Objects []interface{}
}

// First возвращает первый элемент из коллекции
func (c *Collection) First() (any, error) {
	if err := c.notEmptySlice(); nil != err {
		return nil, err
	}

	return c.Objects[0], nil
}

// Last возвращает последний элемент из коллекции
func (c *Collection) Last() (any, error) {
	if err := c.notEmptySlice(); nil != err {
		return nil, err
	}

	return c.Objects[len(c.Objects)-1], nil
}

// Add добавляет элемент в коллекцию
func (c *Collection) Add(object any) {
	c.Objects = append(c.Objects, object)
}

// Remove удаляет элемент из коллекции
func (c *Collection) Remove(object any) {
	for i, obj := range c.Objects {
		if reflect.DeepEqual(obj, object) {
			c.Objects = append(c.Objects[:i], c.Objects[i+1:]...)
			if 0 == len(c.Objects) {
				c.Clear()
			}
			return
		}
	}
}

// Clear очищает всю коллекцию
func (c *Collection) Clear() {
	c.Objects = nil
}

// Contains проверяет, содержится ли элемент в коллекции
func (c *Collection) Contains(object any) (exists bool) {
	exists = false
	for _, obj := range c.Objects {
		if reflect.DeepEqual(obj, object) {
			exists = true
			return
		}
	}
	return
}

// Matching находит совпадение в коллекции
func (c *Collection) Matching(criteria map[string]interface{}) (any, error) {
	if err := c.notEmptySlice(); nil != err {
		return nil, err
	}
	exists := false
	index := 0
	for k, v := range criteria {
		if ok, i := InArrayByField(v, &c.Objects, k); ok {
			if 1 == len(criteria) {
				exists = true
				index = i
				break
			}
			for _, v := range criteria {
				if ok, _ := InArrayByField(v, &c.Objects, k); !ok {
					exists = false
					continue
				}
				exists = true
			}
			if exists {
				index = i
			}
		}
	}
	if !exists {
		return nil, errors.New("matching not found")
	}

	return c.Objects[index], nil
}

func (c *Collection) notEmptySlice() error {
	if len(c.Objects) == 0 {
		return errors.New("empty slice of interfaces")
	}

	return nil
}

// Filter вернет новый объект Collection с отфильтрованными условиями которые переданы в функции фильтр
// Пример:
// type A struct {
//	Name string
//}
// func main() {
//	var a []A
//	for i := 0; i < 5; i++ {
//		a = append(a, A{Name: fmt.Sprintf("Hello%d", i)})
//	}
//	c, _ := Filter(a, func(a any) bool {
//		if a.(A).Name != "Hello2" {
//			return false
//		}
//		return true
//	})
//	fmt.Println(c.Last())
//}
func Filter(objects any, filter func(T any) bool) (Collection, error) {
	var collection Collection
	if reflect.TypeOf(objects).Kind() != reflect.Slice {
		return collection, errors.New("interface type is not slice")
	}

	s := reflect.ValueOf(objects)
	for i := 0; i < s.Len(); i++ {
		if !filter(s.Index(i).Interface()) {
			continue
		}

		collection.Objects = append(collection.Objects, s.Index(i).Interface())
	}
	return collection, nil
}
