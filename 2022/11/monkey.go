package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Monkey 0:
//  Starting items: 79, 98
//  Operation: new = old * 19
//  Test: divisible by 23
//    If true: throw to monkey 2
//    If false: throw to monkey 3

type Monkey struct {
	Name      int
	Items     []uint64
	Operation *Operation
	Test      uint64 // divisible by this number

	TestTrueMonkey  int
	TestFalseMonkey int

	ItemsInspected int
}

func (m *Monkey) HasItems() bool {
	return len(m.Items) > 0
}

func (m *Monkey) GetItemForInspection() uint64 {
	m.ItemsInspected++

	var item uint64
	item, m.Items = m.Items[0], m.Items[1:]
	return item
}

func (m *Monkey) CatchItem(item uint64) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) String() string {

	startingItems := ""
	for _, v := range m.Items {
		startingItems = fmt.Sprintf("%s%d, ", startingItems, v)
	}

	startingItems = strings.TrimSuffix(startingItems, ", ")

	return fmt.Sprintf(`Monkey %d:
  Starting items: %s
  Operation: %s
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`, m.Name, startingItems, m.Operation, m.Test, m.TestTrueMonkey, m.TestFalseMonkey)
}

func NewMonkey(name int) *Monkey {
	return &Monkey{
		Name:           name,
		ItemsInspected: 0,
	}
}

func (m *Monkey) SetOperation(operation string) {
	m.Operation = NewOperation(operation)
}

type Operation struct {
	LHS      string
	RHS      string
	Operator string
}

func NewOperation(operation string) *Operation {
	parts := strings.Split(operation, " ")

	return &Operation{
		LHS:      parts[2],
		Operator: parts[3],
		RHS:      parts[4],
	}
}

func (o *Operation) Evalute(old uint64) uint64 {
	var lhs uint64 = 0
	var rhs uint64 = 0

	if o.LHS == "old" {
		lhs = old
	} else {
		lhs, _ = strconv.ParseUint(o.LHS, 10, 64)
	}
	if o.RHS == "old" {
		rhs = old
	} else {
		rhs, _ = strconv.ParseUint(o.RHS, 10, 64)
	}

	switch o.Operator {
	case "+":
		return lhs + rhs
	case "*":
		return lhs * rhs
	default:
		panic(fmt.Sprintf("unknown operator: %v", o.Operator))
	}

}

func (o *Operation) String() string {
	return fmt.Sprintf("new = %s %s %s", o.LHS, o.Operator, o.RHS)
}
