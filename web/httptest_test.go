package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello World")
}

func TestHttpTest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
