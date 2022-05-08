package main

import (
	"fmt"
	"testing"
)

func ContohFunction() {
	fmt.Println("Ini contoh function")
}

func ContohFunctionWithParam(name string) {
	fmt.Println("Contoh function with param name", name)
}

func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func ContohFunctionWithReturn() {
	firstName := "Joko"
	lastName := "Susanto"
	fullName := getFullName(firstName, lastName)
	fmt.Println("Function with return: ", fullName)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func ContohVariadicFunction() {
	total := sumAll(1, 2, 3, 4, 5) // 15
	fmt.Println("Varadic function args: ", total)
}

func ContohVariadicFunctionSlice() {
	numbers := []int{10, 10, 10}
	total := sumAll(numbers...) // 30
	fmt.Println("Varadic function slice args: ", total)
}

func filterKataKotor(word string) string {
	if word == "Anjay" {
		return "A***y"
	}
	return word
}

func whiteWord(name string, filter func(string) string) {
	fmt.Println("Function as parameter: Halo ", filter(name))
}

func ContohFunctionAsParam() {
	whiteWord("Anjay", filterKataKotor)
	whiteWord("Joko", filterKataKotor)
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive((value - 1))
	}
}

func RecursiveFunction() {
	fmt.Println("factorial loop", factorialLoop(10))
	fmt.Println("recursive factorial loop", factorialRecursive(10))
}

func TestFunc(t *testing.T) {
	ContohFunction()
	ContohFunctionWithParam("Joko")
	ContohFunctionWithReturn()
	ContohVariadicFunction()
	ContohVariadicFunctionSlice()
	ContohFunctionAsParam()
	RecursiveFunction()
}
