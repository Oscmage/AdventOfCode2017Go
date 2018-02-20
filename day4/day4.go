package day4

import (
	"AdventOfCode2017/helpers"
	"sort"
	"strings"
)

const fileName = "day4/input"

func FirstPart() int {
	return validatePassphrase(true)
}

func validatePassphrase(allowAnagram bool) int {
	lines := helpers.ReadFileByLine(fileName)
	m := make(map[string]bool)
	count := 0

	for _, line := range lines {
		if validPassphrase(line, m, allowAnagram) {
			count++
		}
	}
	return count
}

func SecondPart() int {
	return validatePassphrase(false)
}

func validPassphrase(line string, m map[string]bool, allowAnagram bool) bool {
	col := strings.Fields(line)

	for _, el := range col {
		if !allowAnagram {
			s := strings.Split(el, "")
			sort.Strings(s)
			el = strings.Join(s, "")

		}
		if _, ok := m[el]; ok {
			return false
		}
		m[el] = true
	}
	return true
}
