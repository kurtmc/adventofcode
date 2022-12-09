package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Part1Solver struct {
	grid             [][]rune
	startingPosition *Point
	headPosition     *Point
	tailPosition     *Point
	tailPositions    []Point
}

func NewPart1Solver() *Part1Solver {
	grid := createGrid(5)
	grid[2][2] = 's'
	result := &Part1Solver{
		grid:             grid,
		startingPosition: NewPoint(1, 2),
		headPosition:     NewPoint(1, 2),
		tailPosition:     NewPoint(1, 2),
	}

	result.expandGrid(10)

	printGrid(result.grid)

	fmt.Println(result.headPosition)
	fmt.Println(result.tailPosition)
	fmt.Println(result.startingPosition)

	return result
}

func (s *Part1Solver) Line(l string) {
	instruction := strings.Split(l, " ")

	direction := instruction[0]
	count, _ := strconv.Atoi(instruction[1])

	for i := 0; i < count; i++ {
		s.moveHead(direction)
	}

}

func (s *Part1Solver) End() string {
	fmt.Println(s.tailPositions)
	for _, v := range s.tailPositions {
		s.grid[v.Y][v.X] = '#'
	}

	printGrid(s.grid)

	m := make(map[Point]bool, 0)
	for _, v := range s.tailPositions {
		m[v] = true
	}

	return fmt.Sprintf("%d", len(m)+1)
}

func isTouching(head, tail *Point) bool {
	result := (head.X-1) <= tail.X && tail.X <= (head.X+1) && (head.Y-1) <= tail.Y && tail.Y <= (head.Y+1)
	return result
}

func (s *Part1Solver) moveHead(direction string) {
	fmt.Printf("\nmove %s\n", direction)
	s.grid[s.headPosition.Y][s.headPosition.X] = '.'
	s.grid[s.tailPosition.Y][s.tailPosition.X] = '.'
	if direction == "R" || direction == "L" {
		horizontalTranslation := 1
		if direction == "L" {
			horizontalTranslation = -1
		}

		if len(s.grid[s.headPosition.Y])-s.headPosition.X == 1 {
			s.expandGrid(1)
		} else if (s.headPosition.X + horizontalTranslation) < 0 {
			s.expandGrid(1)
		}

		if *s.headPosition == *s.tailPosition {
			s.headPosition.X = s.headPosition.X + horizontalTranslation

		} else {
			s.headPosition.X = s.headPosition.X + horizontalTranslation
			if !isTouching(s.headPosition, s.tailPosition) {
				s.tailPosition.X = s.headPosition.X - horizontalTranslation
				s.tailPosition.Y = s.headPosition.Y
				s.tailPositions = append(s.tailPositions, *s.tailPosition)
			}
		}
	} else if direction == "U" || direction == "D" {
		verticalTranslation := 1
		if direction == "U" {
			verticalTranslation = -1
		}

		if len(s.grid)-s.headPosition.Y == 1 {
			s.expandGrid(1)
		} else if (s.headPosition.Y + verticalTranslation) < 0 {
			s.expandGrid(1)
		}

		if *s.headPosition == *s.tailPosition {
			s.headPosition.Y = s.headPosition.Y + verticalTranslation

		} else {
			s.headPosition.Y = s.headPosition.Y + verticalTranslation
			if !isTouching(s.headPosition, s.tailPosition) {
				fmt.Println("NOT TOUCHING")
				s.tailPosition.X = s.headPosition.X
				s.tailPosition.Y = s.headPosition.Y - verticalTranslation
				s.tailPositions = append(s.tailPositions, *s.tailPosition)
			}
		}
	}
	s.grid[s.startingPosition.Y][s.startingPosition.X] = 's'
	s.grid[s.tailPosition.Y][s.tailPosition.X] = 'T'
	s.grid[s.headPosition.Y][s.headPosition.X] = 'H'
}

func createGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]rune, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j] = '.'
		}
	}

	return grid
}

func expandGridHelper(grid [][]rune, amount int) [][]rune {
	newGrid := createGrid(len(grid) + amount*2)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			newGrid[i+amount][j+amount] = grid[i][j]
		}
	}

	return newGrid
}

func (s *Part1Solver) expandGrid(amount int) {
	newGrid := expandGridHelper(s.grid, amount)

	s.grid = newGrid
	s.startingPosition = NewPoint(s.startingPosition.X+amount, s.startingPosition.Y+amount)
	s.headPosition = NewPoint(s.headPosition.X+amount, s.headPosition.Y+amount)
	s.tailPosition = NewPoint(s.tailPosition.X+amount, s.tailPosition.Y+amount)

	for i := range s.tailPositions {
		s.tailPositions[i].X = s.tailPositions[i].X + amount
		s.tailPositions[i].Y = s.tailPositions[i].Y + amount
	}

}

func printGrid(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
