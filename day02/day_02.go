package day02

import (
	"advent-of-code-2022/utils"
	"bufio"
	"fmt"
	"strings"
)

func Day02() {
	part1()
	part2()
}

func parseLine(line string) (string, string) {
	return string(line[0]), string(line[2])
}

func getShapeScore(shape string) int {
	switch shape {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	default:
		return -1
	}
}

func getOutcomeScore(player1 string, player2 string) int {
	// loss
	if (player1 == "A" && player2 == "C") ||
		(player1 == "B" && player2 == "A") ||
		(player1 == "C" && player2 == "B") {
		return 0
	}

	// draw
	if player1 == player2 {
		return 3
	}

	// win
	return 6
}

func part1() {
	score := 0

	scanner := bufio.NewScanner(strings.NewReader(utils.ReadInputFile("day02/input.txt")))
	for scanner.Scan() {
		line := scanner.Text()

		player1, player2 := parseLine(line)

		switch player2 {
		case "X":
			player2 = "A"
			break
		case "Y":
			player2 = "B"
			break
		case "Z":
			player2 = "C"
			break
		}

		shapeScore := getShapeScore(player2)
		outcomeScore := getOutcomeScore(player1, player2)

		score += shapeScore + outcomeScore
	}

	fmt.Println("Part 1: ", score)
}

func part2() {
	score := 0

	scanner := bufio.NewScanner(strings.NewReader(utils.ReadInputFile("day02/input.txt")))
	for scanner.Scan() {
		line := scanner.Text()

		player1, outcome := parseLine(line)

		player2 := getPlayer2(player1, outcome)

		shapeScore := getShapeScore(player2)
		outcomeScore := getOutcomeScore(player1, player2)

		score += shapeScore + outcomeScore
	}

	fmt.Println("Part 2: ", score)
}

func getPlayer2(player1 string, outcome string) string {

	nested := map[string]map[string]string{
		"A": {
			"X": "C",
			"Y": "A",
			"Z": "B",
		},
		"B": {
			"X": "A",
			"Y": "B",
			"Z": "C",
		},
		"C": {
			"X": "B",
			"Y": "C",
			"Z": "A",
		},
	}

	return nested[player1][outcome]
}
