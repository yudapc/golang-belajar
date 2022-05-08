package main

import (
	"fmt"
	"testing"
)

func TestInterfaceKosong(t *testing.T) {
	var person interface{}
	person = "Joko"
	person = 10
	person = true
	person = "Budi"
	fmt.Println("Interface Kosong: ", person)
}
