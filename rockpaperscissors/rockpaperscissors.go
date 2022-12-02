package rockpaperscissors

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type RPSMove int

const (
	None RPSMove = iota
	Rock
	Paper
	Scissors
)

type RPSOutcome int

const (
	Lose RPSOutcome = iota
	Draw
	Win
)

type RPSRound struct {
	Opponent RPSMove
	Self     RPSMove
	Outcome  RPSOutcome
}

type RPSRounds []RPSRound

func (r *RPSRound) parsePuzzle1(s string) {
	sInputs := strings.Split(s, " ")
	sOpponent := sInputs[0]
	sSelf := sInputs[1]

	if sOpponent == "A" {
		r.Opponent = Rock
	} else if sOpponent == "B" {
		r.Opponent = Paper
	} else if sOpponent == "C" {
		r.Opponent = Scissors
	}

	if sSelf == "X" {
		r.Self = Rock
	} else if sSelf == "Y" {
		r.Self = Paper
	} else if sSelf == "Z" {
		r.Self = Scissors
	}

	if (r.Self == Rock && r.Opponent == Scissors) ||
		(r.Self == Paper && r.Opponent == Rock) ||
		(r.Self == Scissors && r.Opponent == Paper) {
		r.Outcome = Win
	} else if r.Self == r.Opponent {
		r.Outcome = Draw
	} else {
		r.Outcome = Lose
	}
}

func (r *RPSRound) parsePuzzle2(s string) {
	sInputs := strings.Split(s, " ")
	sOpponent := sInputs[0]
	sOutcome := sInputs[1]

	if sOpponent == "A" {
		r.Opponent = Rock
	} else if sOpponent == "B" {
		r.Opponent = Paper
	} else if sOpponent == "C" {
		r.Opponent = Scissors
	}

	if sOutcome == "X" {
		r.Outcome = Lose
	} else if sOutcome == "Y" {
		r.Outcome = Draw
	} else if sOutcome == "Z" {
		r.Outcome = Win
	}

	if (r.Opponent == Paper && r.Outcome == Lose) ||
		(r.Opponent == Rock && r.Outcome == Draw) ||
		(r.Opponent == Scissors && r.Outcome == Win) {
		r.Self = Rock
	} else if (r.Opponent == Scissors && r.Outcome == Lose) ||
		(r.Opponent == Paper && r.Outcome == Draw) ||
		(r.Opponent == Rock && r.Outcome == Win) {
		r.Self = Paper
	} else {
		r.Self = Scissors
	}
}

func ParseDay2Puzzle1Input() RPSRounds {

	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	rpsRounds := make([]RPSRound, 0)

	for scanner.Scan() {
		s := scanner.Text()

		rpsRound := new(RPSRound)
		rpsRound.parsePuzzle1(s)

		rpsRounds = append(rpsRounds, *rpsRound)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rpsRounds
}

func ParseDay2Puzzle2Input() RPSRounds {

	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	rpsRounds := make([]RPSRound, 0)

	for scanner.Scan() {
		s := scanner.Text()

		rpsRound := new(RPSRound)
		rpsRound.parsePuzzle2(s)

		rpsRounds = append(rpsRounds, *rpsRound)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rpsRounds
}

func (r *RPSRounds) CalculateScore() int {
	iTotalScore := 0
	for _, elem := range *r {
		iRoundScore := elem.calculateRoundScore()
		iTotalScore += iRoundScore
	}
	return iTotalScore
}

func (r *RPSRound) calculateRoundScore() int {

	iScoreChoice := int(r.Self)
	iScoreOutcome := int(r.Outcome) * 3

	iScore := iScoreChoice + iScoreOutcome

	return iScore
}
