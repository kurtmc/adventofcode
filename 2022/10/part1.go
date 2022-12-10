package main

import (
	"fmt"
)

type Part1Solver struct {
	cpu                *CPU
	sumSignalStrengths int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		cpu:                NewCPU(),
		sumSignalStrengths: 0,
	}
}

func (s *Part1Solver) Line(l string) {
	s.cpu.LoadInstruction(l)
}

func (s *Part1Solver) End() string {
	s.cpu.Reset()
	//fmt.Printf("current instruction: %s\n", s.cpu.currentInstruction)
	//fmt.Printf("programCounter: %d\n", s.cpu.programCounter)
	//fmt.Printf("clock: %d\n", s.cpu.clock)
	//fmt.Printf("register_X: %d\n", s.cpu.register_X)
	//fmt.Printf("instructionClocksRemaining: %d\n", s.cpu.instructionClocksRemaining)

	for !s.cpu.Halt() {
		//fmt.Printf("clock: %d\n", s.cpu.clock)
		//fmt.Printf("register_X: %d\n", s.cpu.register_X)
		if isSignalStrengthClock(s.cpu.clock) {
			fmt.Printf("register_X: %d\n", s.cpu.register_X)
			fmt.Printf("signal strength: %d\n", (s.cpu.clock+1)*s.cpu.register_X)

			s.sumSignalStrengths = s.sumSignalStrengths + (s.cpu.clock+1)*s.cpu.register_X

		}

		fmt.Printf("clock: %d\n", s.cpu.clock)
		fmt.Printf("register_x: %d\n", s.cpu.register_X)

		s.cpu.AdvanceClock()

		//if isSignalStrengthClock(s.cpu.clock) {
		//	fmt.Printf("register_X: %d\n", s.cpu.register_X)
		//}

		//fmt.Printf("current instruction: %s\n", s.cpu.currentInstruction)
		//fmt.Printf("programCounter: %d\n", s.cpu.programCounter)
		//fmt.Printf("clock: %d\n", s.cpu.clock)
		//fmt.Printf("register_X: %d\n", s.cpu.register_X)
		//fmt.Printf("instructionClocksRemaining: %d\n", s.cpu.instructionClocksRemaining)
	}

	//return fmt.Sprintf("X: %d, clock: %d", s.cpu.register_X, s.cpu.clock)
	return fmt.Sprintf("%d", s.sumSignalStrengths)
}

func isSignalStrengthClock(clock int) bool {
	// 18, 58, 98, etc
	return ((clock+1)-20)%40 == 0
}
