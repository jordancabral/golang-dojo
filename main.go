package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"golang-dojo/app/mocks"
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
		setHandler(item.Path, item.Response)
	}

	http.ListenAndServe(":4000", nil)
}

func setHandler(path string, response string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, mocks.ResponseHello(path), html.EscapeString(r.URL.Path))
	})
}
