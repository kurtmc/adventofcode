package main

import (
	"fmt"
	"strconv"
)

type Part1Solver struct {
	previous       int
	increasesCount int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		previous:       -1,
		increasesCount: 0,
	}
}

func (s *Part1Solver) Line(l string) {
	i, err := strconv.Atoi(l)

	if err != nil {
		panic(err)
	}

	if s.previous != -1 {
		if i > s.previous {
			s.increasesCount++
		}
	}

	s.previous = i
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.increasesCount)
}
