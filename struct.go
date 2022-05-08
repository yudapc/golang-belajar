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
