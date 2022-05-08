package main

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	person1 := Person{Name: "Joko"}
	person2 := &person1

	person2.Name = "Budi"

	fmt.Println("person1", person1)
	fmt.Println("person2", person2)
}
