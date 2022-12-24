package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Part2Solver struct {
	packets [][]interface{}
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		packets: make([][]interface{}, 0),
	}
}

func (s *Part2Solver) Line(l string) {
	if l != "" {
		s.packets = append(s.packets, convert(l))
	}
}

func (s *Part2Solver) End() string {
	s.packets = append(s.packets, convert("[[2]]"))
	s.packets = append(s.packets, convert("[[6]]"))

	sort.Slice(s.packets, func(i, j int) bool {
		cmp := compare(s.packets[i], s.packets[j])
		return cmp == -1
	})

	dividerIndexA, dividerIndexB := 0, 0
	for i, v := range s.packets {
		if reflect.DeepEqual(v, convert("[[2]]")) {
			dividerIndexA = 1 + i
		}
		if reflect.DeepEqual(v, convert("[[6]]")) {
			dividerIndexB = 1 + i
		}
		fmt.Println(v)
	}

	return fmt.Sprintf("%d", dividerIndexA*dividerIndexB)
}
