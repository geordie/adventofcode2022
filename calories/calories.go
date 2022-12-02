package calories

import (
	"bufio"
	"log"
	"os"

	util "github.com/geordie/adventofcode2021/util"
)

type ElfCalories []int

func ParseDay1Input() ElfCalories {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	elfCalories := make([]int, 10)
	iCalories := 0

	for scanner.Scan() {
		s := scanner.Text()

		if len(s) == 0 {
			elfCalories = append(elfCalories, iCalories)
			iCalories = 0
		} else {
			iCalories = iCalories + util.GetIntFromString(s)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elfCalories
}
