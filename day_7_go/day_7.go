package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseCommand(input string) {
	tokens := strings.Split(input, " ")

	if tokens[1] == "ls" {

	} else if tokens[1] == "cd" {
	} else {
		log.Fatal("Invalid command")
	}
}

func main() {
	instructions, err := os.Open("input/sample.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			parseCommand(line)
		}
		// fmt.Println(line)
	}

}
