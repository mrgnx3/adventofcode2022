package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
puzzle input:
datastream buffer

detect start of packet marker in datastream
start of packet = 4 characters all different

read file
find start of packet
return number of characters from start of line to last character in start of packet




*/

type packet struct {
	Pos1 int
}

func findStartingPacketLocation(file string, depth int) packet {
	startingpacket := packet{
		Pos1: -1,
	}
	for i := 0; i < len(file); i++ {
		fileList := file[i : i+depth]
		if assertNoMatches(fileList, i, depth) {
			startingpacket.Pos1 = i
			break
		}
	}
	return startingpacket
}

func assertNoMatches(file string, i int, depth int) bool {
	/*
		for i in list length
		get value at i
		remove i from list
		if i in list without i return true
	*/
	noMatches := true
	for idx := 1; idx < len(file); idx++ {
		if strings.Contains(file[idx:], string(file[idx-1])) {
			noMatches = false
		}
		if !noMatches {
			break
		}
	}
	return noMatches
}

func PuzzleSix() map[string]int {
	file, err := os.Open("data/input.txt")
	//file, err := os.Open("data/test.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}
	fileScanner := bufio.NewScanner(file)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	for fileScanner.Scan() {
		startingPacketLocation1 := findStartingPacketLocation(fileScanner.Text(), 4)
		startingPacketLocation2 := findStartingPacketLocation(fileScanner.Text(), 14)
		return map[string]int{"roundOne": startingPacketLocation1.Pos1 + 4, "roundTwo": startingPacketLocation2.Pos1 + 14}
	}
	return nil
}

func main() {
	log.Printf("PuzzleSix(): %v", PuzzleSix())
}
