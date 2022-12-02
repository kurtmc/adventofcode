package main

import (
	"fmt"
	"strings"
)

type Part2Solver struct {
	totalScore int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{}
}

func (s *Part2Solver) Line(l string) {
	parts := strings.Split(l, " ")
	opponent := getHandFromLetter(parts[0])

	scoreForResult := 0
	desiredResult := parts[1]
	var me Hand
	if desiredResult == "X" { // lose
		me = opponent - 1
		if me < 0 {
			me = me + 3
		}
	} else if desiredResult == "Y" { // draw
		scoreForResult = 3
		me = opponent

	} else { // win
		scoreForResult = 6
		me = (opponent + 1) % 3
	}

	scoreForHand := 1 + int(me)

	fmt.Printf("score for hand: %d, ", scoreForHand)

	switch r := compare(me, opponent); r {
	case 0:
		scoreForResult = 3
	case -1:
		scoreForResult = 6
	}

	fmt.Printf("+ score for result: %d, = ", scoreForResult)

	totalScore := scoreForHand + scoreForResult
	fmt.Printf("%d\n", totalScore)

	s.totalScore = s.totalScore + totalScore
}

func (s *Part2Solver) End() string {
	return fmt.Sprintf("%d", s.totalScore)
}
