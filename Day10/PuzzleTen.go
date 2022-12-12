package main

import (
	"2022/utils"
	"log"
	"strconv"
	"strings"
)

/*
Puzzle input
cpu program

clock circuit
tick = cycle

cpu reg = x = 1
instructions:
	addx V takes 2cycles then x += v
	noop takes 1 cycle

signal strength = cycle (idx) * x
find sum of 20th, 60th, 100th, 140th, 180th, and 220th signal strengths


*/

var SignalStengthTargets = []int{19, 59, 99, 139, 179, 219}

func draw(futures map[int]int) {
	utils.Section()
	CRT, divide := 0, 0
	for idx, _ := range utils.GetRange(0, len(futures)) {
		if (idx)%40 == 0 {
			log.Print()
			CRT = 0
			if divide%6 == 0 {
				utils.Section()
			}
			divide += 1
		}
		sprite(CRT, futures[idx])
		CRT += 1
	}
}

func sprite(idx, value int) {
	if idx-1 <= value && idx+1 >= value {
		print("#")
	} else {
		print(".")
	}
}

func runProgram(commands []string) int {
	x, v := 1, 0
	cycle := 0
	futures := map[int]int{}
	for _, command := range commands {
		if command == "noop" {
			cycle += 1
			futures[cycle] = x
		} else {
			cycle += 1
			futures[cycle] = x
			v, _ = strconv.Atoi(strings.Split(command, " ")[1])
			x += v
			cycle += 1
			futures[cycle] = x
		}
	}

	total := 0
	for _, i := range SignalStengthTargets {
		total += (i + 1) * futures[i]
	}
	draw(futures)
	return total
}

func PuzzleTen() map[string]int {
	//input := utils.ReadInput(utils.GetWD() + "/Day10/data/test.txt")
	input := utils.ReadInput(utils.GetWD() + "/Day10/data/input.txt")
	ans1 := runProgram(input)
	ans2 := 2
	return map[string]int{"ans1": ans1, "ans2": ans2}
}

func main() {
	log.Printf("PuzzleTen(): %v", PuzzleTen())
}
