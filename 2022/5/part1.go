package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part1Solver struct {
	stacks []*Stack
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
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

		fmt.Printf("move %d from index %d to index %d\n", count, fromIndex, toIndex)

		for i := 0; i < count; i++ {
			crate := s.stacks[fromIndex].Pop()
			s.stacks[toIndex].Push(crate)
		}

	} else if l == "" {
		for i := 0; i < len(s.stacks); i++ {
			fmt.Printf("stack[%d] = %c\n", i, s.stacks[i])
		}
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
			fmt.Printf("%c\n", l[index])

			if l[index] != " "[0] {

				s.stacks[i].stack = append([]byte{l[index]}, s.stacks[i].stack...)
			}
		}

	}

}

func (s *Part1Solver) End() string {
	fmt.Println(s.stacks)
	result := ""
	for i := 0; i < len(s.stacks); i++ {
		result = result + string(s.stacks[i].Pop())

	}
	return result
}

type Stack struct {
	stack []byte
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]byte, 0),
	}
}

func (s *Stack) Push(r byte) {
	s.stack = append(s.stack, r)
}
func (s *Stack) Pop() byte {
	element := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return element
}
