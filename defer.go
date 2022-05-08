package main

import "fmt"

func clearData() {
	fmt.Println("Selesai memanggil function didalam defer")
}

func runApplication() {
	defer clearData()
	fmt.Println("Run Application")
}

func ContohDefer() {
	runApplication()
}
