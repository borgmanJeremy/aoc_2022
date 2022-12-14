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

type Operand int

const (
	Add Operand = iota
	Mult
)

type OperationData struct {
	operator Operand
	value    string
}

type Monkey struct {
	items           []int
	operation       OperationData
	testConstant    int
	trueMonkey      int
	falseMonkey     int
	inspectionCount int
}

func parseItems(scanner *bufio.Scanner) []int {
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, ":")
	nums := strings.Split(tokens[1], ",")
	for i := 0; i < len(nums); i++ {
		nums[i] = strings.Trim(nums[i], " ")
	}
	itemList := make([]int, 0)
	for _, num := range nums {
		iNum, _ := strconv.Atoi(num)
		itemList = append(itemList, iNum)
	}
	return itemList
}

func parseOperationData(scanner *bufio.Scanner) OperationData {
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, ":")
	operands := strings.Split(tokens[1], "=")
	for i := 0; i < len(operands); i++ {
		operands[i] = strings.Trim(operands[i], " ")
	}
	rhsTokens := strings.Split(operands[1], " ")
	var operationData OperationData
	if rhsTokens[1] == "+" {
		operationData = OperationData{operator: Add, value: rhsTokens[2]}
	} else if rhsTokens[1] == "*" {
		operationData = OperationData{operator: Mult, value: rhsTokens[2]}
	}
	return operationData
}

func parseTestConstant(scanner *bufio.Scanner) int {
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	divisor, _ := strconv.Atoi(tokens[len(tokens)-1])
	return divisor
}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	monkeyList := make([]Monkey, 0)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if tokens[0] == "Monkey" {
			items := parseItems(scanner)
			operationData := parseOperationData(scanner)
			testConstant := parseTestConstant(scanner)

			trueClause := parseTestConstant(scanner)
			falseClause := parseTestConstant(scanner)

			monkey := Monkey{
				items:           items,
				operation:       operationData,
				testConstant:    testConstant,
				trueMonkey:      trueClause,
				falseMonkey:     falseClause,
				inspectionCount: 0,
			}
			monkeyList = append(monkeyList, monkey)
		}
	}

	for i := 0; i < 20; i++ {
		// Process Monkeys
		for idx, monkey := range monkeyList {
			for _, item := range monkey.items {
				monkeyList[idx].inspectionCount++
				var num int
				if monkey.operation.value == "old" {
					num = item
				} else {
					num, _ = strconv.Atoi(monkey.operation.value)
				}

				// Inspect Item
				if monkey.operation.operator == Add {
					item += num
				} else if monkey.operation.operator == Mult {
					item *= num
				}

				// Decrease Worry
				item = item / 3

				// Pass item
				if item%monkey.testConstant == 0 {
					monkeyList[monkey.trueMonkey].items = append(monkeyList[monkey.trueMonkey].items, item)
				} else {
					monkeyList[monkey.falseMonkey].items = append(monkeyList[monkey.falseMonkey].items, item)
				}
			}
			monkeyList[idx].items = make([]int, 0)
		}
	}

	for idx, monkey := range monkeyList {
		fmt.Printf("Monkey %d inspected %d items\n", idx, monkey.inspectionCount)
	}
}
