package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Mock
type Mock struct {
	Path     string
	Response string
}

func main() {
	fmt.Println("Hola Mundo")

	file, error := ioutil.ReadFile("config.json")
	if error != nil {
		panic(error)
	}

	mock := []Mock{}

	json.Unmarshal(file, &mock)

	fmt.Println(mock)

	for _, item := range mock {
		handler(item.Path, item.Response)
	}

	http.ListenAndServe(":4000", nil)
}

func handler(path string, response string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, response)
	})
}

func (mock Mock) handler() {
	http.HandleFunc(mock.Path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, mock.Response)
	})
}
