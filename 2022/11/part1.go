package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Part1Solver struct {
	monkeys       []*Monkey
	currentMonkey *Monkey
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		monkeys: make([]*Monkey, 0),
	}
}

func (s *Part1Solver) Line(l string) {
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
				item, _ := strconv.Atoi(part)
				s.currentMonkey.Items = append(s.currentMonkey.Items, item)
			}
		} else if strings.HasPrefix(l, "  Operation: ") {
			s.currentMonkey.SetOperation(strings.TrimPrefix(l, "  Operation: "))
		} else if strings.HasPrefix(l, "  Test: divisible by ") {
			s.currentMonkey.Test, _ = strconv.Atoi(strings.TrimPrefix(l, "  Test: divisible by "))
		} else if strings.HasPrefix(l, "    If true: ") {
			r := strings.TrimPrefix(l, "    If true: throw to monkey ")
			s.currentMonkey.TestTrueMonkey, _ = strconv.Atoi(r)
		} else if strings.HasPrefix(l, "    If false: ") {
			r := strings.TrimPrefix(l, "    If false: throw to monkey ")
			s.currentMonkey.TestFalseMonkey, _ = strconv.Atoi(r)
		}
	}
}

func (s *Part1Solver) End() string {

	for _, monkey := range s.monkeys {
		fmt.Println(monkey)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	rounds := 20

	for i := 0; i < rounds; i++ {
		for _, monkey := range s.monkeys {
			fmt.Printf("Monkey %d:\n", monkey.Name)

			for monkey.HasItems() {
				item := monkey.GetItemForInspection()

				fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
				newWorryLevel := monkey.Operation.Evalute(item)
				fmt.Printf("    Worry level is operated and becomes %d.\n", newWorryLevel)
				fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", newWorryLevel/3)
				newWorryLevel = newWorryLevel / 3

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

func (s Part1Solver) GetMonkeyBusiness() int {
	inspectionCounts := make([]int, 0)
	for _, monkey := range s.monkeys {
		inspectionCounts = append(inspectionCounts, monkey.ItemsInspected)
	}
	sort.Ints(inspectionCounts)
	top := inspectionCounts[len(inspectionCounts)-2:]
	return top[0] * top[1]
}
