package main

import (
	"fmt"
	"strconv"
)

type Part2Solver struct {
	bitLength   int
	numberCount int
	bitCounts   []int
	numbers     []string
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		numbers: make([]string, 0),
	}
}

func (s *Part2Solver) Line(l string) {
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

	s.numbers = append(s.numbers, l)

}

func getBitDist(numbers []string) string {
	numberCount := len(numbers)
	bitCounts := make([]int, len(numbers[0]))

	for _, num := range numbers {
		for i, b := range num {
			if b == '1' {
				bitCounts[i] = bitCounts[i] + 1
			}
		}
	}

	gammaStr := ""
	for _, v := range bitCounts {
		if float64(v)/float64(numberCount) >= 0.5 {
			gammaStr = gammaStr + "1"
		} else {
			gammaStr = gammaStr + "0"
		}

	}

	return gammaStr
}

func invertStr(s string) string {
	result := ""

	for _, v := range s {
		if v == '1' {
			result = result + "0"
		} else {
			result = result + "1"
		}
	}

	return result
}

func getNumWithCommonBits(numbers []string, invert bool) int64 {
	i := 0
	l := make([]string, len(numbers))
	copy(l, numbers)
	gammaStr := getBitDist(l)
	for len(l) > 1 {
		gammaStr = getBitDist(l)
		if invert {
			gammaStr = invertStr(gammaStr)
		}
		newList := make([]string, 0)
		for _, v := range l {
			if v[i] == gammaStr[i] {
				newList = append(newList, v)
			}
		}
		l = newList
		i++
	}
	fmt.Println(l)

	result, err := strconv.ParseInt(l[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return result
}

func (s *Part2Solver) End() string {
	oxygenGeneratorRating := getNumWithCommonBits(s.numbers, false)
	co2ScrubberRating := getNumWithCommonBits(s.numbers, true)

	fmt.Println(oxygenGeneratorRating)
	fmt.Println(co2ScrubberRating)

	return fmt.Sprintf("%d", oxygenGeneratorRating*co2ScrubberRating)
}
