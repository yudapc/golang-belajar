package main

import (
	"fmt"
	"testing"
)

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

func TestBreak(t *testing.T) {
	Break()
}

func TestContinue(t *testing.T) {
	Continue()
}
