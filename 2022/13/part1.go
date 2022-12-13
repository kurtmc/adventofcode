package main

import "encoding/json"

type Part1Solver struct {
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
}

func (s *Part1Solver) End() string {
	return ""
}

func convert(input string) []interface{} {
	result := make([]interface{}, 0)
	json.Unmarshal([]byte(input), &result)
	return result
}
