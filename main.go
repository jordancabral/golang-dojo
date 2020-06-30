package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jordancabral/golang-dojo/app/mocks"
)

// Mock
type Mock struct {
	Path     string
	Response string
	Code     int
}

func main() {
	fmt.Println("Starting Mock Server")

	file, error := ioutil.ReadFile("config.json")
	if error != nil {
		panic(error)
	}

	mock := []Mock{}

	json.Unmarshal(file, &mock)

	fmt.Println("Loaded mocks:")
	fmt.Printf("%+v\n", mock)

	for _, item := range mock {
		setHandler(item.Path, item.Response, item.Code)
	}

	http.ListenAndServe(":4000", nil)
}

func setHandler(path string, response string, statusCode int) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		response, error := mocks.ResponseHello(response)
		if error != nil {
			http.Error(w, "File not found for this path", http.StatusNotImplemented)
			return
		}
		fmt.Printf("\nResponse with code:%d for path:%s", statusCode, path)
		w.WriteHeader(statusCode)
		fmt.Fprintf(w, response)
	})
}
