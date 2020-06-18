package mocks

import "fmt"

func ResponseHello(body string) string {
	s := fmt.Sprintf("%s Response", body)
	return s
}
