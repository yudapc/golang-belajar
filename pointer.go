package main

import "fmt"

func ContohPengguaanPointer() {
	person1 := Person{Name: "Joko"}
	person2 := &person1

	person2.Name = "Budi"

	fmt.Println("person1", person1)
	fmt.Println("person2", person2)
}
