package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func ContohReflect() {
	sample := Sample{Name: "Joko"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println("Sample reflect: ", structField)
}

// validation library
func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func TestRefrect(t *testing.T) {
	ContohReflect()
}

func TestValidationValid(t *testing.T) {
	sample := Sample{Name: "Anwar"}
	fmt.Println("is valid?", isValid(sample))
}

func TestValidationInvalid(t *testing.T) {
	sample := Sample{Name: ""}
	fmt.Println("is valid?", isValid(sample))
}
