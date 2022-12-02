package main

import (
	"fmt"
	"sort"

	calories "github.com/geordie/adventofcode2022/calories"
)

func main() {
	day1puzzle1()
	day1puzzle2()
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
