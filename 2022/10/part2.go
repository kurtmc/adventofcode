package main

import (
	"fmt"
)

type Part2Solver struct {
	cpu *CPU
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		cpu: NewCPU(),
	}
}

func (s *Part2Solver) Line(l string) {
	s.cpu.LoadInstruction(l)
}

func (s *Part2Solver) End() string {
	s.cpu.Reset()

	if s.cpu.register_X >= -1 && s.cpu.register_X <= 1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}

	for !s.cpu.Halt() {
		if s.cpu.clock%40 == 0 {
			fmt.Printf("\n")
		}

		crtHorizontalPosition := s.cpu.clock % 40
		if s.cpu.register_X >= (crtHorizontalPosition-1) && s.cpu.register_X <= (crtHorizontalPosition+1) {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		s.cpu.AdvanceClock()
	}

	crtHorizontalPosition := s.cpu.clock % 40
	if s.cpu.register_X >= (crtHorizontalPosition-1) && s.cpu.register_X <= (crtHorizontalPosition+1) {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}

	fmt.Printf("\n")
	return fmt.Sprintf("")
}
