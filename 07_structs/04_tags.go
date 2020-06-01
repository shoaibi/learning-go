package main

import (
	"fmt"
	"reflect"
)

type TagStruct struct {
	// Tags can be anything
	// 3rd party libraries make use of it for json marshaling, validation, etc
	// Format:
	Name string `required max:"100"`
}

func main() {
	// extracting tags:
	t := reflect.TypeOf(TagStruct{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
