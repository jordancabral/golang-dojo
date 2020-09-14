package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordancabral/golang-dojo/app/mocks"
	. "github.com/jordancabral/golang-dojo/app/model"
	"github.com/jordancabral/golang-dojo/app/repository"
	"github.com/softbrewery/gojoi/pkg/joi"
)

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

	//Create config (for testing pourposes)
	//mockTest := Mock{Path: "/hola", Method: "GET", Response: "app/mocks/test.json", Code: 200, ProxyMode: false, ProxyUrl: "http://www.google.com"}
	//repository.CreateConfig(mockTest)

	// Init Gin
	r := gin.Default()

	// mock config CRUD
	// example curl: curl -XPOST http://localhost:8080/mock --header 'Content-Type: application/json' --data-raw '{"path":"test","Method":"GET", "Response": "app/mocks/test.json", "Code": 200, "ProxyMode": false, "ProxyUrl": "http://www.google.com"}'
	// TODO: extract to module
	r.POST("/mock", func(c *gin.Context) {
		mockTest := Mock{}
		c.ShouldBind(&mockTest)

		validate := validateMock(mockTest)
		if nil != validate {
			fmt.Println("mock mal armado")
			c.JSON(400, gin.H{"message": "mock mal armado"})
		} else {
			repository.CreateConfig(mockTest)
			c.JSON(200, gin.H{"message": "ok"})
		}
	})

	// Load mock configs from DB and configure the routes
	result := repository.GetAllConfigs()
	for _, item := range result {
		validate := validateMock(item)
		if nil != validate {
			fmt.Println("mock mal armado")
		} else {
			mock := item
			r.Handle(mock.Method, mock.Path, func(c *gin.Context) {

				// If proxy is enabled, make the request
				if mock.ProxyMode {
					loadProxy(c, mock)
					return
				}

				// If proxy is disabled load mock
				response, error := mocks.LoadMock(mock.Response)
				if error != nil {
					message := "File not found for this path: " + item.Response
					c.JSON(http.StatusNotImplemented, gin.H{"message": message})
					return
				}
				c.Data(http.StatusOK, "application/json", []byte(response))
				return
			})
		}
	}
	r.Run()

}

// TODO: extract to module
func loadProxy(c *gin.Context, mock Mock) {
	fmt.Printf("Proxy enabled, URL: %s Method: %s", mock.ProxyUrl, mock.Method)
	var myBody io.Reader
	myHeaders := c.Request.Header
	if mock.Method != "GET" {
		myBody = c.Request.Body
	}
	request, requestError := http.NewRequest(mock.Method, mock.ProxyUrl, myBody)
	if requestError != nil {
		panic(requestError)
	}

	request.Header = myHeaders
	client := &http.Client{}
	proxyResponse, proxyError := client.Do(request)
	if proxyError != nil {
		message := "Proxy error: " + mock.ProxyUrl + proxyError.Error()
		c.JSON(http.StatusBadGateway, gin.H{"message": message})
		return
	}
	fmt.Println("Proxy response status:", proxyResponse.Status)

	for key, header := range proxyResponse.Header {
		for _, h := range header {
			c.Header(key, h)
		}
	}
	io.Copy(c.Writer, proxyResponse.Body)
	return
}
