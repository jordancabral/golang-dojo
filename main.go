package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jordancabral/golang-dojo/app/mocks"
	"github.com/softbrewery/gojoi/pkg/joi"
)

// Mock
type Mock struct {
	Path      string
	Method    string
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

func validateMock(mock Mock) error {

	schemaCustomHeader := joi.Struct().Keys(joi.StructKeys{
		"key":   joi.String().NonZero(),
		"value": joi.String().NonZero(),
	})

	schemaMock := joi.Struct().Keys(joi.StructKeys{
		"Path":      joi.String().NonZero(),
		"Method":    joi.String().NonZero(),
		"Response":  joi.String().NonZero(),
		"Code":      joi.Int().NonZero(),
		"Headers":   joi.Slice().Items(schemaCustomHeader),
		"ProxyMode": joi.Bool().Required(),
		"ProxyUrl":  joi.String().NonZero(),
	})
	err := joi.Validate(mock, schemaMock)
	return err
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

	paths := make(map[string][]Mock)

	for _, item := range mock {
		validate := validateMock(item)
		if nil != validate {
			fmt.Println("mock mal armado")
			fmt.Println(item.Path)
			fmt.Println(validate)
		} else {
			pathArray := paths[item.Path]
			pathArray = append(pathArray, item)
			paths[item.Path] = pathArray
		}
	}

	for key, val := range paths {
		setPath(key, val)
	}
	//setHandler(item.Path, item.Response, item.Code, item.Headers, item.ProxyMode, item.ProxyUrl, item.Method)

	http.ListenAndServe(":4000", nil)
}

func setPath(path string, mocks []Mock) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		for _, val := range mocks {
			if val.Method == r.Method {
				setHandler(val.Path, val.Response, val.Code, val.Headers, val.ProxyMode, val.ProxyUrl, val.Method, w, r)
				return
			}
		}

		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return

	})
}

func setHandler(path string, response string, statusCode int, Headers []CustomHeader, ProxyMode bool, ProxyUrl string, method string, w http.ResponseWriter, r *http.Request) {

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
			for _, h := range header {
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
}
