package day5

import (
	"AdventOfCode2017/helpers"
)

const fileName = "day5/input"

/*
	First part of the question
*/
func FirstPart() int {
	return general(false)
}

func general(withWeirdCondition bool) int {
	var lines []int = helpers.ParseToInt(helpers.ReadFileByLine(fileName))

	steps, pointer := 0, 0
	for pointer < len(lines) {
		oldPointer := pointer
		pointer = pointer + lines[pointer]
		if withWeirdCondition && lines[oldPointer] >= 3 {
			lines[oldPointer]--
		} else {
			lines[oldPointer]++
		}
		steps++
	}
	return steps
}

/*
	Second part of the question
*/
func SecondPart() int {
	return general(true)
}
