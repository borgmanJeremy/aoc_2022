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
}

func makeCPU(instructions []string) *CPU {
	cpu := new(CPU)

	cpu.clock = 0
	cpu.regX = 1
	cpu.add = 0
	cpu.strength = 0
	cpu.pendingAdd = false
	cpu.instructionPointer = 0
	cpu.instructions = make([]string, len(instructions))
	copy(cpu.instructions, instructions)
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
	if cpu.clock == 20 {
		fmt.Println("clock: ", cpu.clock, "regx: ", cpu.regX)
		cpu.strength += (cpu.clock * cpu.regX)
	} else if (cpu.clock-20)%40 == 0 {
		fmt.Println("clock: ", cpu.clock, "regx: ", cpu.regX)
		cpu.strength += (cpu.clock * cpu.regX)
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
	cpu := makeCPU(instructionList)

	for cpu.instructionPointer < len(cpu.instructions) {
		cpu.tick()
	}
	fmt.Println("strength: ", cpu.strength)
}
