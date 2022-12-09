package forest

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/geordie/adventofcode2022/util"
)

type Forest [][]int

func SolveDay7Puzzle1() {
	solvePuzzle1()
}

func solvePuzzle1() {
	file, err := os.Open("input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0
	var forest Forest

	// Build the forest from the input
	for scanner.Scan() {
		sLine := scanner.Text()
		if i == 0 {
			forest = make(Forest, len(sLine))
			for i := 0; i < len(sLine); i++ {
				forest[i] = make([]int, len(sLine))
			}
		}
		forest.parseLine(sLine, i)
		i++
	}

	// Find visible trees in the forest
	iVisibleTrees := forest.findVisible()

	fmt.Println("DAY 7, PUZZLE 1 ANSWER: ", iVisibleTrees)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (f Forest) parseLine(sLine string, row int) {
	for idx, item := range sLine {
		f[row][idx] = util.GetIntFromString(string(item))
	}
}

func (f Forest) findVisible() int {

	iVisibleTrees := 0

	width := len(f)
	// Initialize an slice of slices to track visible trees
	visibleTrees := make(Forest, width)
	for w := 0; w < width; w++ {
		visibleTrees[w] = make([]int, width)
	}

	// Set the trees as the four corners of the forest to visible
	visibleTrees[0][0] = 1
	visibleTrees[0][width-1] = 1
	visibleTrees[width-1][0] = 1
	visibleTrees[width-1][width-1] = 1

	// Look for visible trees in columns
	for j := 1; j <= width-2; j++ {
		iTallestHeight := -1
		// Look from top down
		for i := 0; i <= width-1; i++ {
			iCurTreeHeight := f[i][j]
			if iCurTreeHeight > iTallestHeight {
				visibleTrees[i][j] = 1
				iTallestHeight = iCurTreeHeight
			}
		}
		iTallestHeight = -1
		// Look from bottom up
		for i := width - 1; i >= 0; i-- {
			iCurTreeHeight := f[i][j]
			if iCurTreeHeight > iTallestHeight {
				visibleTrees[i][j] = 1
				iTallestHeight = iCurTreeHeight
			}
		}
	}

	// Look for visible trees in rows
	for i := 1; i <= width-2; i++ {
		iTallestHeight := -1
		// Look from left to right
		for j := 0; j <= width-1; j++ {
			iCurTreeHeight := f[i][j]
			if iCurTreeHeight > iTallestHeight {
				visibleTrees[i][j] = 1
				iTallestHeight = iCurTreeHeight
			}
		}
		iTallestHeight = -1
		// Look from right to left
		for j := width - 1; j >= 0; j-- {
			iCurTreeHeight := f[i][j]
			if iCurTreeHeight > iTallestHeight {
				visibleTrees[i][j] = 1
				iTallestHeight = iCurTreeHeight
			}
		}
	}

	// Count visible trees
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			iVisibleTrees += visibleTrees[i][j]
		}
	}
	return iVisibleTrees
}
