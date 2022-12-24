package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Part1Solver struct {
	a, b  string
	index int
	sum   int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{index: 1}
}

func (s *Part1Solver) Line(l string) {

	if l == "" {
		cmp := compare(convert(s.a), convert(s.b))
		fmt.Println(cmp)
		if cmp == 0 || cmp == -1 {
			s.sum = s.sum + s.index
		}

		s.index++
		s.a = ""
		s.b = ""
	} else {
		if s.a == "" {
			s.a = l
		} else {
			s.b = l
		}
	}
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.sum)
}

func convert(input string) []interface{} {
	result := make([]interface{}, 0)
	json.Unmarshal([]byte(input), &result)
	return result
}

func compare(a, b []interface{}) int {
	fmt.Printf("compare(%v, %v)\n", a, b)

	lenA := len(a)
	lenB := len(b)

	min := lenA
	if lenB < min {
		min = lenB
	}

	for i := 0; i < min; i++ {
		k := i
		v := a[i]
		if reflect.TypeOf(v) == reflect.TypeOf(b[k]) {
			if reflect.TypeOf(v).Kind() == reflect.Float64 {
				if v.(float64) < b[k].(float64) {
					return -1
				} else if v.(float64) > b[k].(float64) {
					return 1
				}
			} else {
				cmp := compare(v.([]interface{}), b[k].([]interface{}))
				if cmp != 0 {
					return cmp
				}
			}
		} else {
			var newA, newB []interface{}
			if reflect.TypeOf(v).Kind() == reflect.Float64 {
				newA = []interface{}{v}
				newB = b[k].([]interface{})
			} else {
				newA = v.([]interface{})
				newB = []interface{}{b[k]}
			}
			cmp := compare(newA, newB)
			if cmp != 0 {
				return cmp
			}
		}

	}

	if lenA < lenB {
		return -1
	}
	if lenA > lenB {
		return 1
	}
	return 0
}
