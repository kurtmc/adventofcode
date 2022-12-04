package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part1Solver struct {
	fullyEnclosedCount int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
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

	if fullyEnclosed(elfARange, elfBRange) {
		s.fullyEnclosedCount = s.fullyEnclosedCount + 1
	}

	fmt.Printf("%d %d %d %d\n", elfAMin, elfAMax, elfBMin, elfBMax)
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.fullyEnclosedCount)
}

type Range struct {
	From int
	To   int
}

func fullyEnclosed(a, b Range) bool {
	if a.From >= b.From && a.To <= b.To {
		return true
	} else if b.From >= a.From && b.To <= a.To {
		return true
	}
	return false
}
