package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part2Solver struct {
	horizontal int
	depth      int
	aim        int
	step       int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{}
}

func (s *Part2Solver) Line(l string) {
	parts := strings.Split(l, " ")

	action := parts[0]
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	if action == "forward" {
		s.horizontal = s.horizontal + value
		s.depth = s.depth + s.aim*value
	} else if action == "down" {
		s.aim = s.aim + value
	} else if action == "up" {
		s.aim = s.aim - value
	} else {
		panic(fmt.Sprintf("unknown action: %s", action))
	}

	s.step++

}

func (s *Part2Solver) End() string {
	return fmt.Sprintf("%d", s.horizontal*s.depth)
}
