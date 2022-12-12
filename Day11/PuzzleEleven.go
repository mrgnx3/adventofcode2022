package main

import (
	"2022/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*
Puzzle input
notes

monkeys react to level of worry

Starting Item := worry level
Operation := how worry level changes monkey behaviour
Test := if else what happens the item

worry level set after inspection but before test
worry level % 3 rounded down to whole number

1 turn each monkey inspects and throws every item one at a time in order listed
called a round
new items go to the end of the list


read input
generate monkeys, round
simulate round
count how many times each monkey inspects an item
after 20 round find two most active monkeys


structs:
	monkeys: starting items, operation, test
	round: monkeys

find 2 most active monkeys, * monkey business
*/

type monkey struct {
	StartingItems []int
	Operation     func(int) int
	Test          func(int) *monkey
	Count         int
	Divisor       int
}

func buildMonkeys(input []string) []*monkey {
	var monkeys []*monkey
	monkeyNumber := 0
	for _, line := range input {
		if strings.Contains(line, "Monkey") {
			monkeys = append(monkeys, &monkey{
				Count: 0,
			})
		} else if line == "" {
			monkeyNumber += 1
		}
	}
	monkeyNumber = 0
	for idx, line := range input {
		if strings.Contains(line, "Starting items") {
			monkeys[monkeyNumber].StartingItems = getStartingItems(line)
		} else if strings.Contains(line, "Operation") {
			monkeys[monkeyNumber].Operation = getOperation(line)
		} else if strings.Contains(line, "Test") {
			monkeys[monkeyNumber].Test = getTest(monkeys, input[idx], input[idx+1], input[idx+2], monkeys[monkeyNumber])
		} else if line == "" {
			monkeyNumber += 1
		}
	}
	return monkeys
}

func getTest(monkeys []*monkey, line1 string, line2 string, line3 string, mnky *monkey) func(int) *monkey {
	divisibleBy, _ := strconv.Atoi(strings.Split(line1, "divisible by ")[1])
	monkeyTrue, _ := strconv.Atoi(strings.Split(line2, "throw to monkey ")[1])
	monkeyFalse, _ := strconv.Atoi(strings.Split(line3, "throw to monkey ")[1])
	mnky.Divisor = divisibleBy
	return func(x int) *monkey {
		if x%divisibleBy == 0 {
			return monkeys[monkeyTrue]
		} else {
			return monkeys[monkeyFalse]
		}
	}
}

func getOperation(line string) func(int) int {
	op := strings.Split(strings.TrimSpace(line), "=")[1]
	switch {
	case strings.Contains(op, "*"):
		mult := strings.Split(strings.TrimSpace(op), "*")
		switch {
		case strings.TrimSpace(mult[0]) == "old" && strings.TrimSpace(mult[1]) == "old":
			return func(i int) int {
				return i * i
			}
		case strings.TrimSpace(mult[0]) == "old":
			num, _ := strconv.Atoi(strings.TrimSpace(mult[1]))
			return func(i int) int {
				return i * num
			}
		default:
			num, _ := strconv.Atoi(strings.TrimSpace(mult[0]))
			return func(i int) int {
				return i * num
			}
		}
	default:
		mult := strings.Split(strings.TrimSpace(op), "+")
		switch {
		case strings.TrimSpace(mult[0]) == "old" && strings.TrimSpace(mult[1]) == "old":
			return func(i int) int {
				return i + i
			}
		case strings.TrimSpace(mult[0]) == "old":
			num, _ := strconv.Atoi(strings.TrimSpace(mult[1]))
			return func(i int) int {
				return i + num
			}
		default:
			num, _ := strconv.Atoi(strings.TrimSpace(mult[0]))
			return func(i int) int {
				return i + num
			}
		}
	}
}

func getStartingItems(line string) []int {
	startingItemsStr := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ",")
	return utils.StrSliceToInt(startingItemsStr)
}

func processRound(m *monkey, divisor int) {
	for _, item := range m.StartingItems {
		//newValue := m.Operation(item) / 3 // first part
		newValue := (m.Operation(item) % divisor) // second part
		m.Test(newValue).StartingItems = append(m.Test(newValue).StartingItems, newValue)
		m.Count += 1
	}
	m.StartingItems = []int{}
}

func calcMonkeyBusiness(monkeys []*monkey) int {
	var count []int
	for _, m := range monkeys {
		count = append(count, m.Count)
	}
	sort.Ints(count)
	return count[len(count)-1] * count[len(count)-2]
}

func PuzzleEleven() map[string]int {
	//input := utils.ReadInput(utils.GetWD() + "/Day11/data/test.txt")
	input := utils.ReadInput(utils.GetWD() + "/Day11/data/input.txt")
	monkeys := buildMonkeys(input)

	divisors := monkeys[0].Divisor
	for idx := 1; idx < len(monkeys); idx++ {
		divisors *= monkeys[idx].Divisor
	}

	for range utils.GetRange(1, 10000) {
		for _, mnky := range monkeys {
			processRound(mnky, divisors)
		}
	}

	ans1 := calcMonkeyBusiness(monkeys)
	ans2 := 2
	return map[string]int{"ans1": ans1, "ans2": ans2}
}

func main() {
	log.Printf("PuzzleEleven(): %v", PuzzleEleven())
}
