package signals

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Buffer []rune

func SolveDay6Puzzle2() {

	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	buffer := make(Buffer, 14)
	i := 0

	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			// If no duplicates were found, we have the answer
			if !buffer.hasDups() {
				fmt.Println("Day 6, PUZZLE 2 ANSWER: ", i)
				return
			}
			// Shift the buffer
			buffer = append(buffer[1:14], c)
			i++
		}
	}
}

func SolveDay6Puzzle1() {

	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	buffer := make(Buffer, 4)
	i := 0

	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			// If no duplicates were found, we have the answer
			if !buffer.hasDups() {
				fmt.Println("Day 6, PUZZLE 1 ANSWER: ", i)
				return
			}
			// Shift the buffer
			buffer = append(buffer[1:4], c)
			i++
		}
	}
}

func (b Buffer) hasDups() bool {

	keys := make(map[rune]bool)
	for _, r := range b {
		if r == 0 {
			return true
		} else if keys[r] == true {
			return true
		} else {
			keys[r] = true
		}
	}
	return false
}
