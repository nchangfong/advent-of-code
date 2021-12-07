package day6

import (
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

func run(days int, start []int) int {
	for i := 0; i < days; i++ {
		for j := 0; j < len(start); j++ {
			if start[j] == 0 {
				start = append(start, 9)
				start[j] = 6
			} else {
				start[j]--
			}
		}
		// fmt.Printf("After %d days: %s\n", i+1, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(start)), ","), "[]"))
	}

	return len(start)
}

func run3(days int, start []int) int {
	group := make([]int, 9)

	for _, v := range start {
		group[v] += 1
	}

	total := len(start)

	for i := 0; i < days; i++ {
		tmp := group[0]
		group[0] = group[1]
		group[1] = group[2]
		group[2] = group[3]
		group[3] = group[4]
		group[4] = group[5]
		group[5] = group[6]
		group[6] = group[7] + tmp
		group[7] = group[8]
		group[8] = tmp

		total += tmp
	}

	return total
}

func Part1(filename string, days int) int {
	lines := helpers.ReadStrings(filename)
	var state []int

	for _, line := range lines {

		values := strings.Split(line, ",")

		for _, v := range values {
			i, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			state = append(state, i)
		}

	}
	count := run(days, state)
	return count
}

func Part2(filename string, days int) int {
	lines := helpers.ReadStrings(filename)
	var state []int

	for _, line := range lines {

		values := strings.Split(line, ",")

		for _, v := range values {
			i, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			state = append(state, i)
		}

	}
	count := run3(days, state)
	return count
}
