package main

import (
	"fmt"
	"strconv"
)

type Part1Solver struct {
	grid        [][]int
	columnIndex int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		columnIndex: 0,
	}
}

func (s *Part1Solver) Line(l string) {
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

func (s *Part1Solver) End() string {
	width := len(s.grid[0])
	height := len(s.grid)

	// all perimeter trees are visible
	visibleTreesCount := 2*height + (2 * width) - 4

	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if isVisible(s.grid, x, y) {
				visibleTreesCount++
			}
		}
	}

	//fmt.Println(s.grid)
	return fmt.Sprintf("%d", visibleTreesCount)
}

func isVisible(grid [][]int, x, y int) bool {
	var fromBottom, fromTop, fromLeft, fromRight bool = true, true, true, true
	treeHeight := grid[y][x]

	for i := y - 1; i >= 0; i-- {
		if treeHeight <= grid[i][x] {
			fromTop = false
		}
	}
	for i := y + 1; i < len(grid); i++ {
		if treeHeight <= grid[i][x] {
			fromBottom = false
		}
	}

	for i := x - 1; i >= 0; i-- {
		if treeHeight <= grid[y][i] {
			fromLeft = false
		}
	}
	for i := x + 1; i < len(grid); i++ {
		if treeHeight <= grid[y][i] {
			fromRight = false
		}
	}

	return fromBottom || fromTop || fromLeft || fromRight
}
