package main

import (
	"fmt"
)

func Slice() {
	bilangan := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice: ", bilangan)
}

func SliceMake() {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Baru"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}

func CopySlice() {
	fromSlice := []int{1, 2, 3, 4, 5}

	// copy slice from fromSlice to toSlice
	toSlice := make([]int, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("Copy slice: ", toSlice)
}

func SliceAppend() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("Slice append: ", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println("daySlice2: ", daySlice2)
	fmt.Println("days: ", days)
}
