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

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}

}
