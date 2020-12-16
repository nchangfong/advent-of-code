package day10

import (
	"bufio"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
)

func readInput(filename string) []int {
	file, err := os.Open(path.Join("testdata", filename))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	res := []int{0}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, val)
	}

	sort.Ints(res)
	res = append(res, res[len(res)-1]+3)

	return res
}

func part1(a []int) int {
	ones, threes := 0, 0

	for i := 1; i < len(a); i++ {
		diff := a[i] - a[i-1]

		switch diff {
		case 1:
			ones++
		case 3:
			threes++
		}
	}

	return ones * threes
}

func part2(a []int) int {
	count := 1
	run := 0

	for i := 1; i < len(a); i++ {
		diff := a[i] - a[i-1]
		if diff != 1 {
			switch run {
			case 2:
				count *= 2
			case 3:
				count *= 4
			case 4:
				count *= 7
			}
			run = 0
		} else {
			run++
		}
	}

	return count
}
