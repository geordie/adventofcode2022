package directories

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/geordie/adventofcode2022/util"
)

type Node struct {
	name      string
	size      int
	totalSize int
	children  []*Node
	parent    *Node
}

var fs *Node

func SolveDay7Puzzle1() {
	fs = new(Node)
	solvePuzzle1()
}

func SolveDay7Puzzle2() {
	fs = new(Node)
	solvePuzzle2()
}

func solvePuzzle2() {
	file, err := os.Open("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs.name = "/"
	nodeCur := fs
	scanner := bufio.NewScanner(file)

	i := 0

	// Build the file system from the input set of commands
	for scanner.Scan() {
		sLine := scanner.Text()
		nodeCur = nodeCur.parseCommand(sLine, scanner)
		i++
	}

	fs.buildDirectorySizes()

	iFileSystemSize := 70000000
	iRequiredFreeSpace := 30000000
	iCurrentSizeUsed := fs.totalSize
	iSizeNeeded := iCurrentSizeUsed - (iFileSystemSize - iRequiredFreeSpace)

	nodeToDelete := fs.findMinDeletableDirectory(iSizeNeeded)

	fmt.Println("DAY 7, PUZZLE 2 ANSWER: ", nodeToDelete.totalSize)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solvePuzzle1() {
	file, err := os.Open("input/day7_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs.name = "/"
	nodeCur := fs
	scanner := bufio.NewScanner(file)

	i := 0

	// Build the file system from the input set of commands
	for scanner.Scan() {
		sLine := scanner.Text()
		nodeCur = nodeCur.parseCommand(sLine, scanner)
		i++
	}

	fs.buildDirectorySizes()

	iResult := fs.sumDirsNoBiggerThan(100000)

	fmt.Println("DAY 7, PUZZLE 1 ANSWER: ", iResult)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (nodeCur *Node) parseCommand(sLine string, scanner *bufio.Scanner) *Node {
	// Handle ls commands
	if sLine[0:4] == "$ ls" {
		for scanner.Scan() {
			sNextLine := scanner.Text()
			if sNextLine[0:1] == "$" {
				nodeCur = nodeCur.parseCommand(sNextLine, scanner)
				return nodeCur
			} else if sNextLine[0:3] == "dir" {
				subDir := new(Node)
				subDir.name = sNextLine[4:]
				subDir.size = 0
				subDir.parent = nodeCur
				nodeCur.children = append(nodeCur.children, subDir)
			} else /* assume file */ {
				parts := strings.Split(sNextLine, " ")
				file := new(Node)
				file.size = util.GetIntFromString(parts[0])
				file.name = parts[1]
				file.parent = nodeCur
				nodeCur.children = append(nodeCur.children, file)
			}
		}
	} else if sLine[0:5] == "$ cd " { // Handle cd commands
		sNextDir := sLine[5:]
		if sNextDir == "/" {
			nodeCur = fs
		} else if sNextDir == ".." {
			if nodeCur.name != "/" {
				nodeCur = nodeCur.parent
			}
			return nodeCur
		} else {
			for _, subNode := range nodeCur.children {
				if subNode.name == sNextDir {
					return subNode
				}
			}
		}
	}
	return nodeCur
}

func (node *Node) Print(bRecursive bool) {
	sParent := ""
	if node.parent != nil {
		sParent = node.parent.name
	}
	fmt.Println("Name:", node.name, "Size:", node.size, "Parent:", sParent, "TotalSize: ", node.totalSize)
	if bRecursive {
		for _, childNode := range node.children {
			childNode.Print(bRecursive)
		}
	}
}

func (node *Node) buildDirectorySizes() {

	// Depth first
	for _, childNode := range node.children {
		childNode.buildDirectorySizes()
	}

	// Files have size > 0
	if node.size > 0 {
		node.totalSize = node.size
	}

	// Aggregate sizes up the tree but not above the root
	if node.name != "/" {
		node.parent.totalSize += node.totalSize
	}
}

func (node *Node) sumDirsNoBiggerThan(iMaxSize int) int {

	iResult := 0

	// Depth first
	for _, childNode := range node.children {
		iResult += childNode.sumDirsNoBiggerThan(iMaxSize)
	}

	if node.size == 0 && node.totalSize <= iMaxSize {
		return iResult + node.totalSize
	}

	return iResult
}

func (node *Node) findMinDeletableDirectory(iMinSize int) *Node {
	var nodeMinDeletable *Node

	// Depth first
	for _, childNode := range node.children {
		candidateNode := childNode.findMinDeletableDirectory(iMinSize)
		if nodeMinDeletable == nil ||
			(candidateNode != nil &&
				candidateNode.totalSize >= iMinSize &&
				candidateNode.totalSize < nodeMinDeletable.totalSize) {
			nodeMinDeletable = candidateNode
		}
	}

	if node.size > 0 {
		return nodeMinDeletable
	} else if node.size == 0 && node.totalSize >= iMinSize {
		if nodeMinDeletable == nil ||
			(node.totalSize < nodeMinDeletable.totalSize) {
			return node
		}
	}

	return nodeMinDeletable
}
