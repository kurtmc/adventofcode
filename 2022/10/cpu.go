package main

import (
	"strconv"
	"strings"
)

type CPU struct {
	// Instruction-pointer
	ip         int
	register_X int
	clock      int
	program    []string

	instructionClocksRemaining int
	currentInstruction         string
}

func NewCPU() *CPU {
	return &CPU{
		program: make([]string, 0),
	}
}

func (c *CPU) Reset() {
	// Reset register
	c.register_X = 1

	// Reset clock
	c.clock = 1

	// Reset instruction pointer to zero.
	c.ip = 0

	// load initial instruction
	c.currentInstruction = c.program[c.ip]

}

func (c *CPU) LoadInstruction(instruction string) {
	c.program = append(c.program, instruction)
}

func (c *CPU) ExecuteInstruction() {
	if c.currentInstruction == "noop" {
	} else if strings.HasPrefix(c.currentInstruction, "addx") {
		parts := strings.Split(c.currentInstruction, " ")
		v, _ := strconv.Atoi(parts[1])
		c.register_X = c.register_X + v
	}
}

func (c *CPU) AdvanceClock() {

	c.clock++

	if c.instructionClocksRemaining == 0 {
		c.ExecuteInstruction()

		// load next instruction
		c.ip++
		c.currentInstruction = c.program[c.ip]
		if c.currentInstruction == "noop" {
			c.instructionClocksRemaining = 1
			c.instructionClocksRemaining--
			c.ExecuteInstruction()
		} else if strings.HasPrefix(c.currentInstruction, "addx") {
			c.instructionClocksRemaining = 2
			c.instructionClocksRemaining--
		}
	} else {
		c.instructionClocksRemaining--
	}

}

func (c *CPU) Halt() bool {
	return c.instructionClocksRemaining == 0 && (len(c.program)-1) == c.ip
}
