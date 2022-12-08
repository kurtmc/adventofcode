package main

import (
	"fmt"
	"strconv"
)

type Part2Solver struct {
	grid        [][]int
	columnIndex int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		columnIndex: 0,
	}
}

func (s *Part2Solver) Line(l string) {
	if s.grid == nil {
		s.grid = make([][]int, 0)
	}
	s.grid = append(s.grid, make([]int, len(l)))
	for i, v := range l {
		treeHeight, _ := strconv.Atoi(string(v))
		s.grid[s.columnIndex][i] = treeHeight

	}
	s.columnIndex++
}

func (s *Part2Solver) End() string {
	width := len(s.grid[0])
	height := len(s.grid)

	maxScenicScore := 0

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			score := scenicScore(s.grid, x, y)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	return fmt.Sprintf("%d", maxScenicScore)
}

func scenicScore(grid [][]int, x, y int) int {
	var fromBottom, fromTop, fromLeft, fromRight int = 0, 0, 0, 0
	treeHeight := grid[y][x]

	fmt.Printf("treeHeight: %d\n", treeHeight)

	for i := y - 1; i >= 0; i-- {
		fmt.Printf("fromTop:\n")
		fmt.Printf("grid[%d][%d]: %d\n", i, x, grid[i][x])
		if treeHeight > grid[i][x] {
			fromTop++
		} else {
			fromTop++
			break
		}
	}
	for i := y + 1; i < len(grid); i++ {
		fmt.Printf("fromBottom:\n")
		fmt.Printf("grid[%d][%d]: %d\n", i, x, grid[i][x])
		if treeHeight > grid[i][x] {
			fromBottom++
		} else {
			fromBottom++
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		if treeHeight > grid[y][i] {
			fromLeft++
		} else {
			fromLeft++
			break
		}
	}
	for i := x + 1; i < len(grid); i++ {
		if treeHeight > grid[y][i] {
			fromRight++
		} else {
			fromRight++
			break
		}
	}

	fmt.Printf("%d * %d * %d * %d = %d\n", fromBottom, fromTop, fromLeft, fromRight, fromBottom*fromTop*fromLeft*fromRight)
	fmt.Printf("expecting %d * %d * %d * %d = %d\n", 2, 1, 1, 2, 4)
	return fromBottom * fromTop * fromLeft * fromRight
}
