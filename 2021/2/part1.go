package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part1Solver struct {
	horizontal int
	depth      int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		horizontal: 0,
		depth:      0,
	}
}

func (s *Part1Solver) Line(l string) {
	parts := strings.Split(l, " ")

	action := parts[0]
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	if action == "forward" {
		s.horizontal = s.horizontal + value
	} else if action == "down" {
		s.depth = s.depth + value
	} else if action == "up" {
		s.depth = s.depth - value
	} else {
		panic(fmt.Sprintf("unknown action: %s", action))
	}
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.horizontal*s.depth)
}
