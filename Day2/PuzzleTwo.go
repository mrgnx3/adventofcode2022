package main

import (
	"log"
	"os"
	"strings"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

type round struct {
	Opponent string
	You      string
}

func (r *round) handValue() int {
	switch r.You {
	case "X":
		return ROCK
	case "Y":
		return PAPER
	default:
		return SCISSORS
	}
}

func filter(roundString string) round {
	splitRound := strings.Split(roundString, " ")
	return round{
		Opponent: splitRound[0],
		You:      splitRound[1],
	}
}

func (r *round) yourScore() int {
	return r.getResult() + r.handValue()
}

func (r *round) getResult() int {
	switch {
	case r.Opponent == "A" && r.You == "X":
		return 3
	case r.Opponent == "B" && r.You == "X":
		return 0
	case r.Opponent == "C" && r.You == "X":
		return 6
	case r.Opponent == "A" && r.You == "Y":
		return 6
	case r.Opponent == "B" && r.You == "Y":
		return 3
	case r.Opponent == "C" && r.You == "Y":
		return 0
	case r.Opponent == "A" && r.You == "Z":
		return 0
	case r.Opponent == "B" && r.You == "Z":
		return 6
	default:
		return 3
	}
}

func (r *round) updateValues() {
	switch {
	case r.Opponent == "A" && r.You == "X":
		r.You = "Z"
	case r.Opponent == "B" && r.You == "X":
		r.You = "X"
	case r.Opponent == "C" && r.You == "X":
		r.You = "Y"
	case r.Opponent == "A" && r.You == "Y":
		r.You = "X"
	case r.Opponent == "B" && r.You == "Y":
		r.You = "Y"
	case r.Opponent == "C" && r.You == "Y":
		r.You = "Z"
	case r.Opponent == "A" && r.You == "Z":
		r.You = "Y"
	case r.Opponent == "B" && r.You == "Z":
		r.You = "Z"
	default:
		r.You = "X"
	}
}

func PuzzleTwo() map[string]int {
	file, err := os.ReadFile("data/input.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}

	roundsStr := strings.Split(string(file), "\n")

	scoreOne := roundTotal(roundsStr, false)
	scoreTwo := roundTotal(roundsStr, true)

	return map[string]int{"roundOne": scoreOne, "roundTwo": scoreTwo}
}

func roundTotal(roundsStr []string, secondPuzzle bool) int {
	score := 0
	for _, roundStr := range roundsStr {
		if roundStr != "" {
			res := filter(roundStr)
			if secondPuzzle {
				res.updateValues()
			}
			score += res.yourScore()
		}
	}
	return score
}

func main() {
	log.Printf("PuzzleTwo(): %v", PuzzleTwo())
}
