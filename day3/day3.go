package day3

import (
	"fmt"
)

const stepsTo = 25

type Pos struct {
	x, y int
}

func (pos *Pos) updatePos(x int, y int) {
	pos.x = pos.x + x
	pos.y = pos.y + y
}

func FirstPart() (int, int) {
	m := make(map[int]Pos)

	c := 1
	currentPos := Pos{x: 0, y: 0}
	s := 3
	m[c] = currentPos
	currentPos.updatePos(1, 0)
	c++
	fmt.Println(currentPos.x)
	fmt.Println(currentPos.y)
	fmt.Println("----------")
	for c <= stepsTo {
		// Up
		for i := 1; i < s; i++ {
			m[c] = currentPos
			currentPos.updatePos(0, -1)
			c++
		}
		move(&currentPos, -1, 1, &c)
		fmt.Println(currentPos.x)
		fmt.Println(currentPos.y)
		fmt.Println("----------")

		// Left
		for i := 1; i < s; i++ {
			m[c] = currentPos
			currentPos.updatePos(-1, 0)
			c++
		}
		move(&currentPos, 1, 1, &c)
		fmt.Println(currentPos.x)
		fmt.Println(currentPos.y)
		fmt.Println("----------")

		// Down
		for i := 1; i < s; i++ {
			m[c] = currentPos
			currentPos.updatePos(0, 1)
			c++
		}
		move(&currentPos, 1, -1, &c)
		fmt.Println(currentPos.x)
		fmt.Println(currentPos.y)
		fmt.Println("----------")

		// Right
		for i := 1; i < s; i++ {
			m[c] = currentPos
			currentPos.updatePos(1, 0)
			c++
		}

		fmt.Println(currentPos.x)
		fmt.Println(currentPos.y)
		fmt.Println("----------")
		fmt.Println(c)
		s += s + 2
	}
	return m[c].x, m[c].y
}

func move(pos *Pos, x int, y int, c *int) {
	pos.updatePos(x, y)
	*c = *c + 1
}
