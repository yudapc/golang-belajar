package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloJoko() string {
	return "Hello Joko"
}

func TestHelloJokoManual(t *testing.T) {
	expected := "Hello Joko"
	result := HelloJoko()
	if result != "Hello Joko" {
		fmt.Println("got: ", result)
		fmt.Println("expected: ", expected)
		t.Error("Harusnya Hello Joko")
	}
}

func TestHelloJokoWithAssert(t *testing.T) {
	expected := "Hello Joko"
	result := HelloJoko()
	assert.Equal(t, expected, result)
}

func TestSkip(t *testing.T) {
	t.Skip("Sekip dulu ya guys..")
}

func sayHelloPeople(name string) string {
	return "Hello " + name
}

func TestHelloPeopleTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Name is Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
		{
			name:     "Name is Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sayHelloPeople(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
