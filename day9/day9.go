package day9

import (
	"AdventOfCode2017/helpers"
	"fmt"
	"os"
)

const fileName = "day9/input"

func FirstPart() int {
	sum := 0
	sortOfAStack, pointer := make([]string, 0), 0
	temp := helpers.ReadFileByLine(fileName)
	if len(temp) > 1 || len(temp) == 0 {
		fmt.Println("Something is seriously wrong")
		os.Exit(1)
	}

	for _, element := range temp[0] {

	}
	return sum
}
