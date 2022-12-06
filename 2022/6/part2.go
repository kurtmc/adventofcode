package main

import "fmt"

type Part2Solver struct {
	result int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{}
}

func (s *Part2Solver) Line(l string) {
	s.result = StartOfMessage(l)
}

func (s *Part2Solver) End() string {
	return fmt.Sprintf("%d", s.result)
}

func StartOfMessage(i string) int {

	uniqLen := 14

	b := make([]rune, 0)

	for k, v := range i {
		b = append(b, v)
		if len(b) > uniqLen {
			b = b[1:]
		}

		uniq := make(map[rune]int)
		for _, r := range b {
			uniq[r] = 1
		}

		if len(uniq) == uniqLen {
			return k + 1
		}
	}

	return 0
}
