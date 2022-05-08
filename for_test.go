package main

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	counter := 1

	for counter <= 3 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
}

func TestForLoopWithStatement(t *testing.T) {
	for counter := 1; counter <= 2; counter++ {
		fmt.Println("Perlulangan dengan statement ke", counter)
	}
}

func TestForRange(t *testing.T) {
	names := []string{"Eko", "Budi", "Joko"}
	for index, name := range names {
		fmt.Println("For range index", index, "=", name)
	}
}
