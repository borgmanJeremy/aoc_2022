package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkUnique(input string) bool {
	charList := strings.Split(input, "")
	sort.Strings(charList)

	for idx := 0; idx < len(charList)-1; idx++ {
		if charList[idx] == charList[idx+1] {
			return false
		}
	}
	return true
}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	scanner.Scan()
	line := scanner.Text()
	for idx := 0; idx < len(line)-3; idx++ {
		// fmt.Println("idx: ", idx+4, " ", line[idx:idx+4])
		if checkUnique(line[idx : idx+4]) {
			fmt.Println("Part 1: ", idx+4)
			break
		}
	}
}
