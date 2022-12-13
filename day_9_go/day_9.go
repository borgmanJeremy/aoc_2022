package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	x int
	y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func decodeDirection(dir string) Direction {
	switch dir {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right

	}
	log.Fatal("Invalid direction")
	return 0
}

func moveHead(headPos Point, dir Direction) Point {
	switch dir {
	case Up:
		headPos.y += 1
	case Down:
		headPos.y -= 1
	case Left:
		headPos.x -= 1
	case Right:
		headPos.x += 1
	}
	return headPos
}

func inColumn(headPos, tailPos Point) bool {
	return headPos.x == tailPos.x
}

func inRow(headPos, tailPos Point) bool {
	return headPos.y == tailPos.y
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveTail(headPos, tailPos Point) Point {
	// Check if overlapping
	if headPos == tailPos {
		return tailPos
	}

	if intAbs(headPos.x-tailPos.x) <= 1 && intAbs(headPos.y-tailPos.y) <= 1 {
		return tailPos
	}

	// Handle vertical and horizontal cases
	if inColumn(headPos, tailPos) {
		if headPos.y > (tailPos.y + 1) {
			tailPos.y = headPos.y - 1
		} else if headPos.y < (tailPos.y - 1) {
			tailPos.y = headPos.y + 1
		}
		return tailPos
	} else if inRow(headPos, tailPos) {
		if headPos.x > (tailPos.x + 1) {
			tailPos.x = headPos.x - 1
		} else if headPos.x < (tailPos.x - 1) {
			tailPos.x = headPos.x + 1
		}
		return tailPos
	}

	// Handle diagonal cases
	if headPos.x > tailPos.x && headPos.y > tailPos.y {
		tailPos.x += 1
		tailPos.y += 1
		return tailPos
	} else if headPos.x > tailPos.x && headPos.y < tailPos.y {
		tailPos.x += 1
		tailPos.y -= 1
		return tailPos
	} else if headPos.x < tailPos.x && headPos.y > tailPos.y {
		tailPos.x -= 1
		tailPos.y += 1
		return tailPos
	} else {
		tailPos.x -= 1
		tailPos.y -= 1
		return tailPos
	}
}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	headPos := Point{0, 0}
	knot1Pos := Point{0, 0}
	knot2Pos := Point{0, 0}
	knot3Pos := Point{0, 0}
	knot4Pos := Point{0, 0}
	knot5Pos := Point{0, 0}
	knot6Pos := Point{0, 0}
	knot7Pos := Point{0, 0}
	knot8Pos := Point{0, 0}
	knot9Pos := Point{0, 0}

	headHistory := make(map[Point]bool)
	knot9 := make(map[Point]bool)

	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, " ")

		dir := decodeDirection(token[0])
		magnitude, _ := strconv.Atoi(token[1])

		for i := 0; i < magnitude; i++ {
			headPos = moveHead(headPos, dir)
			knot1Pos = moveTail(headPos, knot1Pos)
			knot2Pos = moveTail(knot1Pos, knot2Pos)
			knot3Pos = moveTail(knot2Pos, knot3Pos)
			knot4Pos = moveTail(knot3Pos, knot4Pos)
			knot5Pos = moveTail(knot4Pos, knot5Pos)
			knot6Pos = moveTail(knot5Pos, knot6Pos)
			knot7Pos = moveTail(knot6Pos, knot7Pos)
			knot8Pos = moveTail(knot7Pos, knot8Pos)
			knot9Pos = moveTail(knot8Pos, knot9Pos)

			headHistory[headPos] = true
			knot9[knot9Pos] = true
		}

	}
	fmt.Println(len(knot9))

}
