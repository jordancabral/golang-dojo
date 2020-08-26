package model

import "github.com/Kamva/mgm"

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

// func NewMock(name string, pages int) *Mock {
// 	return &Mock{
// 		Name:  name,
// 		Pages: pages,
// 	}
// }

type Mocks []Mock

type CustomHeader struct {
	Key   string
	Value string
}
