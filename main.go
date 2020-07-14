package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jordancabral/golang-dojo/app/mocks"
)

// Mock
type Mock struct {
	Path      string
	Method	  string
	Response  string
	Code      int
	Headers   []CustomHeader
	ProxyMode bool   `json:"proxy_mode"`
	ProxyUrl  string `json:"proxy_url"`
}

type CustomHeader struct {
	Key   string
	Value string
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
		setHandler(item.Path, item.Response, item.Code, item.Headers, item.ProxyMode, item.ProxyUrl, item.Method)
	}

	http.ListenAndServe(":4000", nil)
}

func setHandler(path string, response string, statusCode int, Headers []CustomHeader, ProxyMode bool, ProxyUrl string, method string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			http.Error(w, "Invalid Method", http.StatusNotFound)
			return	
		}
		if ProxyMode == true {
			fmt.Printf("\nEsto deberia ser el proxy, url:%s", ProxyUrl)
			fmt.Printf("\nMethod:%s", r.Method)
			var myBody io.Reader
			myHeaders := r.Header
			if r.Method != "GET" {
				myBody = r.Body
			}
			request, requestError := http.NewRequest(r.Method, ProxyUrl, myBody)
			if requestError != nil {
				panic(requestError)
			}

			request.Header = myHeaders
			client := &http.Client{}
			proxyResponse, proxyError := client.Do(request)
			if proxyError != nil {
				http.Error(w, "Rompió el proxy:"+ProxyUrl, http.StatusBadGateway)
				return
			}
			fmt.Println("Response status:", proxyResponse.Status)

			for key, header := range proxyResponse.Header {
				for _,h := range header {
					w.Header().Set(key, h)
				}
			}
			
			io.Copy(w, proxyResponse.Body)
			return
		}
		response, error := mocks.ResponseHello(response)
		if error != nil {
			http.Error(w, "File not found for this path", http.StatusNotImplemented)
			return
		}
		fmt.Printf("\nResponse with code:%d for path:%s, headers:%s", statusCode, path, Headers)
		for _, header := range Headers {
			w.Header().Set(header.Key, header.Value)
		}
		w.WriteHeader(statusCode)
		fmt.Fprintf(w, response)
	})
}
