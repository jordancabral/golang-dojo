package mocks

import (
	"fmt"
	"io/ioutil"
)

func ResponseHello(fileName string) string {

	file, error := ioutil.ReadFile(fileName)
	if error != nil {
		panic(error)
	}

	fmt.Println(file)
	fmt.Printf("json %s", string(file))

	return string(file)
}
