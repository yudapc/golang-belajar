package main

import (
	"fmt"
	"testing"
)

func clearData() {
	fmt.Println("Selesai memanggil function didalam defer")
}

func runApplication() {
	defer clearData()
	fmt.Println("Run Application")
}

func TestDefer(t *testing.T) {
	runApplication()
}
