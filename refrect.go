package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func ContohReflect() {
	sample := Sample{Name: "Joko"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println("Sample reflect: ", structField)
}
