package model

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/Kamva/mgm"
)

// Mock
type Mock struct {
	mgm.DefaultModel `bson:",inline"`
	Path             string
	Method           string
	Response         string
	Code             int
	Headers          []CustomHeader
	ProxyMode        bool   `bson:"proxy_mode"`
	ProxyUrl         string `bson:"proxy_url"`
}

// NewMock ...
func NewMock(method string, path string, response string, code int, proxy_mode bool, proxy_url string) Mock {

	responseEncoded := base64.StdEncoding.EncodeToString([]byte(response))

	mock := Mock{
		Method:    method,
		Path:      path,
		Response:  responseEncoded,
		Code:      code,
		ProxyMode: proxy_mode,
		ProxyUrl:  proxy_url}

	return mock
}

type Mocks []Mock

type CustomHeader struct {
	Key   string
	Value string
}

// GetDecodedResponse ...
func (mock *Mock) GetDecodedResponse() ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(mock.Response)
	if err != nil {
		fmt.Println("decode error:", err)
		return nil, errors.New("Cant decode response")
	}
	return decoded, nil
}
