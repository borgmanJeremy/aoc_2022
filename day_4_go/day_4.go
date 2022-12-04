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

func getPairs(input string) []string {
	pairs := strings.Split(input, ",")
	return pairs
}

func getRange(input string) (int, int) {
	val := strings.Split(input, "-")
	low, _ := strconv.Atoi(val[0])
	high, _ := strconv.Atoi(val[1])

	return low, high
}

func findFullOverlap(low_1, high_1, low_2, high_2 int) bool {
	if low_1 > high_2 || low_2 > high_1 {
		return false
	}
	if low_1 <= low_2 && high_1 >= high_2 {
		return true
	}
	if low_2 <= low_1 && high_2 >= high_1 {
		return true
	}
	return false
}

func findAnyOverlap(low_1, high_1, low_2, high_2 int) bool {
	if low_1 > high_2 || low_2 > high_1 {
		return false
	}
	return true
}

func main() {
	f, err := os.Open("input/input.txt")
	checkError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count_p1 := 0
	count_p2 := 0
	for scanner.Scan() {
		pairs := getPairs(scanner.Text())
		low_1, high_1 := getRange(pairs[0])
		low_2, high_2 := getRange(pairs[1])

		if findFullOverlap(low_1, high_1, low_2, high_2) {
			count_p1 += 1
		}
		if findAnyOverlap(low_1, high_1, low_2, high_2) {
			count_p2 += 1
		}
	}

	fmt.Println("Part 1: ", count_p1)
	fmt.Println("Part 2: ", count_p2)
}
