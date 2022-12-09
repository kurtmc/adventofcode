package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part2Solver struct {
	grid             [][]rune
	startingPosition *Point
	rope             []*Point
	tailPositions    []Point
}

func NewPart2Solver() *Part2Solver {
	grid := createGrid(1)

	ropeLength := 3

	rope := make([]*Point, 0)
	for i := 0; i < ropeLength; i++ {
		rope = append(rope, NewPoint(0, 0))
	}

	result := &Part2Solver{
		grid:             grid,
		startingPosition: NewPoint(0, 0),
		rope:             rope,
		tailPositions:    make([]Point, 0),
	}

	result.tailPositions = append(result.tailPositions, Point{X: 0, Y: 0})

	return result
}

func (s *Part2Solver) Line(l string) {
	instruction := strings.Split(l, " ")

	direction := instruction[0]
	count, _ := strconv.Atoi(instruction[1])

	s.plotRope()
	fmt.Println("initial state:")
	s.PrintGridState()

	for i := 0; i < count; i++ {
		fmt.Printf("\nmove %s\n", direction)
		s.moveHead2(direction)
		s.PrintGridState()
	}

}

func (s *Part2Solver) PrintGridState() {
	printGrid(s.grid)
	fmt.Printf("s.startingPosition: (x: %d, y: %d)\n", s.startingPosition.X, s.startingPosition.Y)
	fmt.Printf("s.rope: ")
	for i, v := range s.rope {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("(x: %d, y: %d)", v.X, v.Y)
	}
	fmt.Printf("\n")
}

func (s *Part2Solver) End() string {
	//fmt.Println(s.tailPositions)

	printGrid(s.grid)

	m := make(map[Point]bool, 0)
	for _, v := range s.tailPositions {
		m[v] = true
	}

	return fmt.Sprintf("%d", len(m))
}

func (s *Part2Solver) moveHead2(direction string) {

	headIndex := len(s.rope) - 1
	tailIndex := headIndex - 1

	for _, knot := range s.rope {
		s.grid[knot.Y][knot.X] = '.'
	}
	if direction == "R" || direction == "L" {
		horizontalTranslation := 1
		if direction == "L" {
			horizontalTranslation = -1
		}

		if len(s.grid[s.rope[headIndex].Y])-s.rope[headIndex].X == 1 {
			s.expandGrid(1)
		} else if (s.rope[headIndex].X + horizontalTranslation) < 0 {
			s.expandGrid(1)
		}

		if *s.rope[headIndex] == *s.rope[tailIndex] {
			s.rope[headIndex].X = s.rope[headIndex].X + horizontalTranslation
		} else {
			s.rope[headIndex].X = s.rope[headIndex].X + horizontalTranslation
			for i := len(s.rope) - 1; i >= 1; i-- {
				headIndex := i
				tailIndex := headIndex - 1
				if !isTouching(s.rope[headIndex], s.rope[tailIndex]) {
					s.rope[tailIndex].X = s.rope[headIndex].X - horizontalTranslation
					s.rope[tailIndex].Y = s.rope[headIndex].Y
				}
			}
		}
	} else if direction == "U" || direction == "D" {
		verticalTranslation := 1
		if direction == "U" {
			verticalTranslation = -1
		}

		if len(s.grid)-s.rope[headIndex].Y == 1 {
			s.expandGrid(1)
		} else if (s.rope[headIndex].Y + verticalTranslation) < 0 {
			s.expandGrid(1)
		}

		if *s.rope[headIndex] == *s.rope[tailIndex] {
			s.rope[headIndex].Y = s.rope[headIndex].Y + verticalTranslation

		} else {
			s.rope[headIndex].Y = s.rope[headIndex].Y + verticalTranslation
			for i := len(s.rope) - 1; i >= 1; i-- {
				headIndex := i
				tailIndex := headIndex - 1
				if !isTouching(s.rope[headIndex], s.rope[tailIndex]) {
					s.rope[tailIndex].X = s.rope[headIndex].X
					s.rope[tailIndex].Y = s.rope[headIndex].Y - verticalTranslation
				}
			}
		}
	}

	s.plotRope()
}

func (s *Part2Solver) moveHead(direction string) {

	headIndex := 1
	tailIndex := 0

	for _, knot := range s.rope {
		s.grid[knot.Y][knot.X] = '.'
	}
	if direction == "R" || direction == "L" {
		horizontalTranslation := 1
		if direction == "L" {
			horizontalTranslation = -1
		}

		if len(s.grid[s.rope[headIndex].Y])-s.rope[headIndex].X == 1 {
			s.expandGrid(1)
		} else if (s.rope[headIndex].X + horizontalTranslation) < 0 {
			s.expandGrid(1)
		}

		if *s.rope[headIndex] == *s.rope[tailIndex] {
			s.rope[1].X = s.rope[1].X + horizontalTranslation

		} else {
			s.rope[headIndex].X = s.rope[headIndex].X + horizontalTranslation
			if !isTouching(s.rope[headIndex], s.rope[tailIndex]) {
				s.rope[tailIndex].X = s.rope[headIndex].X - horizontalTranslation
				s.rope[tailIndex].Y = s.rope[headIndex].Y

				s.tailPositions = append(s.tailPositions, *s.rope[tailIndex])
			}
		}
	} else if direction == "U" || direction == "D" {
		verticalTranslation := 1
		if direction == "U" {
			verticalTranslation = -1
		}

		if len(s.grid)-s.rope[headIndex].Y == 1 {
			s.expandGrid(1)
		} else if (s.rope[headIndex].Y + verticalTranslation) < 0 {
			s.expandGrid(1)
		}

		if *s.rope[headIndex] == *s.rope[tailIndex] {
			s.rope[headIndex].Y = s.rope[headIndex].Y + verticalTranslation

		} else {
			s.rope[headIndex].Y = s.rope[headIndex].Y + verticalTranslation
			if !isTouching(s.rope[headIndex], s.rope[tailIndex]) {
				fmt.Println("NOT TOUCHING")
				s.rope[tailIndex].X = s.rope[headIndex].X
				s.rope[tailIndex].Y = s.rope[headIndex].Y - verticalTranslation

				s.tailPositions = append(s.tailPositions, *s.rope[tailIndex])
			}
		}
	}

	s.plotRope()
}

func (s *Part2Solver) plotRope() {
	ropeRunes := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'H'}

	s.grid[s.startingPosition.Y][s.startingPosition.X] = 's'

	for i := 0; i < len(s.rope); i++ {
		j := len(ropeRunes) - len(s.rope) + i
		//fmt.Printf("i = %d, j = %d, ropeRunes[%d] = %c\n", i, j, j, ropeRunes[j])
		s.grid[s.rope[i].Y][s.rope[i].X] = ropeRunes[j]
	}
}

func (s *Part2Solver) expandGrid(amount int) {
	fmt.Println("expanding grid")
	newGrid := expandGridHelper(s.grid, amount)

	s.grid = newGrid
	s.startingPosition = NewPoint(s.startingPosition.X+amount, s.startingPosition.Y+amount)

	for i := range s.rope {
		s.rope[i] = NewPoint(s.rope[i].X+amount, s.rope[i].Y+amount)
	}

	for i := range s.tailPositions {
		s.tailPositions[i].X = s.tailPositions[i].X + amount
		s.tailPositions[i].Y = s.tailPositions[i].Y + amount
	}
}
