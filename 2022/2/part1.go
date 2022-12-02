package main

import (
	"fmt"
	"strings"
)

type Part1Solver struct {
	totalScore int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
	parts := strings.Split(l, " ")
	opponent := getHandFromLetter(parts[0])
	me := getHandFromLetter(parts[1])

	scoreForHand := 1 + int(me)

	fmt.Printf("score for hand: %d, ", scoreForHand)

	scoreForResult := 0
	switch r := compare(me, opponent); r {
	case 0:
		scoreForResult = 3
	case 1:
		scoreForResult = 6
	}

	fmt.Printf("+ score for result: %d, = ", scoreForResult)

	totalScore := scoreForHand + scoreForResult
	fmt.Printf("%d\n", totalScore)

	s.totalScore = s.totalScore + totalScore
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.totalScore)
}

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

func getHandFromLetter(l string) Hand {
	if l == "A" || l == "X" {
		return Rock
	} else if l == "B" || l == "Y" {
		return Paper
	} else { // C or Z
		return Scissors
	}
}

func compare(a, b Hand) int {
	if a == b {
		return 0
	}

	if a == Rock { // 0
		if b == Paper { // 1
			return -1
		}
		if b == Scissors { // 2
			return 1
		}
	}
	if a == Paper { // 1
		if b == Rock { // 0
			return 1
		}
		if b == Scissors { // 2
			return -1
		}
	}
	if a == Scissors { // 2
		if b == Rock { // 0
			return -1
		}
		if b == Paper { // 1
			return 1
		}
	}

	return -999 // invalid
}
