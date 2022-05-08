package main

import "fmt"

func SwitchCase() {
	name := "Joko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}
}
