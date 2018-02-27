package day7

import (
	"AdventOfCode2017/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "day7/input"

type element struct {
	val      int
	children []string
}

type totWeightElement struct {
	elem      *element
	totWeight int
}

func parseChildren(line []string, m map[string]*element) []string {
	children := make([]string, 0)
	for i := 3; i < len(line); i++ {
		child := strings.TrimSuffix(line[i], ",")
		childElem, ok := m[child]
		if ok {
			// Child already in map
			children = append(children, child)
		} else {
			// Child not in map, lets create a one
			childElem = &element{}
			m[child] = childElem
			children = append(children, child)
		}
	}
	return children
}

func parse() map[string]*element {
	lines := helpers.ReadFileByLineAndSpace(fileName)
	m := make(map[string]*element)

	for _, line := range lines {

		// Parse the children
		children := parseChildren(line, m)
		name := line[0]
		// Removed "(" and ")" from the string
		valAsString := strings.TrimSuffix(strings.TrimPrefix(line[1], "("), ")")
		v, _ := strconv.Atoi(valAsString)

		if elem, ok := m[name]; ok {
			// Element is referenced to somewhere and already created
			elem.val = v
			elem.children = children
		} else {
			// Element doesn't exist so lets add it!
			newElem := &element{v, children}
			m[name] = newElem
		}
	}
	return m
}

func FirstPart() string {
	m := parse()
	childMap := make(map[string]bool)

	// Find all children and add them to the childMap
	for _, element := range m {
		for _, child := range element.children {
			childMap[child] = true
		}
	}

	//Check which element is not a child and return it
	for k := range m {
		if !childMap[k] {
			return k
		}
	}
	return "Found no bottom program, something is obviously wrong"
}

func SecondPart() int {
	m := parse()
	root := FirstPart()

	balanced, val := balanced(root, m)

	if balanced {
		fmt.Println("Found no inbalance")
		os.Exit(1)
	}
	return val
}

func balanced(node string, m map[string]*element) (bool, int) {
	n := m[node]

	children := make([]totWeightElement, 0)
	totWeight := n.val

	for _, child := range n.children {
		balanced, val := balanced(child, m)
		if !balanced {
			return false, val
		}
		newElement := totWeightElement{m[child], val}
		totWeight += val
		children = append(children, newElement)
	}

	uniqueElem, ok := findUniqueElem(children)

	if ok != nil {
		for _, child := range children {
			if uniqueElem.totWeight != child.totWeight {
				if uniqueElem.totWeight >= child.totWeight {
					return false, (uniqueElem.elem.val - (uniqueElem.totWeight - child.totWeight))
				}
				return false, (uniqueElem.elem.val - (child.totWeight - uniqueElem.totWeight))
			}
		}
		fmt.Println("Something seriously wrong")
		os.Exit(1)
	} else {
		//Balanced
		return true, totWeight
	}
	return false, 0
}

func findUniqueElem(children []totWeightElement) (totWeightElement, error) {
	m := make(map[int]int)
	for _, child := range children {
		m[child.totWeight] = m[child.totWeight] + 1
	}

	for _, val := range m {
		if val == 1 {
			for _, child := range children {
				if val == child.totWeight {
					return child, nil
				}
			}
		}
	}
	return totWeightElement{}, nil
}
