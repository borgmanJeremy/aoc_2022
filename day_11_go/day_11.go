package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	items           []uint64
	operation       OperationData
	testConstant    uint64
	trueMonkey      uint64
	falseMonkey     uint64
	inspectionCount int
}

func parseItems(scanner *bufio.Scanner) []uint64 {
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, ":")
	nums := strings.Split(tokens[1], ",")
	for i := 0; i < len(nums); i++ {
		nums[i] = strings.Trim(nums[i], " ")
	}
	itemList := make([]uint64, 0)
	for _, num := range nums {
		iNum, _ := strconv.ParseUint(num, 10, 64)
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

func parseTestConstant(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	divisor, _ := strconv.ParseUint(tokens[len(tokens)-1], 10, 64)
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

	var gcf uint64 = 1
	for _, monkey := range monkeyList {
		gcf *= monkey.testConstant
	}

	for i := 0; i < 10000; i++ {
		// Process Monkeys
		for idx, monkey := range monkeyList {
			for _, item := range monkey.items {
				monkeyList[idx].inspectionCount++
				var num uint64
				if monkey.operation.value == "old" {
					num = item
				} else {
					num, _ = strconv.ParseUint(monkey.operation.value, 10, 64)
				}

				// Inspect Item
				if monkey.operation.operator == Add {
					item += num
				} else if monkey.operation.operator == Mult {
					item *= num
				}

				// Decrease Worry
				//item = item / 3
				item = item % gcf
				// Pass item
				if item%monkey.testConstant == 0 {
					monkeyList[monkey.trueMonkey].items = append(monkeyList[monkey.trueMonkey].items, item)
				} else {
					monkeyList[monkey.falseMonkey].items = append(monkeyList[monkey.falseMonkey].items, item)
				}
			}
			monkeyList[idx].items = make([]uint64, 0)
		}
	}

	sortedInspectionCount := make([]int, 0)

	for _, monkey := range monkeyList {
		sortedInspectionCount = append(sortedInspectionCount, monkey.inspectionCount)
		fmt.Println(monkey.inspectionCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortedInspectionCount)))
	for _, count := range sortedInspectionCount {
		_ = count
		//fmt.Println(count)
	}
	fmt.Println("solution: ", sortedInspectionCount[0]*sortedInspectionCount[1])

}
