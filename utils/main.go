package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetWD() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

func GetRange(lowerIndex, higherIndex int64) []int64 {
	var fullRange []int64
	for i := lowerIndex; i <= higherIndex; i++ {
		fullRange = append(fullRange, i)
	}
	return fullRange
}

func ReadInput(file string) []string {
	datafile, err := os.Open(file)
	//datafile, err := os.Open("data/test.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}
	fileScanner := bufio.NewScanner(datafile)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(datafile)

	var input []string

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	return input
}

func Contains(int64s []int64, x int64) bool {
	for _, i := range int64s {
		if x == i {
			return true
		}
	}
	return false
}

func Sum(int64s []int64) int64 {
	var total int64
	total = 0
	for _, i := range int64s {
		total += i
	}
	return total
}

func Section() {
	log.Print("********************************")
}

func StrSliceToInt(strList []string) []int {
	var i []int
	for _, s := range strList {
		sConv, _ := strconv.Atoi(strings.TrimSpace(s))
		i = append(i, sConv)
	}
	return i
}
