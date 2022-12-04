package sections

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/geordie/adventofcode2021/util"
)

type Range struct {
	rangeStart int
	rangeEnd   int
}

type RangePair struct {
	range1 Range
	range2 Range
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

		sRange1 := ranges[0]
		range1Bounds := strings.Split(sRange1, "-")
		iRange1Start := util.GetIntFromString(range1Bounds[0])
		iRange1End := util.GetIntFromString(range1Bounds[1])

		sRange2 := ranges[1]
		range2Bounds := strings.Split(sRange2, "-")
		iRange2Start := util.GetIntFromString(range2Bounds[0])
		iRange2End := util.GetIntFromString(range2Bounds[1])

		if iRange1Start <= iRange2Start && iRange1End >= iRange2End {
			iCount++
		} else if iRange2Start <= iRange1Start && iRange2End >= iRange1End {
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

		// fmt.Println("LINE:", sLine)

		ranges := strings.Split(sLine, ",")

		sRange1 := ranges[0]
		range1Bounds := strings.Split(sRange1, "-")
		iRange1Start := util.GetIntFromString(range1Bounds[0])
		iRange1End := util.GetIntFromString(range1Bounds[1])
		// fmt.Println("Range1: ", iRange1Start, iRange1End)

		sRange2 := ranges[1]
		range2Bounds := strings.Split(sRange2, "-")
		iRange2Start := util.GetIntFromString(range2Bounds[0])
		iRange2End := util.GetIntFromString(range2Bounds[1])
		// fmt.Println("Range2: ", iRange2Start, iRange2End)

		if iRange1Start <= iRange2Start && iRange1End >= iRange2Start {
			iCount++
			// fmt.Println("OVERLAP 1")
		} else if iRange2Start <= iRange1Start && iRange2End >= iRange1Start {
			iCount++
			// fmt.Println("OVERLAP 2")
		}
	}

	fmt.Println("DAY 4, PUZZLE 2 ANSWER: ", iCount)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}