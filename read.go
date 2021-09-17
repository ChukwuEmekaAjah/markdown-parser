package md

import (
	"io/ioutil"
	"strings"
)

func readFile(filename string) (string, error) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getFileLines(data string) []string {

	return strings.Split(data, "\n")
}

func ParseFile(filename string) ([]string, error) {

	data, err := readFile(filename)

	if err != nil {
		return []string{}, err
	}

	return getFileLines(data), nil
}
