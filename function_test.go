package main

import "testing"

func TestFunc(t *testing.T) {
	ContohFunction()
	ContohFunctionWithParam("Joko")
	ContohFunctionWithReturn()
	ContohVariadicFunction()
	ContohVariadicFunctionSlice()
	ContohFunctionAsParam()
	RecursiveFunction()
}
