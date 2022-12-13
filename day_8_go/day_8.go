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

type mapElement struct {
	height  int
	visible bool
}

func printMap(treeMap [][]mapElement) {
	// print map
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			if treeMap[i][j].visible == true {
				fmt.Printf("\033[1;34m%d\033[0m ", treeMap[i][j].height)
			} else {
				fmt.Print(treeMap[i][j].height, " ")
			}
		}
		fmt.Println()
	}

	// fmt.Println()
	// for i := 0; i < len(treeMap); i++ {
	// 	for j := 0; j < len(treeMap[i]); j++ {
	// 		if treeMap[i][j].visible == true {
	// 			fmt.Print(treeMap[i][j].height, " ")
	// 		} else {
	// 			fmt.Print(" ", " ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)
	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)

	// Read in map
	treeMap := make([][]mapElement, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]mapElement, 0)
		for i := 0; i < len(line); i++ {
			row = append(row, mapElement{int(line[i]) - 48, false})
		}
		treeMap = append(treeMap, row)
	}

	// Mark Edges visible
	for i := 0; i < len(treeMap[0]); i++ {
		treeMap[0][i].visible = true
		treeMap[len(treeMap)-1][i].visible = true
	}

	for i := 0; i < len(treeMap[0]); i++ {
		treeMap[i][0].visible = true
		treeMap[i][len(treeMap)-1].visible = true
	}

	// Fill forward row visibility
	for i := 1; i < len(treeMap); i++ {
		highest := treeMap[i][0].height
		for j := 1; j < len(treeMap[i]); j++ {
			if treeMap[i][j].height > highest {
				treeMap[i][j].visible = true
				highest = treeMap[i][j].height
			}
		}
	}

	// Fill reverse row visibility
	for i := 1; i < len(treeMap); i++ {
		highest := treeMap[i][len(treeMap)-1].height
		for j := len(treeMap[i]) - 1; j > 0; j-- {
			if treeMap[i][j-1].height > highest {
				treeMap[i][j-1].visible = true
				highest = treeMap[i][j-1].height
			}
		}
	}

	// Fill downward column visibility
	for i := 1; i < len(treeMap); i++ {
		highest := treeMap[0][i].height
		for j := 1; j < len(treeMap[i]); j++ {
			if treeMap[j][i].height > highest {
				treeMap[j][i].visible = true
				highest = treeMap[j][i].height
			}
		}
	}

	// Fill upward column visibility
	for i := 1; i < len(treeMap); i++ {
		highest := treeMap[len(treeMap)-1][i].height
		for j := len(treeMap[i]) - 1; j > 0; j-- {
			if treeMap[j-1][i].height > highest {
				treeMap[j-1][i].visible = true
				highest = treeMap[j-1][i].height
			}
		}
	}

	printMap(treeMap)

	// Count visible
	visibleCount := 0
	for i := 0; i < len(treeMap); i++ {
		for j := 0; j < len(treeMap[i]); j++ {
			if treeMap[i][j].visible == true {
				visibleCount++
			}
		}
	}
	fmt.Println("\nPart 1 ", visibleCount)
}
