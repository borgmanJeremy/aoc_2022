package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open("input/input.txt")
	checkError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var calorieList []int
	var count int = 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			num, err := strconv.Atoi(scanner.Text())
			checkError(err)
			count += num
		} else {
			calorieList = append(calorieList, count)
			count = 0
		}
	}

	calorieList = append(calorieList, count)
	sort.Sort(sort.Reverse(sort.IntSlice(calorieList)))
	fmt.Println("Part 1: ", calorieList[0])
	fmt.Println("Part 2: ", calorieList[0]+calorieList[1]+calorieList[2])

}
