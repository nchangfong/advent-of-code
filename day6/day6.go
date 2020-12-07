package day6

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

func uniqueChars(input string, m map[string]bool) {
	s := strings.Split(input, "")

	for _, v := range s {
		m[v] = true
	}

}

func readInput() []string {
	file, err := os.Open(path.Join("testdata", "input"))

	if err != nil {
		log.Fatal(err)
	}

	var res []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func run() []int {
	lines := readInput()
	part1 := 0
	part2 := 0
	m := make(map[string]int)
	groupSize := 0
	for _, line := range lines {
		if line != "" {
			for _, v := range strings.Split(line, "") {
				m[v]++
			}
			groupSize++
		} else {
			for _, v := range m {
				if v == groupSize {
					part2++
				}
			}
			part1 += len(m)
			groupSize = 0
			m = make(map[string]int)
		}

	}
	return []int{part1, part2}
}
