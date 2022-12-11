//+build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`type monkey struct {
	items   []int
	op      func(int) int
	div     int
	target1 int
	target2 int
}
`)
	fmt.Println(`var input = []monkey {`)

	monkeys := strings.Split(input, "\n\n")
	for _, m := range monkeys {
		lines := strings.Split(m, "\n")
		items := strings.Replace(lines[1], "  Starting items: ", "", 1)
		op := strings.Replace(lines[2], "  Operation: new = ", "", 1)
		test := strings.Replace(lines[3], "  Test: divisible by ", "", 1)
		target1 := strings.Replace(lines[4], "    If true: throw to monkey ", "", 1)
		target2 := strings.Replace(lines[5], "    If false: throw to monkey ", "", 1)

		fmt.Println("\t{")
		fmt.Printf("\t\titems: []int{%s},\n", items)
		fmt.Printf("\t\top: func(old int) int { return %s },\n", op)
		fmt.Printf("\t\tdiv: %s,\n", test)
		fmt.Printf("\t\ttarget1: %s,\n", target1)
		fmt.Printf("\t\ttarget2: %s,\n", target2)
		fmt.Println("\t},")
	}

	fmt.Println("}")
}

const input = `Monkey 0:
  Starting items: 50, 70, 89, 75, 66, 66
  Operation: new = old * 5
  Test: divisible by 2
    If true: throw to monkey 2
    If false: throw to monkey 1

Monkey 1:
  Starting items: 85
  Operation: new = old * old
  Test: divisible by 7
    If true: throw to monkey 3
    If false: throw to monkey 6

Monkey 2:
  Starting items: 66, 51, 71, 76, 58, 55, 58, 60
  Operation: new = old + 1
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 79, 52, 55, 51
  Operation: new = old + 6
  Test: divisible by 3
    If true: throw to monkey 6
    If false: throw to monkey 4

Monkey 4:
  Starting items: 69, 92
  Operation: new = old * 17
  Test: divisible by 19
    If true: throw to monkey 7
    If false: throw to monkey 5

Monkey 5:
  Starting items: 71, 76, 73, 98, 67, 79, 99
  Operation: new = old + 8
  Test: divisible by 5
    If true: throw to monkey 0
    If false: throw to monkey 2

Monkey 6:
  Starting items: 82, 76, 69, 69, 57
  Operation: new = old + 7
  Test: divisible by 11
    If true: throw to monkey 7
    If false: throw to monkey 4

Monkey 7:
  Starting items: 65, 79, 86
  Operation: new = old + 5
  Test: divisible by 17
    If true: throw to monkey 5
    If false: throw to monkey 0`
