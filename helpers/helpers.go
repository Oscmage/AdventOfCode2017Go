package helpers

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	return string(file), err
}

func ReadFileByLine(path string) []string {
	file, err := ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}

func ReadFileByLineAndSpace(path string) [][]string {
	rows := ReadFileByLine(path)

	var matrix [][]string
	for _, element := range rows {
		rowSplitBySpace := strings.Fields(element)
		matrix = append(matrix, rowSplitBySpace)
	}

	return matrix
}
