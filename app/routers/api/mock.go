package routers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/jordancabral/golang-dojo/app/model"
	"github.com/jordancabral/golang-dojo/app/repository"
)

// MockRoute ...
func MockRoute(c *gin.Context) {

	mockPath := c.Param("path")
	mock, error := repository.GetConfig(c.Request.Method, mockPath)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "path not found"})
		return
	}

	// If proxy is enabled, make the request
	if mock.ProxyMode {
		loadProxy(c, mock)
		return
	}

	// If proxy is disabled load mock
	decoded, err := mock.GetDecodedResponse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding response"})
		return
	}

	c.Data(http.StatusOK, "application/json", decoded)
	return

}

// TODO: extract to module
func loadProxy(c *gin.Context, mock *Mock) {
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
