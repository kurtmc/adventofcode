package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part2Solver struct {
	stacks []*Stack
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{}
}

func (s *Part2Solver) Line(l string) {
	if l == "" {
		fmt.Println("starting:")
		PrintStacks(s.stacks)
	}
	if strings.HasPrefix(l, " 1 ") || l == "" {
		return
	}

	if strings.HasPrefix(l, "move ") {
		parts := strings.Split(l, " ")
		count, _ := strconv.Atoi(parts[1])
		fromIndex, _ := strconv.Atoi(parts[3])
		fromIndex = fromIndex - 1
		toIndex, _ := strconv.Atoi(parts[5])
		toIndex = toIndex - 1

		crates := s.stacks[fromIndex].PopN(count)
		s.stacks[toIndex].PushN(crates)

		fmt.Printf("%s:\n\n", l)
		PrintStacks(s.stacks)
	} else {
		if s.stacks == nil {
			stackCount := (len(l) + 1) / 4
			s.stacks = make([]*Stack, 0)
			for i := 0; i < stackCount; i++ {
				s.stacks = append(s.stacks, NewStack())
			}

		}
		for i := 0; i < len(s.stacks); i++ {
			index := 1 + i*4
			if l[index] != " "[0] {
				s.stacks[i].stack = append([]byte{l[index]}, s.stacks[i].stack...)
			}
		}

	}
}

func (s *Part2Solver) End() string {
	result := ""
	fmt.Println("ending:")
	PrintStacks(s.stacks)
	//for i := 0; i < len(s.stacks); i++ {
	//	fmt.Printf("s.stacks[%d] = %c\n", i, s.stacks[i])
	//}
	for i := 0; i < len(s.stacks); i++ {
		result = result + string(s.stacks[i].Pop())

	}
	// MMZMZGVGS
	return result
}

func PrintStacks(stacks []*Stack) {
	maxLen := 0
	for _, s := range stacks {
		if len(s.stack) > maxLen {
			maxLen = len(s.stack)
		}
	}

	for i := maxLen - 1; i >= 0; i-- {
		for j := 0; j < len(stacks); j++ {
			if len(stacks[j].stack) > i {
				fmt.Printf("[%c]", stacks[j].stack[i])
			} else {
				fmt.Printf("   ")
			}
			if j != len(stacks)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	for j := 0; j < len(stacks); j++ {
		fmt.Printf(" %d ", j+1)
		if j != len(stacks)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}
