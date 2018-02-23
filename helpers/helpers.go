package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
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

func ReadFirstLineWithSpaces(path string) []string {
	arr := ReadFileByLineAndSpace(path)
	return arr[0]
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

func ParseToInt(lines []string) []int {
	intLines := make([]int, len(lines), cap(lines))

	for index, line := range lines {
		element, _ := strconv.ParseInt(line, 10, 0)
		asInt := int(element)
		intLines[index] = asInt
	}
	return intLines
}
