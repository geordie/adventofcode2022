package main

import (
	"fmt"
	"sort"

	calories "github.com/geordie/adventofcode2022/day1"
	rps "github.com/geordie/adventofcode2022/day2"
	rucksacks "github.com/geordie/adventofcode2022/day3"
	sections "github.com/geordie/adventofcode2022/day4"
	crates "github.com/geordie/adventofcode2022/day5"
	signals "github.com/geordie/adventofcode2022/day6"
)

func main() {
	day6puzzle1()
	day6puzzle2()
}

func day6puzzle2() {
	signals.SolveDay6Puzzle2()
}

func day6puzzle1() {
	signals.SolveDay6Puzzle1()
}

func day5puzzle2() {
	crates.SolveDay5Puzzle2()
}

func day5puzzle1() {
	crates.SolveDay5Puzzle1()
}

func day4puzzle2() {
	sections.SolveDay4Puzzle2()
}

func day4puzzle1() {
	sections.SolveDay4Puzzle1()
}

func day3puzzle2() {
	rucksacks.SolveDay3Puzzle2()
}

func day3puzzle1() {
	rucksacks.SolveDay3Puzzle1()
}

func day2puzzle2() {
	rpsRounds := rps.ParseDay2Puzzle2Input()
	iScore := rpsRounds.CalculateScore()
	fmt.Println("DAY 2, PUZZLE 2 ANSWER: ", iScore)
}

func day2puzzle1() {
	rpsRounds := rps.ParseDay2Puzzle1Input()
	iScore := rpsRounds.CalculateScore()
	fmt.Println("DAY 2, PUZZLE 1 ANSWER: ", iScore)
}

func day1puzzle2() {
	elfCalories := calories.ParseDay1Input()

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	top3Total := elfCalories[0] + elfCalories[1] + elfCalories[2]
	fmt.Println("DAY 1, PUZZLE 2 ANSWER: ", top3Total)
}

func day1puzzle1() {
	elfCalories := calories.ParseDay1Input()

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	fmt.Println("DAY 1, PUZZLE 1 ANSWER: ", elfCalories[0])
}
