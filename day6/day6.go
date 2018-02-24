package day6

import (
	"AdventOfCode2017/helpers"
	"bytes"
	"strconv"
)

const fileName = "day6/input"

func FirstPart() int {
	arr := helpers.ParseToInt(helpers.ReadFirstLineWithSpaces(fileName))
	m := make(map[string]bool)

	largestPointer, count := 0, 0

	for {
		count++
		m[toString(arr)] = true
		largestPointer = findLargestPos(arr)

		// Distribute
		largestVal := arr[largestPointer]
		arr[largestPointer] = 0
		internalPointer := largestPointer + 1
		for largestVal > 0 {
			if internalPointer == len(arr) {
				internalPointer = 0
			}
			arr[internalPointer]++
			internalPointer++
			largestVal--
		}
		if m[toString(arr)] {
			return count
		}
	}
}

// Should be redefined and not used, this can be done smarter
func findLargestPos(arr []int) int {
	largest := 0
	for pos, elem := range arr {
		if elem > arr[largest] {
			largest = pos
		}
	}
	return largest
}

func toString(arr []int) string {
	var buffer bytes.Buffer

	for _, elem := range arr {
		buffer.WriteString(strconv.Itoa(elem))
	}
	return buffer.String()
}
