package main

import "fmt"

type Part1Solver struct {
	result int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
	s.result = StartOfPacket(l)
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.result)
}

func StartOfPacket(i string) int {

	b := make([]rune, 0)

	for k, v := range i {
		b = append(b, v)
		if len(b) > 4 {
			b = b[1:]
		}

		uniq := make(map[rune]int)
		for _, r := range b {
			uniq[r] = 1
		}

		if len(uniq) == 4 {
			return k + 1
		}
	}

	return 0
}
