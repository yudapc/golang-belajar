package main

import "fmt"

func (c Customer) sayHello() {
	fmt.Println("Hello ", c.Name)
}

func RunStructMethod() {
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}
	joko.sayHello()
}
