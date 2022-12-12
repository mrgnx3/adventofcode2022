package main

import (
	"github.com/deckarep/golang-set"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
puzzle input:
paired section assignments

sections have unique ids
elves assigned to a range of sections

read file
split pairs
filter ranges
find slices entirely within the other slice
print count

*/

type assignment struct {
	Assignment string
}

type elfPair struct {
	ElfOne assignment
	ElfTwo assignment
}

func split(pair string) elfPair {
	elves := strings.Split(pair, ",")
	return elfPair{
		ElfOne: assignment{elves[0]},
		ElfTwo: assignment{elves[1]},
	}
}

func (a assignment) getRange() []int {
	var fullRange []int
	splitVal := strings.Split(a.Assignment, "-")
	lowerIndex, _ := strconv.Atoi(splitVal[0])
	higherIndex, _ := strconv.Atoi(splitVal[1])
	for i := lowerIndex; i <= higherIndex; i++ {
		fullRange = append(fullRange, i)
	}
	return fullRange
}

func sliceToSet(mySlice []int) mapset.Set {
	mySet := mapset.NewSet()
	for _, ele := range mySlice {
		mySet.Add(ele)
	}
	return mySet
}

func findOverlap(pair elfPair) []int {
	setElfOne := sliceToSet(pair.ElfOne.getRange())
	setElfTwo := sliceToSet(pair.ElfTwo.getRange())

	if setElfTwo.IsSubset(setElfOne) {
		return pair.ElfOne.getRange()
	} else if setElfOne.IsSubset(setElfTwo) {
		return pair.ElfTwo.getRange()
	}
	return nil
}

func findAnyOverlap(pair elfPair) int {
	setElfOne := pair.ElfOne.getRange()
	setElfTwo := pair.ElfTwo.getRange()

	for _, overlapOne := range setElfOne {
		for _, overlapTwo := range setElfTwo {
			if overlapOne == overlapTwo {
				return 1
			}
		}
	}
	return 0
}

func PuzzleThree() map[string]int {
	//file, err := os.ReadFile("data/test.txt")
	file, err := os.ReadFile("data/input.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}
	var listOfPairs []elfPair
	for _, pair := range strings.Split(string(file), "\n") {
		if pair != "" {
			listOfPairs = append(listOfPairs, split(pair))
		}
	}

	var overlap []interface{}
	var anyOverlap []int
	for _, pair := range listOfPairs {
		overlapPair := findOverlap(pair)
		if overlapPair != nil {
			overlap = append(overlap, overlapPair)
		}
		anyOverlapPair := findAnyOverlap(pair)
		if anyOverlapPair != 0 {
			anyOverlap = append(anyOverlap, anyOverlapPair)
		}
	}

	//roundOne := count(overlap)

	//sumOfItems1 := calculateTotal(packs, false)
	//sumOfItems2 := calculateTotal(packs, true)

	return map[string]int{"roundOne": len(overlap), "roundTwo": len(anyOverlap)}
}

func main() {
	log.Printf("PuzzleThree(): %v", PuzzleThree())
}
