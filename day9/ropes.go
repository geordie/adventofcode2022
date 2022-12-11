package ropes

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/geordie/adventofcode2021/util"
)

type Direction int

const (
	None Direction = iota
	Right
	Left
	Up
	Down
)

type RopeMove struct {
	dir Direction
	mag int
}

type Position struct {
	x int
	y int
}

func SolveDay9Puzzle1() {
	solvePuzzle1()
}

func solvePuzzle1() {
	file, err := os.Open("input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	posHistory := make([]Position, 0)
	posHistory = append(posHistory, Position{x: 0, y: 0})

	ropeMoves := make([]RopeMove, 0)

	scanner := bufio.NewScanner(file)

	i := 0
	posHead := Position{}
	posTail := Position{}

	// Build the sequence of rope movements from the input
	for scanner.Scan() {
		sLine := scanner.Text()
		ropeMove := RopeMove{}
		ropeMove.parse(sLine)
		ropeMoves = append(ropeMoves, ropeMove)
		i++
	}

	for _, rm := range ropeMoves {
		for i := 0; i < rm.mag; i++ {
			posHead.moveHeadOne(rm)
			posTail.followHead(posHead)

			pos := Position{x: posTail.x, y: posTail.y}
			if !contains(posHistory, pos) {
				posHistory = append(posHistory, pos)
			}
		}
	}

	fmt.Println("DAY 9, PUZZLE 1 ANSWER: ", len(posHistory))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (rm *RopeMove) parse(sLine string) {
	bDir := sLine[0]
	sMagnitude := sLine[2:]

	rm.parseDirection(bDir)
	rm.parseMagnitude(sMagnitude)
}

func (rm *RopeMove) parseDirection(bDir byte) {
	if bDir == 'R' {
		rm.dir = Right
	} else if bDir == 'L' {
		rm.dir = Left
	} else if bDir == 'U' {
		rm.dir = Up
	} else if bDir == 'D' {
		rm.dir = Down
	}
}

func (rm *RopeMove) parseMagnitude(sMagnitude string) {
	iMagnitude := util.GetIntFromString(sMagnitude)
	rm.mag = iMagnitude
}

func (p *Position) moveHeadOne(rm RopeMove) {
	if rm.dir == Right {
		p.x++
	} else if rm.dir == Left {
		p.x--
	} else if rm.dir == Up {
		p.y++
	} else if rm.dir == Down {
		p.y--
	}
}

func (posTail *Position) followHead(posHead Position) {
	if posTail.x == posHead.x {
		if posHead.y-posTail.y == 2 {
			posTail.y++
		} else if posTail.y-posHead.y == 2 {
			posTail.y--
		}
	} else if posTail.y == posHead.y {
		if posHead.x-posTail.x == 2 {
			posTail.x++
		} else if posTail.x-posHead.x == 2 {
			posTail.x--
		}
	} else if posHead.x-posTail.x == 2 {
		posTail.x++
		posTail.y = posHead.y
	} else if posTail.x-posHead.x == 2 {
		posTail.x--
		posTail.y = posHead.y
	} else if posHead.y-posTail.y == 2 {
		posTail.y++
		posTail.x = posHead.x
	} else if posTail.y-posHead.y == 2 {
		posTail.y--
		posTail.x = posHead.x
	}
}

func contains(positions []Position, p Position) bool {
	for _, pos := range positions {
		if pos == p {
			return true
		}
	}
	return false
}
