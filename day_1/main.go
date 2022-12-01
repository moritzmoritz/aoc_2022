package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	totalCalories int
}

func main() {
	// First part

	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read line by line
	scanner := bufio.NewScanner(file)
	elves := []Elf{}

	calories := 0
	for scanner.Scan() {
		elfCalories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			elves = append(elves, Elf{totalCalories: calories})
			calories = 0
			continue
		}

		calories += elfCalories
	}

	elfMostCalories := elves[0]

	// Find elf with most calories
	for i := 0; i < len(elves); i++ {
		elf := elves[i]
		if elf.totalCalories > elfMostCalories.totalCalories {
			elfMostCalories = elf
		}
	}

	fmt.Println((elfMostCalories))

	sort.Slice(elves[:], func(i, j int) bool {
		return elves[i].totalCalories > elves[j].totalCalories
	})

	// Second part

	// Get total calories of top three elves
	topThreeElves := elves[0:3]
	sum := 0
	for i := 0; i < len(topThreeElves); i++ {
		sum += topThreeElves[i].totalCalories
	}

	fmt.Println(sum)

}
