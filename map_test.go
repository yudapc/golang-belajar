package main

import (
	"fmt"
	"testing"
)

func MapTypeData() {
	person := map[string]string{
		"name":    "Joko",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}

func TestMapTypeData(t *testing.T) {
	MapTypeData()
}
