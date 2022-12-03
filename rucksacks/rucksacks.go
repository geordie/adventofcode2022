package rucksacks

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type Comparment []byte

type Rucksack struct {
	Comparment1 Comparment
	Comparment2 Comparment
}

type Rucksacks []Rucksack

func ParseDay3Puzzle1Input() Rucksacks {

	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	rucksacks := make([]Rucksack, 0)

	iSum := 0
	for scanner.Scan() {
		s := scanner.Text()

		iRucksackSize := len(s)
		iCompartmentSize := iRucksackSize / 2

		s1 := s[0:iCompartmentSize]
		s2 := s[iCompartmentSize:iRucksackSize]

		rucksack := Rucksack{Comparment1: []byte(s1), Comparment2: []byte(s2)}
		iCommonItem := rucksack.FindCommonItem()

		iSum += iCommonItem

		rucksacks = append(rucksacks, rucksack)
	}

	fmt.Println("DAY 3, PUZZLE 1 ANSWER: ", iSum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rucksacks
}

func (r *Rucksack) FindCommonItem() int {
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
