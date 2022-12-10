package main

import (
	"fmt"
	"strconv"
	"strings"
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
	//fmt.Printf("current instruction: %s\n", s.cpu.currentInstruction)
	//fmt.Printf("programCounter: %d\n", s.cpu.programCounter)
	//fmt.Printf("clock: %d\n", s.cpu.clock)
	//fmt.Printf("register_X: %d\n", s.cpu.register_X)
	//fmt.Printf("instructionClocksRemaining: %d\n", s.cpu.instructionClocksRemaining)

	s.cpu.currentInstruction = s.cpu.program[s.cpu.programCounter]

	for s.cpu.programCounter < len(s.cpu.program) {
		if isSignalStrengthClock(s.cpu.clock) {
			fmt.Printf("register_X: %d\n", s.cpu.register_X)
			fmt.Printf("signal strength: %d\n", (s.cpu.clock+2)*s.cpu.register_X)

			s.sumSignalStrengths = s.sumSignalStrengths + (s.cpu.clock+2)*s.cpu.register_X

		}

		s.cpu.Clock()

		if isSignalStrengthClock(s.cpu.clock) {
			fmt.Printf("register_X: %d\n", s.cpu.register_X)
		}

		//fmt.Printf("current instruction: %s\n", s.cpu.currentInstruction)
		//fmt.Printf("programCounter: %d\n", s.cpu.programCounter)
		//fmt.Printf("clock: %d\n", s.cpu.clock)
		//fmt.Printf("register_X: %d\n", s.cpu.register_X)
		//fmt.Printf("instructionClocksRemaining: %d\n", s.cpu.instructionClocksRemaining)
	}

	//return fmt.Sprintf("X: %d, clock: %d", s.cpu.register_X, s.cpu.clock)
	return fmt.Sprintf("%d", s.sumSignalStrengths)
}

type CPU struct {
	register_X     int
	clock          int
	program        []string
	programCounter int

	executing                  bool
	instructionClocksRemaining int
	currentInstruction         string
}

func isSignalStrengthClock(clock int) bool {
	// 18, 58, 98, etc
	return ((clock+2)-20)%40 == 0
}

func NewCPU() *CPU {
	return &CPU{
		register_X: 1,
		clock:      0,
		executing:  false,

		program:        make([]string, 0),
		programCounter: 0,
	}
}

func (c *CPU) LoadInstruction(instruction string) {
	c.program = append(c.program, instruction)
}

// clock the CPU
func (c *CPU) Clock() {
	if c.instructionClocksRemaining > 0 {
		c.instructionClocksRemaining--
		c.clock++
		return
	}
	//c.clock++

	//if c.instructionClocksRemaining == 0 {
	//	if strings.HasPrefix(c.currentInstruction, "addx") {
	//		parts := strings.Split(c.currentInstruction, " ")
	//		v, _ := strconv.Atoi(parts[1])
	//		c.register_X = c.register_X + v
	//	}

	//	c.programCounter++
	//	if c.programCounter < len(c.program) {
	//		c.currentInstruction = c.program[c.programCounter]
	//	}

	//	if c.currentInstruction == "noop" {
	//		c.instructionClocksRemaining = 0
	//		return
	//	}
	//	if strings.HasPrefix(c.currentInstruction, "addx") {
	//		c.instructionClocksRemaining = 1
	//		return
	//	}
	//} else {
	//	c.instructionClocksRemaining--
	//	return
	//}

	if c.currentInstruction == "noop" && c.instructionClocksRemaining == 0 {
		//c.clock++
		//c.instructionClocksRemaining = 0
	}
	if strings.HasPrefix(c.currentInstruction, "addx") && c.instructionClocksRemaining == 0 {
		//c.clock++
		//c.clock++
		parts := strings.Split(c.currentInstruction, " ")
		v, _ := strconv.Atoi(parts[1])
		c.register_X = c.register_X + v
	}

	c.programCounter++
	if c.programCounter < len(c.program) {
		c.currentInstruction = c.program[c.programCounter]
		if c.currentInstruction == "noop" {
			c.instructionClocksRemaining = 0
		} else if strings.HasPrefix(c.currentInstruction, "addx") {
			c.instructionClocksRemaining = 1
		}
	}

	c.clock++

}

//func (c *CPU) Execute(instruction string) {
//	if c.executing
//
//	c.currentInstruction = instruction
//
//	if instruction == "noop" {
//		c.instructionClockCycles = 1
//	} else if strings.HasPrefix(instruction, "addx") {
//		c.instructionClockCycles = 2
//
//		parts := strings.Split(instruction, " ")
//		v, _ := strconv.Atoi(parts[1])
//
//		for i := 0; i < 2; i++ {
//			c.clock++
//		}
//
//		c.register_X = c.register_X + v
//	}
//}
