package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

/*
puzzle input:
each item in each rucksack

rucksacks for the journey
Two compartments
One type per compartment
One item mixed up per rucksack
same amount of items in both packs
find common char
prioritise
sum

read file
split packs by compartment
find common item
prioritise
sum up

*/

type pack struct {
	CompartmentOne string
	CompartmentTwo string
	WholePack      string
}

type groupedPack struct {
	GroupOne   pack
	GroupTwo   pack
	GroupThree pack
}

func (p groupedPack) commonBadge() rune {
	for _, singleChar := range p.GroupOne.WholePack {
		if strings.Contains(p.GroupTwo.WholePack, string(singleChar)) && strings.Contains(p.GroupThree.WholePack, string(singleChar)) {
			return singleChar
		}
	}
	panic("no common badge")
}

func (p pack) commonValue() rune {
	for _, singleChar := range p.CompartmentOne {
		if strings.Contains(p.CompartmentTwo, string(singleChar)) {
			return singleChar
		}
	}
	panic("no common value found")
}

func (p pack) commonBadge() rune {
	for _, singleChar := range p.CompartmentOne {
		if strings.Contains(p.CompartmentTwo, string(singleChar)) {
			return singleChar
		}
	}
	panic("no common value found")
}

func PuzzleThree() map[string]int {
	//file, err := os.ReadFile("data/test.txt")
	file, err := os.ReadFile("data/input.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}

	packs := strings.Split(string(file), "\n")

	sumOfItems1 := calculateTotal(packs, false)
	sumOfItems2 := calculateTotal(packs, true)

	return map[string]int{"roundOne": sumOfItems1, "roundTwo": sumOfItems2}
}

func calculateTotal(packs []string, secondPuzzle bool) int {
	var packList []pack
	for _, packStr := range packs {
		if packStr != "" {
			strLen := len(packStr)
			packList = append(packList, pack{
				CompartmentOne: packStr[0:(strLen / 2)],
				CompartmentTwo: packStr[strLen/2 : strLen],
				WholePack:      packStr,
			})
		}
	}

	var commonValue []rune
	if secondPuzzle {
		groupedPackList := groupPack(packList)
		for _, groupPack := range groupedPackList {
			commonValue = append(commonValue, groupPack.commonBadge())
		}
	} else {
		for _, packItem := range packList {
			commonValue = append(commonValue, packItem.commonValue())
		}
	}
	return prioritiseItems(commonValue)
}

func groupPack(packList []pack) []groupedPack {
	index := 0
	var groupPack []groupedPack
	for idx, _ := range packList {
		if (idx+1)%3 == 0 {
			groupPack = append(groupPack, groupedPack{
				GroupOne:   packList[idx-2],
				GroupTwo:   packList[idx-1],
				GroupThree: packList[idx],
			})
			index += 1
		}
	}
	return groupPack
}

func prioritiseItems(commonItems []rune) int {
	total := 0
	for _, commonItem := range commonItems {
		total += convert(commonItem)
	}
	return total
}

func convert(in rune) int {
	if unicode.IsUpper(in) {
		return int(in - 'A' + 27)
	} else {
		return int(in - 'a' + 1)
	}
}

func main() {
	log.Printf("PuzzleThree(): %v", PuzzleThree())
}
