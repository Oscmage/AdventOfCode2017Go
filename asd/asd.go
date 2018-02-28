package asd

import (
	"fmt"
	"strconv"
)

func victory(A string, B string) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if value(A[i]) > value(B[i]) {
			count++
			fmt.Println(value(A[i]))
			fmt.Println(value(B[i]))
		}
	}
	return count
}

func value(a byte) int {
	switch a {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	}
	val, _ := strconv.Atoi(string(a))
	return val

}

func Solution(S string) int {
	val, _ := strconv.ParseInt(S, 2, 10)
	count := 0

	for val > 0 {
		if val%2 == 0 {
			val = val / 2
		} else {
			val--
		}
		count++
	}
	return count
}

func SolutionTwo(S string) int {
	count, removeZeroes := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] != '0' {
			break
		}
		removeZeroes++
	}
	S = S[removeZeroes:]

	for S != "0" && S != "" {
		lastBit := S[len(S)-1:]
		if lastBit == "1" {
			S = S[:len(S)-1] + "0"
		} else {
			S = S[:len(S)-1]
		}
		count++
	}
	return count
}
