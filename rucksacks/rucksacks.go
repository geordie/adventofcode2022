package rucksacks

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type Comparment []byte

type Rucksack struct {
	Comparment1 Comparment
	Comparment2 Comparment
}

func SolveDay3Puzzle2() {

	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	iSum := 0
	for scanner.Scan() {

		batch := make([]string, 3)
		batch[0] = scanner.Text()

		for i := 1; i < 3; i++ {
			scanner.Scan()
			batch[i] = scanner.Text()
		}

		iCommonItem := findCommonBatchElement(batch)

		iSum += iCommonItem
	}

	fmt.Println("DAY 3, PUZZLE 2 ANSWER: ", iSum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func SolveDay3Puzzle1() {

	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	iSum := 0
	for scanner.Scan() {
		s := scanner.Text()

		iRucksackSize := len(s)
		iCompartmentSize := iRucksackSize / 2

		s1 := s[0:iCompartmentSize]
		s2 := s[iCompartmentSize:iRucksackSize]

		rucksack := Rucksack{Comparment1: []byte(s1), Comparment2: []byte(s2)}
		iCommonItem := rucksack.findCommonItemBetweenCompartments()

		iSum += iCommonItem

	}

	fmt.Println("DAY 3, PUZZLE 1 ANSWER: ", iSum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findCommonBatchElement(batch []string) int {
	iElemVal := 0
	for _, ch := range batch[0] {
		if strings.Contains(batch[1], string(ch)) &&
			strings.Contains(batch[2], string(ch)) {

			iCharVal := int(ch)
			// uppercase
			if iCharVal <= 90 {
				return iCharVal - 38
			} else {
				// lowercase
				return iCharVal - 96
			}
		}
	}
	return iElemVal
}

func (r *Rucksack) findCommonItemBetweenCompartments() int {
	iElemVal := 0
	for _, elem := range r.Comparment1 {
		if bytes.Contains(r.Comparment2, []byte{elem}) {
			if elem <= 90 {
				iElemVal = int(elem) - 38
			} else {
				// lowercase
				iElemVal = int(elem) - 96
			}
			return iElemVal
		}
	}
	return 0
}
