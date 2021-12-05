package day1

import (
	"fmt"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

func Part1() int {
	lines, err := helpers.ReadInts("part1.txt")
	if err != nil {
		fmt.Printf("error: %v", err)
		return -1
	}

	counter := 0

	for i := range lines {
		if i == 0 {
			continue
		}

		if lines[i] > lines[i-1] {
			counter++
		}

	}
	return counter
}

func Part2() int {
	lines, err := helpers.ReadInts("part1.txt")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return -1
	}

	counter := 0

	for i := range lines {
		if i < 3 {
			continue
		}

		if lines[i] > lines[i-3] {
			counter++
		}

	}
	return counter
}
