package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part2Solver struct {
	grid             [][]rune
	startingPosition *Point
	rope             *Knot
	tailPositions    []Point
}

type Knot struct {
	name             rune
	position         *Point
	previousPosition *Point
	nextKnot         *Knot
}

func (k *Knot) move(x, y int) {
	k.previousPosition = NewPoint(k.position.X, k.position.Y)
	k.position.X = x
	k.position.Y = y
}

func NewKnot(name rune, x, y int) *Knot {
	return &Knot{
		name:     name,
		position: NewPoint(x, y),
	}
}

func NewPart2Solver() *Part2Solver {
	grid := createGrid(1)

	ropeLength := 10

	rope := NewKnot('H', 0, 0)
	current := rope
	for i := 0; i < ropeLength-1; i++ {
		current.nextKnot = NewKnot(rune('1'+i), 0, 0)
		current = current.nextKnot
	}

	result := &Part2Solver{
		grid:             grid,
		startingPosition: NewPoint(0, 0),
		rope:             rope,
		tailPositions:    make([]Point, 0),
	}
	fmt.Print("#######################\n")
	fmt.Print("#######################\n")
	fmt.Print("#######################\n")
	fmt.Print("#######################\n")

	current = result.rope
	for current != nil {
		fmt.Printf("%c\n", current.name)
		current = current.nextKnot
	}
	fmt.Print("#######################\n")
	fmt.Print("#######################\n")

	result.tailPositions = append(result.tailPositions, Point{X: 0, Y: 0})

	result.expandGrid(10)

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
		s.moveHead(direction)
		s.PrintGridState()
	}

}

func (s *Part2Solver) PrintRopeState() {
	current := s.rope
	for current != nil {
		fmt.Printf("(%c, (%d, %d))", current.name, current.position.X, current.position.Y)
		if current.nextKnot != nil {
			fmt.Printf(", ")
		}
		current = current.nextKnot
	}

	fmt.Printf("\n")
}

func (s *Part2Solver) PrintGridState() {
	printGrid(s.grid)
	fmt.Printf("s.startingPosition: (x: %d, y: %d)\n", s.startingPosition.X, s.startingPosition.Y)
	fmt.Printf("s.rope: ")
	s.PrintRopeState()
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

func isDiagonal(head, tail *Point) bool {
	if !isTouching(head, tail) {
		return false
	}
	if head.X == tail.X {
		return false
	}
	if head.Y == tail.Y {
		return false
	}
	return true
}

func (s *Part2Solver) moveHead(direction string) {
	fmt.Println("header")
	current := s.rope
	for current != nil {
		s.grid[s.rope.position.Y][s.rope.position.X] = '.'

		current = current.nextKnot
	}

	// expand grid if the head of the rope is on the edge
	if len(s.grid[s.rope.position.Y])-s.rope.position.X == 1 || ((s.rope.position.X - 1) < 0) {
		s.expandGrid(1)
	} else if len(s.grid)-s.rope.position.Y == 1 || ((s.rope.position.Y - 1) < 0) {
		s.expandGrid(1)
	}

	translation := Point{X: 0, Y: 0}
	if direction == "R" {
		translation.X = 1
	} else if direction == "L" {
		translation.X = -1
	} else if direction == "U" {
		translation.Y = -1
	} else if direction == "D" {
		translation.Y = 1
	}

	// move the head of the rope
	head := s.rope
	head.move(head.position.X+translation.X, head.position.Y+translation.Y)

	current = head
	for current != nil {
		if current.nextKnot == nil {
			break
		}

		if !isTouching(current.position, current.nextKnot.position) {
			s.grid[current.nextKnot.position.Y][current.nextKnot.position.X] = '.'

			transX := current.position.X - current.nextKnot.position.X
			transY := current.position.Y - current.nextKnot.position.Y

			if transX > 1 {
				transX = 1
			}
			if transX < -1 {
				transX = -1
			}
			if transY > 1 {
				transY = 1
			}
			if transY < -1 {
				transY = -1
			}

			fmt.Printf("transX = %d, transY = %d\n", transX, transY)

			current.nextKnot.move(current.nextKnot.position.X+transX, current.nextKnot.position.Y+transY)
		}

		current = current.nextKnot
	}

	s.plotRope()
}

func (s *Part2Solver) plotRope() {
	s.grid[s.startingPosition.Y][s.startingPosition.X] = 's'

	knots := make([]*Knot, 0)

	current := s.rope
	for current != nil {
		knots = append([]*Knot{current}, knots...)

		current = current.nextKnot
	}

	for _, knot := range knots {
		s.grid[knot.position.Y][knot.position.X] = knot.name
	}
}

func (s *Part2Solver) expandGrid(amount int) {
	fmt.Println("expanding grid")
	newGrid := expandGridHelper(s.grid, amount)

	s.grid = newGrid
	s.startingPosition = NewPoint(s.startingPosition.X+amount, s.startingPosition.Y+amount)

	current := s.rope
	for current != nil {
		current.position.X = current.position.X + amount
		current.position.Y = current.position.Y + amount

		if current.previousPosition != nil {
			current.previousPosition.X = current.previousPosition.X + 1
			current.previousPosition.Y = current.previousPosition.Y + 1
		}

		current = current.nextKnot
	}

	for i := range s.tailPositions {
		s.tailPositions[i].X = s.tailPositions[i].X + amount
		s.tailPositions[i].Y = s.tailPositions[i].Y + amount
	}
}
