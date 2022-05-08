package main

import (
	"fmt"
	"testing"
)

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Error", message)
	fmt.Println("Berhasil di recover sudah tidak panic lagi")
}

func runApp(error bool) {
	defer endApp()
	if error {
		fmt.Println("PANIC")
		panic("Error")
	}
}

func TestPanicWithRecover(t *testing.T) {
	runApp(true)
}
