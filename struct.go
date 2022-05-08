package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func CreateDataFromStruct() {
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}

	fmt.Println("Create data from struct", joko)
}

// struct method

func (c Customer) sayHello() {
	fmt.Println("Hello ", c.Name)
}

func RunStructMethod() {
	budi := Customer{
		Name:    "Budi",
		Address: "Semarang",
		Age:     35,
	}
	budi.sayHello()
}
