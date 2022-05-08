package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func ContohErrorInterface() {
	_, err := pembagian(1, 0)
	fmt.Println("Error interface", err.Error())
}
