package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string `json:"name-field"`
	age  int
}

func main() {
	user := &User{"John Doe The Fourth", 20}

	field, ok := reflect.TypeOf(user).Elem().FieldByName("name")
	if !ok {
		panic("Field not fount")
	}

	fmt.Println(getStructTag(field))

}

func getStructTag(f reflect.StructField) string {
	//可以对 tag 进行修改
	//f.Tag = "hello"  // 输出 hello
	return string(f.Tag)
}
