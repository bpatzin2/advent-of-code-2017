package main

import (
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
}

func (c coord) copy() coord {
	return coord{x: c.x, y: c.y}
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

func firstValueLargerThan(num int) int {
	squareSize := 1
	c := coord{x: 0, y: 0}
	spiral := spiralState{}
	v := 1
	vals := make(map[coord]int)
	vals[c] = v

	for i := 1; i < num; i++ {
		switch {
		case spiral.up > 0:
			spiral.up -= 1
			c = c.copy()
			c.y += 1
		case spiral.left > 0:
			spiral.left -= 1
			c = c.copy()
			c.x -= 1
		case spiral.down > 0:
			spiral.down -= 1
			c = c.copy()
			c.y -= 1
		case spiral.right > 0:
			spiral.right -= 1
			c = c.copy()
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

		v = sumAdjacents(c, vals)
		if v > num {
			return v
		}
		vals[c] = v

	}

	return abs(c.x) + abs(c.y)
}

func sumAdjacents(c coord, vals map[coord]int) int {
	xs := []int{-1, 0, 1}
	ys := []int{-1, 0, 1}
	sum := 0
	for _, x := range xs {
		for _, y := range ys {
			sum += vals[coord{x: c.x + x, y: c.y + y}]
		}
	}
	return sum
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	p1 := distanceFromCenter(325489)
	fmt.Println("Part 1:", p1)
	p2 := firstValueLargerThan(325489)
	fmt.Println("Part 2", p2)
}
