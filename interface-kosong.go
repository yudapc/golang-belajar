package main

import "fmt"

// interface{} adalah any kalau di typescript yang berarti semua data diperbolehkan

func PenggunaanInterfaceKosong() {
	var person interface{}
	person = "Joko"
	person = 10
	person = true
	person = "Budi"
	fmt.Println("Interface Kosong: ", person)
}
