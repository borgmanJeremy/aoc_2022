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

type OpCode int

const (
	noop OpCode = iota
	addx
)

func decodeOpCode(input string) OpCode {
	switch input {
	case "addx":
		return addx
	case "noop":
		return noop
	}
	log.Fatal("Unknown OpCode")
	return 0
}

type CPU struct {
	clock int
	regX  int
	add   int

	clockQueue int
	pendingAdd bool

	instructionPointer int
	instructions       []string

	strength int

	monitor *[][]rune
}

func makeCPU(instructions []string, monitor *[][]rune) *CPU {
	cpu := new(CPU)

	cpu.clock = 0
	cpu.regX = 1
	cpu.add = 0
	cpu.strength = 0
	cpu.pendingAdd = false
	cpu.instructionPointer = 0
	cpu.instructions = make([]string, len(instructions))
	copy(cpu.instructions, instructions)

	cpu.monitor = monitor

	return cpu
}

func (cpu *CPU) addx(number int) {
	cpu.pendingAdd = true
	cpu.clockQueue = 1
	cpu.add = number
}

func (cpu *CPU) decode(instruction string) {
	tokens := strings.Split(instruction, " ")
	opCode := decodeOpCode(tokens[0])

	switch opCode {
	case addx:
		number, _ := strconv.Atoi(tokens[1])
		_ = number
		cpu.addx(number)
	case noop:
	}
}

func (cpu *CPU) tick() {

	cpu.clock += 1
	// Part 1 calculations
	if cpu.clock == 20 {
		cpu.strength += (cpu.clock * cpu.regX)
	} else if (cpu.clock-20)%40 == 0 {
		cpu.strength += (cpu.clock * cpu.regX)
	}

	// Part 2 calculations
	row := ((cpu.clock - 1) / 40)
	col := ((cpu.clock - 1) % 40)
	spriteCenter := cpu.regX

	if col == spriteCenter || col == spriteCenter+1 || col == spriteCenter-1 {
		(*cpu.monitor)[row][col] = '#'
	}

	if cpu.pendingAdd {
		cpu.clockQueue -= 1
		if cpu.clockQueue == 0 {
			cpu.regX += cpu.add
			cpu.pendingAdd = false
		}
	} else {
		cpu.decode(cpu.instructions[cpu.instructionPointer])
		cpu.instructionPointer += 1
	}

}

func printMonitor(monitor *[][]rune) {
	for _, row := range *monitor {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	instructionList := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instructionList = append(instructionList, line)
	}

	// create new cpu
	monitorHeight := 6
	monitorWidth := 40
	var monitor [][]rune

	for i := 0; i < monitorHeight; i++ {
		monitor = append(monitor, make([]rune, monitorWidth))
		for j := 0; j < monitorWidth; j++ {
			monitor[i][j] = ' '
		}
	}

	cpu := makeCPU(instructionList, &monitor)

	for cpu.instructionPointer < len(cpu.instructions) {
		cpu.tick()
		printMonitor(&monitor)
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("strength: ", cpu.strength)
}
