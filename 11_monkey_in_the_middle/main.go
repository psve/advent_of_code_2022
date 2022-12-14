package main

import (
	"fmt"
	"helper"
	"sort"
	"strconv"
	"strings"
)

const (
	tokenStartingItems = "  Starting items: "
	tokenOperation     = "  Operation: new = old "
	tokenTest          = "  Test: divisible by "
	tokenIfTrue        = "    If true: throw to monkey "
	tokenIfFalse       = "    If false: throw to monkey "
)

// Multiple of all the monkies' divisors. We do all calculations modulo this
// number to keep the values bounded. This works since all the divisors are co-prime.
var modulo int

type monkey struct {
	items     []int
	op        func(old int) int
	test      func(item int) int
	inspected int
}

func (m *monkey) throwItem(decreaseWorry bool, monkeys []*monkey) {
	m.inspected++
	item := m.op(m.items[0]) % modulo
	if decreaseWorry {
		item /= 3
	}
	m.items = m.items[1:]
	target := m.test(item)
	monkeys[target].items = append(monkeys[target].items, item)
}

func (m *monkey) throwAll(decreaseWorry bool, monkeys []*monkey) {
	for len(m.items) != 0 {
		m.throwItem(decreaseWorry, monkeys)
	}
}

func parse(path string) []*monkey {
	modulo = 1
	monkeys := make([]*monkey, 0)
	template := struct {
		items, op, test, ifTrue, ifFalse string
	}{}
	toMonkey := func() *monkey {
		m := &monkey{}

		// Parse the item list
		for _, i := range strings.Split(template.items, ", ") {
			item, _ := strconv.Atoi(i)
			m.items = append(m.items, item)
		}

		// Parse the operation
		parts := strings.Split(template.op, " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "*":
			switch parts[1] {
			case "old":
				m.op = func(old int) int { return old * old }
			default:
				m.op = func(old int) int { return old * val }
			}
		case "+":
			switch parts[1] {
			case "old":
				m.op = func(old int) int { return old + old }
			default:
				m.op = func(old int) int { return old + val }
			}
		}

		// Parse the test
		divisor, _ := strconv.Atoi(template.test)
		ifTrue, _ := strconv.Atoi(template.ifTrue)
		ifFalse, _ := strconv.Atoi(template.ifFalse)
		m.test = func(item int) int {
			if item%divisor == 0 {
				return ifTrue
			}
			return ifFalse
		}

		modulo *= divisor
		return m
	}

	// Go over each line, filling out the template. Once we parse the last line,
	// call 'toMonkey'.
	helper.ForEachLine(path, func(line string) error {
		switch {
		case strings.HasPrefix(line, tokenStartingItems):
			template.items = strings.TrimPrefix(line, tokenStartingItems)
		case strings.HasPrefix(line, tokenOperation):
			template.op = strings.TrimPrefix(line, tokenOperation)
		case strings.HasPrefix(line, tokenTest):
			template.test = strings.TrimPrefix(line, tokenTest)
		case strings.HasPrefix(line, tokenIfTrue):
			template.ifTrue = strings.TrimPrefix(line, tokenIfTrue)
		case strings.HasPrefix(line, tokenIfFalse):
			template.ifFalse = strings.TrimPrefix(line, tokenIfFalse)
			monkeys = append(monkeys, toMonkey())
		}
		return nil
	})

	return monkeys
}

func monkeyBusiness(path string, decreaseWorry bool, rounds int) int {
	monkeys := parse(path)
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			m.throwAll(decreaseWorry, monkeys)
		}
	}

	activity := make(sort.IntSlice, 0, len(monkeys))
	for _, m := range monkeys {
		activity = append(activity, m.inspected)
	}
	activity.Sort()
	return activity[len(activity)-1] * activity[len(activity)-2]
}

func main() {
	fmt.Println(monkeyBusiness("./input", true, 20))
	fmt.Println("---------")
	fmt.Println(monkeyBusiness("./input", false, 10000))
}
