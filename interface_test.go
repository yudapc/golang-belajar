package main

import (
	"fmt"
	"testing"
)

type HasName interface {
	GetName() string
}

func sayHelloFromInterface(hasName HasName) { // interface HasName as a contract
	fmt.Println("Interface: Hello", hasName.GetName())
}

/// implementation interface HasName ke-1

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

/// implementation interface HasName ke-2

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name + " is my animal"
}

func TestInterface(t *testing.T) {
	budi := Person{Name: "Budi"}
	sayHelloFromInterface(budi)

	cat := Animal{Name: "Miaww"}
	sayHelloFromInterface(cat)
}
