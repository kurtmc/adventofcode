package main

import (
	"fmt"
	"strconv"
)

type Part1Solver struct {
	bitLength   int
	numberCount int
	bitCounts   []int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
	if s.bitLength == 0 {
		s.bitLength = len(l)
	}

	if s.bitCounts == nil {
		s.bitCounts = make([]int, s.bitLength)
	}

	for i, v := range []byte(l) {
		if v == '1' {
			s.bitCounts[i] = s.bitCounts[i] + 1
		}
	}

	s.numberCount++

}

func (s *Part1Solver) End() string {
	gammaStr := ""
	epsilonStr := ""
	for _, v := range s.bitCounts {
		if v/(s.numberCount/2) >= 1 {
			gammaStr = gammaStr + "1"
			epsilonStr = epsilonStr + "0"
		} else {
			gammaStr = gammaStr + "0"
			epsilonStr = epsilonStr + "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%d", gamma*epsilon)
}
