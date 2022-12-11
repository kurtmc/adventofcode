package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Part2Solver struct {
	monkeys       []*Monkey
	currentMonkey *Monkey
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		monkeys: make([]*Monkey, 0),
	}
}

func (s *Part2Solver) Line(l string) {
	if strings.HasPrefix(l, "Monkey ") {
		parts := strings.Split(l, " ")
		name, _ := strconv.Atoi(strings.TrimSuffix(parts[1], ":"))
		s.currentMonkey = NewMonkey(name)
		s.monkeys = append(s.monkeys, s.currentMonkey)
	} else {
		if strings.HasPrefix(l, "  Starting items: ") {
			r := strings.TrimPrefix(l, "  Starting items: ")
			parts := strings.Split(r, ", ")
			for _, part := range parts {
				item, _ := strconv.ParseUint(part, 10, 64)
				s.currentMonkey.Items = append(s.currentMonkey.Items, item)
			}
		} else if strings.HasPrefix(l, "  Operation: ") {
			s.currentMonkey.SetOperation(strings.TrimPrefix(l, "  Operation: "))
		} else if strings.HasPrefix(l, "  Test: divisible by ") {
			s.currentMonkey.Test, _ = strconv.ParseUint(strings.TrimPrefix(l, "  Test: divisible by "), 10, 64)
		} else if strings.HasPrefix(l, "    If true: ") {
			r := strings.TrimPrefix(l, "    If true: throw to monkey ")
			s.currentMonkey.TestTrueMonkey, _ = strconv.Atoi(r)
		} else if strings.HasPrefix(l, "    If false: ") {
			r := strings.TrimPrefix(l, "    If false: throw to monkey ")
			s.currentMonkey.TestFalseMonkey, _ = strconv.Atoi(r)
		}
	}
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func (s *Part2Solver) End() string {

	for _, monkey := range s.monkeys {
		fmt.Println(monkey)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	rounds := 10000

	// find LCM for all monkey tests
	tests := []uint64{}
	for _, monkey := range s.monkeys {
		tests = append(tests, monkey.Test)
	}

	lcm := LCM(tests[0], tests[1], tests[2:]...)

	for i := 0; i < rounds; i++ {
		for _, monkey := range s.monkeys {
			fmt.Printf("Monkey %d:\n", monkey.Name)

			for monkey.HasItems() {
				item := monkey.GetItemForInspection()

				fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)

				newWorryLevel := monkey.Operation.Evalute(item)

				if newWorryLevel < item {
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
					fmt.Println("#####################")
				}

				fmt.Printf("    Worry level is operated and becomes %d.\n", newWorryLevel)

				newWorryLevel = newWorryLevel % lcm

				if newWorryLevel%monkey.Test == 0 {

					fmt.Printf("    Current worry level is divisible by %d.\n", monkey.Test)
					fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", newWorryLevel, monkey.TestTrueMonkey)
					s.monkeys[monkey.TestTrueMonkey].CatchItem(newWorryLevel)
				} else {

					fmt.Printf("    Current worry level is not divisible by %d.\n", monkey.Test)
					fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", newWorryLevel, monkey.TestFalseMonkey)
					s.monkeys[monkey.TestFalseMonkey].CatchItem(newWorryLevel)
				}
			}
			fmt.Println()
		}
	}

	for _, monkey := range s.monkeys {
		//fmt.Printf("Monkey %d: %v\n", monkey.Name, monkey.Items)
		fmt.Printf("Monkey %d inspected items %d times.\n", monkey.Name, monkey.ItemsInspected)
	}

	return fmt.Sprintf("%d", s.GetMonkeyBusiness())
}

func (s Part2Solver) GetMonkeyBusiness() int {
	inspectionCounts := make([]int, 0)
	for _, monkey := range s.monkeys {
		inspectionCounts = append(inspectionCounts, monkey.ItemsInspected)
	}
	sort.Ints(inspectionCounts)
	top := inspectionCounts[len(inspectionCounts)-2:]
	return top[0] * top[1]
}
