package day4_sections

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/geordie/adventofcode2021/util"
)

type Range struct {
	Start int
	End   int
}

func SolveDay4Puzzle1() {

	file, err := os.Open("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	iCount := 0
	for scanner.Scan() {
		sLine := scanner.Text()

		ranges := strings.Split(sLine, ",")

		range1 := Range{}
		range1.parseRange(ranges[0])

		range2 := Range{}
		range2.parseRange(ranges[1])

		if range1.contains(range2) || range2.contains(range1) {
			iCount++
		}
	}

	fmt.Println("DAY 4, PUZZLE 1 ANSWER: ", iCount)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func SolveDay4Puzzle2() {

	file, err := os.Open("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	iCount := 0
	for scanner.Scan() {
		sLine := scanner.Text()

		ranges := strings.Split(sLine, ",")

		range1 := Range{}
		range1.parseRange(ranges[0])

		range2 := Range{}
		range2.parseRange(ranges[1])

		if range1.overlaps(range2) {
			iCount++
		}
	}

	fmt.Println("DAY 4, PUZZLE 2 ANSWER: ", iCount)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *Range) parseRange(sRange string) {
	range1Bounds := strings.Split(sRange, "-")

	if len(range1Bounds) != 2 {
		return
	}

	r.Start = util.GetIntFromString(range1Bounds[0])
	r.End = util.GetIntFromString(range1Bounds[1])
}

func (r1 *Range) overlaps(r2 Range) bool {
	if r1.Start <= r2.Start && r1.End >= r2.Start {
		return true
	} else if r2.Start <= r1.Start && r2.End >= r1.Start {
		return true
	}
	return false
}

func (r1 *Range) contains(r2 Range) bool {
	if r1.Start <= r2.Start && r1.End >= r2.End {
		return true
	}
	return false
}
