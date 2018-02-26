package day7

import (
	"AdventOfCode2017/helpers"
	"strconv"
	"strings"
)

const fileName = "day7/input"

type element struct {
	name     string
	val      int
	children []*element
}

func parseChildren(line []string, m map[string]*element) []*element {
	children := make([]*element, 0)
	for i := 3; i < len(line); i++ {
		child := strings.TrimSuffix(line[i], ",")
		childElem, ok := m[child]
		if ok {
			// Child already in map
			children = append(children, childElem)
		} else {
			// Child not in map, lets create a one
			childElem = &element{}
			childElem.name = child
			m[child] = childElem
			children = append(children, childElem)
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
			newElem := &element{name, v, children}
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
			childMap[child.name] = true
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
