package main

import (
	"fmt"
	"strconv"
)

type Part1Solver struct {
	previous int
	max      int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		previous: 0,
		max:      0,
	}
}

func (s *Part1Solver) Line(l string) {
	if l == "" {
		if s.previous > s.max {
			s.max = s.previous
		}
		s.previous = 0
	} else {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		//fmt.Println(i)
		s.previous = s.previous + i
	}
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.max)
}
