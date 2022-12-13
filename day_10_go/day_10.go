package main

import (
	"bufio"
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

	clockQueue int
	pendingAdd bool

	instructionPointer int
	instructions       []string
}

func makeCPU(instructions []string) *CPU {
	cpu := new(CPU)

	cpu.clock = 0
	cpu.regX = 0
	cpu.pendingAdd = false
	cpu.instructionPointer = 0

	copy(cpu.instructions, instructions)
	return cpu
}

func (cpu *CPU) addx(number int) {
	cpu.pendingAdd = true
	cpu.clockQueue = 2
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

	if cpu.pendingAdd {
		if cpu.clockQueue > 0 {
			cpu.clockQueue -= 1
		} else {
			cpu.regX += 1
			cpu.pendingAdd = false
		}
	} else {
		cpu.decode(cpu.instructions[cpu.instructionPointer])
		cpu.instructionPointer += 1
	}

	cpu.clock += 1
}

func main() {
	instructions, err := os.Open("input/sample.txt")
	checkError(err)

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	instructionList := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instructionList = append(instructionList, line)
	}

	// create new cpu
	cpu := makeCPU(instructionList)
	cpu.tick()

}
