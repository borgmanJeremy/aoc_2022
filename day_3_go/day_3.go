package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func splitInHalf(input string) (string, string) {
	length := len(input)
	frontHalf := input[0 : length/2]
	backHalf := input[length/2 : length]

	return frontHalf, backHalf
}

func toNumber(input rune) int {
	num := int(input)
	if num > 64 && num < 91 {
		num = num - 64 + 26
	} else if num > 96 && num < 123 {
		num -= 96
	}
	return num
}

func findFirstDuplicate(input1, input2 string) rune {
	for _, char := range input1 {
		for _, char2 := range input2 {
			if char == char2 {
				return char
			}
		}
	}
	return ' '
}

func findIntersection(input_1, input_2, input_3 string) rune {
	for _, char := range input_1 {
		for _, char2 := range input_2 {
			for _, char3 := range input_3 {
				if char == char2 && char == char3 {
					return char
				}
			}
		}
	}
	return ' '
}

func main() {
	f, err := os.Open("input/input.txt")
	checkError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	count := 0
	score_2 := 0
	group := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf, secondHalf := splitInHalf(line)
		dup := findFirstDuplicate(firstHalf, secondHalf)
		sum += toNumber(dup)

		count += 1
		group = append(group, line)
		if count == 3 {
			score := toNumber(findIntersection(group[0], group[1], group[2]))
			score_2 += score
			count = 0
			group = nil
		}
	}
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", score_2)

}
