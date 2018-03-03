package day9

import (
	"AdventOfCode2017/helpers"
	"errors"
	"fmt"
	"os"
)

const fileName = "day9/input"

type stack struct {
	s []int
}

func (st *stack) Push(i int) {
	st.s = append(st.s, i)
}

func (st *stack) Peek() (int, error) {
	len := len(st.s)
	if len == 0 {
		return 0, errors.New("Empty stack")
	}
	return st.s[len-1], nil
}

func (st *stack) Pop() (int, error) {
	len := len(st.s)
	if len == 0 {
		return 0, errors.New("Stack is empty")
	}

	ret := st.s[len-1]
	st.s = st.s[:len-1]
	return ret, nil
}

func (st *stack) Size() int {
	return len(st.s)
}

func FirstPart() int {
	temp := helpers.ReadFileByLine(fileName)
	if len(temp) > 1 || len(temp) == 0 {
		fmt.Println("Something is seriously wrong")
		os.Exit(1)
	}

	sum := 0
	stack := &stack{make([]int, 0)}

	for _, element := range temp[0] {

	}
	return sum
}
