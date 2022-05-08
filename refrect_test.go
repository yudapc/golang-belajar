package main

import (
	"fmt"
	"testing"
)

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
