package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func PuzzleOne() int {
	file, err := os.ReadFile("One/input.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}
	total := strings.Split(string(file), "\n")
	elves := splitElves(total)
	totals := calculate(elves)

	numberOne := compare(totals)
	delete(elves, numberOne["elf"]-1)

	totals = calculate(elves)
	numberTwo := compare(totals)
	delete(elves, numberTwo["elf"]-1)

	totals = calculate(elves)
	numberThree := compare(totals)
	topThree := numberOne["cals"] + numberTwo["cals"] + numberThree["cals"]
	return topThree
}

/*
compare top 3
two options
1. Find highest, then find second highest, then find third highest
2. Check each time if it compares to the top 3
*/

func compare(totals map[int]int) map[string]int {
	currentHighest := 0
	highest := map[string]int{"elf": 0, "cals": 0}

	for elf, totalCals := range totals {
		if totals[elf] > currentHighest {
			highest["elf"] = elf + 1
			highest["cals"] = totalCals
			currentHighest = totalCals
		}
	}
	return highest
}

func calculate(elves map[int][]int) map[int]int {
	totals := make(map[int]int, 0)
	for elf, cals := range elves {
		totals[elf] = sum(cals)
	}
	return totals
}

func sum(elves []int) int {
	totalCals := 0
	for _, cals := range elves {
		totalCals += cals
	}
	return totalCals
}

func splitElves(total []string) map[int][]int {
	elves := make(map[int][]int, 0)
	index := 0
	for _, calories := range total {
		caloriesInt, err := strconv.Atoi(calories)
		if err != nil {
			index += 1
		} else {
			elves[index] = append(elves[index], caloriesInt)
		}
	}
	return elves
}

func main() {
	print(PuzzleOne())
}
