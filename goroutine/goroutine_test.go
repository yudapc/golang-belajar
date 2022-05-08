package main

import (
	"fmt"
	"testing"
	"time"
)

// Basic go routine

func RunHelloWorld() {
	fmt.Println("Hello World")
}
func TestBasicGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

// Banyak go routine

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
