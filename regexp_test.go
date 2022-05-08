package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto", 10))
}
