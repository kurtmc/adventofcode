package main

import "fmt"

type Part2Solver struct {
	groupSize   int
	group       []string
	prioritySum int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		groupSize: 3,
		group:     make([]string, 0),
	}
}

func (s *Part2Solver) Line(l string) {
	s.group = append(s.group, l)

	if len(s.group) == s.groupSize {
		common := findCommonLetters(s.group)[0]
		s.prioritySum = s.prioritySum + getPriority(common)
		s.group = make([]string, 0)
	}
}

func (s *Part2Solver) End() string {
	return fmt.Sprintf("%d", s.prioritySum)
}

func findCommonLetters(s []string) []rune {
	h := make(map[rune]int, 0)

	for _, v := range s[0] {
		h[v] = 1
	}

	for i := 2; i <= len(s); i++ {
		for _, v := range s[i-1] {
			if count, ok := h[v]; ok && count == i-1 {
				h[v] = i
			}
		}
	}

	result := make([]rune, 0)
	for r, v := range h {
		if v == len(s) {
			result = append(result, r)
		}
	}

	return result
}
