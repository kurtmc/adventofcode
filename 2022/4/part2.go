package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part2Solver struct {
	overlapCount int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{}
}

func (s *Part2Solver) Line(l string) {
	parts := strings.Split(l, ",")
	elfA := strings.Split(parts[0], "-")
	elfB := strings.Split(parts[1], "-")

	elfARange := Range{}
	elfAMin, _ := strconv.Atoi(elfA[0])
	elfAMax, _ := strconv.Atoi(elfA[1])
	elfARange.From = elfAMin
	elfARange.To = elfAMax
	elfBMin, _ := strconv.Atoi(elfB[0])
	elfBMax, _ := strconv.Atoi(elfB[1])
	elfBRange := Range{}
	elfBRange.From = elfBMin
	elfBRange.To = elfBMax

	if overlap(elfARange, elfBRange) {
		s.overlapCount = s.overlapCount + 1

		//fmt.Printf("overlaps: %d %d %d %d\n", elfAMin, elfAMax, elfBMin, elfBMax)
	}

	//fmt.Printf("%d %d %d %d\n", elfAMin, elfAMax, elfBMin, elfBMax)
}

func (s *Part2Solver) End() string {
	return fmt.Sprintf("%d", s.overlapCount)
}

func overlap(a, b Range) bool {
	if a.From >= b.From && a.From <= b.To {
		return true
	}
	if a.To >= b.From && a.To <= b.To {
		return true
	}
	if b.From >= a.From && b.From <= a.To {
		return true
	}
	if b.To >= a.From && b.To <= a.To {
		return true
	}

	return false
}
