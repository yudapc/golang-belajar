package main

import "fmt"

func ForLoop() {
	counter := 1

	for counter <= 3 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
}

func ForLoopWithStatement() {
	for counter := 1; counter <= 2; counter++ {
		fmt.Println("Perlulangan dengan statement ke", counter)
	}
}

func ForRange() {
	names := []string{"Eko", "Budi", "Joko"}
	for index, name := range names {
		fmt.Println("For range index", index, "=", name)
	}
}
