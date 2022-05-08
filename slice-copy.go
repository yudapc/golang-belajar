package main

import "fmt"

func CopySlice() {
	fromSlice := []int{1, 2, 3, 4, 5}

	// copy slice from fromSlice to toSlice
	toSlice := make([]int, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("Copy slice: ", toSlice)
}
