package main

import (
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
}

type spiralState struct {
	up    int
	left  int
	down  int
	right int
}

func distanceFromCenter(num int) int {
	squareSize := 1
	c := coord{x: 0, y: 0}
	spiral := spiralState{}
	for i := 1; i < num; i++ {
		switch {
		case spiral.up > 0:
			spiral.up -= 1
			c.y += 1
		case spiral.left > 0:
			spiral.left -= 1
			c.x -= 1
		case spiral.down > 0:
			spiral.down -= 1
			c.y -= 1
		case spiral.right > 0:
			spiral.right -= 1
			c.x += 1
		default:
			c.x += 1
			squareSize += 2

			spiral = spiralState{
				up:    squareSize - 2,
				left:  squareSize - 1,
				down:  squareSize - 1,
				right: squareSize - 1,
			}
		}
	}

	return abs(c.x) + abs(c.y)
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	result := distanceFromCenter(325489)
	fmt.Println("Part 1:", result)
}
