package day2

import (
	"AdventOfCode2017/helpers"
	"math"
	"strconv"
)

const fileName = "day2/input"

func PartOne() int {
	rows := helpers.ReadFileByLineAndSpace(fileName)

	sum := 0

	for _, row := range rows {
		var min, max int = math.MaxInt64, math.MinInt64

		for _, colEl := range row {
			i, _ := strconv.ParseInt(colEl, 10, 0)
			asInt := int(i)
			if asInt > max {
				max = asInt
			}
			if asInt < min {
				min = asInt
			}
		}
		sum += max - min
	}
	return sum
}
