package crates

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/geordie/adventofcode2021/util"
)

type Move struct {
	Amount int
	From   int
	To     int
}

type Moves []Move

type Stacks [][]string

func SolveDay5Puzzle2() {

	file, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	containerRows := make([]string, 10)
	stacks := make(Stacks, 12)
	scanner := bufio.NewScanner(file)

	i := 0

	// Get containers
	for scanner.Scan() {
		sLine := scanner.Text()

		if sLine[1] == '1' {
			break
		}
		containerRows[i] = sLine
		i++
	}

	// Store containers in stacks
	for _, row := range containerRows {
		iLen := len(row)
		for i = 3; i <= iLen; i = i + 4 {
			crate := row[i-2 : i-1]
			if len(strings.Trim(crate, " ")) == 0 {
				continue
			}
			iStack := ((i + 1) / 4) - 1
			stacks[iStack] = append(stacks[iStack], crate)
		}
	}

	//fmt.Println("STACKS BEFORE: ", stacks)

	// Burn the blank line
	scanner.Scan()

	// Make the container moves
	for scanner.Scan() {
		sMove := scanner.Text()
		move := Move{}
		move.parseMove(sMove)
		stacks.moveContainers2(move)
	}

	//fmt.Println("STACKS AFTER: ", stacks)

	sTopCrates := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			sTopCrates += stack[0]
		}
	}

	fmt.Println("DAY 5, PUZZLE 2 ANSWER: ", sTopCrates)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func SolveDay5Puzzle1() {

	file, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	containerRows := make([]string, 10)
	stacks := make(Stacks, 12)
	scanner := bufio.NewScanner(file)

	i := 0

	// Get containers
	for scanner.Scan() {
		sLine := scanner.Text()

		if sLine[1] == '1' {
			break
		}
		containerRows[i] = sLine
		i++
	}

	// Store containers in stacks
	for _, row := range containerRows {
		iLen := len(row)
		for i = 3; i <= iLen; i = i + 4 {
			crate := row[i-2 : i-1]
			if len(strings.Trim(crate, " ")) == 0 {
				continue
			}
			iStack := ((i + 1) / 4) - 1
			stacks[iStack] = append(stacks[iStack], crate)
		}
	}

	// fmt.Println("STACKS BEFORE: ", stacks)

	// Burn the blank line
	scanner.Scan()

	// Make the container moves
	for scanner.Scan() {
		sMove := scanner.Text()
		move := Move{}
		move.parseMove(sMove)
		stacks.moveContainers(move)
	}

	// fmt.Println("STACKS AFTER: ", stacks)

	sTopCrates := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			sTopCrates += stack[0]
		}
	}

	fmt.Println("DAY 5, PUZZLE 1 ANSWER: ", sTopCrates)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m *Move) parseMove(sMove string) {
	moveParts := strings.Split(sMove, " ")

	if len(moveParts) != 6 {
		return
	}

	m.Amount = util.GetIntFromString(moveParts[1])
	m.From = util.GetIntFromString(moveParts[3]) - 1
	m.To = util.GetIntFromString(moveParts[5]) - 1
}

func (s Stacks) moveContainers(m Move) {
	// Get the containers to move
	toMove := s[m.From][0:m.Amount]

	for _, container := range toMove {
		s[m.To] = append([]string{container}, s[m.To]...)
	}

	s[m.From] = s[m.From][m.Amount:]
}

func (s Stacks) moveContainers2(m Move) {
	// Get the containers to move
	toMove := s[m.From][0:m.Amount]

	for i, j := 0, len(toMove)-1; i < j; i, j = i+1, j-1 {
		toMove[i], toMove[j] = toMove[j], toMove[i]
	}

	for _, container := range toMove {
		s[m.To] = append([]string{container}, s[m.To]...)
	}

	s[m.From] = s[m.From][m.Amount:]
}
