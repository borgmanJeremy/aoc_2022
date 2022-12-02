package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func shapeScore(input string) int {
	switch input {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	log.Fatal("Improper character passed to shape score")
	return 0
}

func scoreOutcome(opp string, you string) int {
	//Draws
	if opp == "A" && you == "X" {
		return 3
	} else if opp == "B" && you == "Y" {
		return 3
	} else if opp == "C" && you == "Z" {
		return 3
		//Wins
	} else if opp == "A" && you == "Y" {
		return 6
	} else if opp == "B" && you == "Z" {
		return 6
	} else if opp == "C" && you == "X" {
		return 6
	}
	return 0
}

func calculateShape(opp string, outcome string) string {
	// Draw
	if outcome == "Y" {
		if opp == "A" {
			return "X"
		}
		if opp == "B" {
			return "Y"
		}
		if opp == "C" {
			return "Z"
		}
	}

	// Lose
	if outcome == "X" {
		if opp == "A" {
			return "Z"
		}
		if opp == "B" {
			return "X"
		}
		if opp == "C" {
			return "Y"
		}
	}

	// Win
	if outcome == "Z" {
		if opp == "A" {
			return "Y"
		}
		if opp == "B" {
			return "Z"
		}
		if opp == "C" {
			return "X"
		}
	}
	log.Fatal("Invalid combo")
	return ""
}

func main() {
	f, err := os.Open("input/input.txt")
	checkError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	score_1 := 0
	score_2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Split(line, " ")
		opp := round[0]
		you := round[1]

		score_1 += shapeScore(you)
		score_1 += scoreOutcome(opp, you)

		shapeToPlay := calculateShape(opp, you)
		score_2 += shapeScore(shapeToPlay)
		score_2 += scoreOutcome(opp, shapeToPlay)
	}
	fmt.Println("Part 1: ", score_1)
	fmt.Println("Part 2: ", score_2)
}
