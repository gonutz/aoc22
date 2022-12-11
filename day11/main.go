package main

import (
	"fmt"
	"sort"
)

func main() {
	part1()
	part2()
}

func part1() {
	monkeys := copyInput()

	inspect := make(map[int]int)

	turn := func(monkeyIndex int) {
		m := &monkeys[monkeyIndex]
		for len(m.items) > 0 {
			inspect[monkeyIndex]++
			m.items[0] = m.op(m.items[0]) / 3
			target := m.target2
			if m.items[0]%m.div == 0 {
				target = m.target1
			}
			monkeys[target].items = append(monkeys[target].items, m.items[0])
			m.items = m.items[1:]
		}
	}

	round := func() {
		for i := range monkeys {
			turn(i)
		}
	}

	for i := 0; i < 20; i++ {
		round()
	}

	var counts []int
	for _, n := range inspect {
		counts = append(counts, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	fmt.Println("part 1)", counts[0]*counts[1])
}

func part2() {
	monkeys := copyInput()

	// If we stay in modulo space of all the factors that we check for, we
	// should be good. So we multiply them all up and modulo items by that
	// product.
	allDivs := 1
	for _, m := range monkeys {
		allDivs *= m.div
	}

	inspect := make(map[int]int)

	turn := func(monkeyIndex int) {
		m := &monkeys[monkeyIndex]
		for len(m.items) > 0 {
			inspect[monkeyIndex]++
			m.items[0] = m.op(m.items[0]) % allDivs
			target := m.target2
			if m.items[0]%m.div == 0 {
				target = m.target1
			}
			monkeys[target].items = append(monkeys[target].items, m.items[0])
			m.items = m.items[1:]
		}
	}

	round := func() {
		for i := range monkeys {
			turn(i)
		}
	}

	for i := 0; i < 10000; i++ {
		round()
	}

	var counts []int
	for _, n := range inspect {
		counts = append(counts, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	fmt.Println("part 2)", counts[0]*counts[1])
}

func copyInput() []monkey {
	m := make([]monkey, len(input))
	for i := range m {
		m[i] = input[i]
		m[i].items = append([]int{}, m[i].items...)
	}
	return m
}

type monkey struct {
	items   []int
	op      func(int) int
	div     int
	target1 int
	target2 int
}

var input = []monkey{
	{
		items:   []int{50, 70, 89, 75, 66, 66},
		op:      func(old int) int { return old * 5 },
		div:     2,
		target1: 2,
		target2: 1,
	},
	{
		items:   []int{85},
		op:      func(old int) int { return old * old },
		div:     7,
		target1: 3,
		target2: 6,
	},
	{
		items:   []int{66, 51, 71, 76, 58, 55, 58, 60},
		op:      func(old int) int { return old + 1 },
		div:     13,
		target1: 1,
		target2: 3,
	},
	{
		items:   []int{79, 52, 55, 51},
		op:      func(old int) int { return old + 6 },
		div:     3,
		target1: 6,
		target2: 4,
	},
	{
		items:   []int{69, 92},
		op:      func(old int) int { return old * 17 },
		div:     19,
		target1: 7,
		target2: 5,
	},
	{
		items:   []int{71, 76, 73, 98, 67, 79, 99},
		op:      func(old int) int { return old + 8 },
		div:     5,
		target1: 0,
		target2: 2,
	},
	{
		items:   []int{82, 76, 69, 69, 57},
		op:      func(old int) int { return old + 7 },
		div:     11,
		target1: 7,
		target2: 4,
	},
	{
		items:   []int{65, 79, 86},
		op:      func(old int) int { return old + 5 },
		div:     17,
		target1: 5,
		target2: 0,
	},
}
