package main

import (
	"fmt"
	"testing"
)

type Customer struct {
	Name, Address string
	Age           int
}

// struct method

func (c Customer) sayHello() {
	fmt.Println("Hello ", c.Name)
}

func TestStruct(t *testing.T) {
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}

	fmt.Println("Create data from struct", joko)
}

func TestStructMethod(t *testing.T) {
	budi := Customer{
		Name:    "Budi",
		Address: "Semarang",
		Age:     35,
	}
	budi.sayHello()
}
