package main

import "fmt"

func MapTypeData() {
	person := map[string]string{
		"name":    "Joko",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
