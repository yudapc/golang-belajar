package main

import "fmt"

func Break() {
	for i := 0; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println("Perulangan break ke", i)
	}
}

func Continue() {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Perulangan continue ke", i)
	}
}
