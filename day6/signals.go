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
	solvePuzzle(14)
}

func SolveDay6Puzzle1() {
	solvePuzzle(4)
}

func solvePuzzle(bufferSize int) {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	buffer := make(Buffer, bufferSize)
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
			buffer = append(buffer[1:bufferSize], c)
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
