package mocks

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func LoadMock(fileName string) (string, error) {

	file, error := ioutil.ReadFile(fileName)
	if error != nil {
		return "", errors.New("file not found")
	}

	fmt.Printf("Loaded mock response:\n%s", string(file))

	return string(file), nil
}
