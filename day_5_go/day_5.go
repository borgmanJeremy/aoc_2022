package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	qty int
	src int
	dst int
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseCols(fileName string) [][]rune {
	col_file, err := os.Open(fileName)
	checkError(err)

	defer col_file.Close()

	var col_list [][]rune

	scanner := bufio.NewScanner(col_file)
	col := []rune{}
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			box := rune(scanner.Text()[0])

			col = append(col, box)

		} else {
			col_list = append(col_list, col)
			col = []rune{}
		}
	}
	col_list = append(col_list, col)
	return col_list
}

func parseInstructions(inputFile string) []instruction {
	instructions, err := os.Open(inputFile)
	checkError(err)

	defer instructions.Close()

	var instructionList []instruction
	scanner := bufio.NewScanner(instructions)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		qty, _ := strconv.Atoi(words[1])
		src, _ := strconv.Atoi(words[3])
		dst, _ := strconv.Atoi(words[5])
		instructionList = append(instructionList, instruction{qty, src - 1, dst - 1})
	}
	return instructionList
}

func part_1(col_list [][]rune, instructionList []instruction) {
	for _, instruction := range instructionList {
		for i := 0; i < instruction.qty; i++ {
			value := col_list[instruction.src][0]
			new_col := append([]rune{value}, col_list[instruction.dst]...)
			col_list[instruction.dst] = new_col
			col_list[instruction.src] = col_list[instruction.src][1:]
		}
	}

	// fmt.Printf("%c \n", col_list)
	for _, col := range col_list {
		fmt.Printf("%c", col[0])
	}
	fmt.Println()
}

func part_2(col_list [][]rune, instructionList []instruction) {
	for _, instruction := range instructionList {
		new_vals := append([]rune{}, col_list[instruction.src][0:instruction.qty]...)
		original_vals := col_list[instruction.dst]
		new_col := append(new_vals, original_vals...)

		col_list[instruction.dst] = new_col
		col_list[instruction.src] = col_list[instruction.src][instruction.qty:]
	}

	// fmt.Printf("%c \n", col_list)
	for _, col := range col_list {
		fmt.Printf("%c", col[0])
	}
}

func main() {
	col_list := parseCols("input/puzzle/col.txt")
	instructionList := parseInstructions("input/puzzle/instructions.txt")
	part_1(col_list, instructionList)

	col_list = parseCols("input/puzzle/col.txt")
	instructionList = parseInstructions("input/puzzle/instructions.txt")
	part_2(col_list, instructionList)
}
