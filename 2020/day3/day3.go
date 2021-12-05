package day3

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readFile() [][]string {
	var res [][]string
	filename := "input"
	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "")
		res = append(res, points)
	}
	return res
}

type pair struct {
	x int
	y int
}

func prod(s []int) int {
	res := 1
	for _, v := range s {
		res *= v
	}
	return res
}

func ski() []int {
	slope := readFile()

	pairs := []pair{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := []int{}
	for _, move := range pairs {
		count := 0
		pos := 0
		for i := 0; i < len(slope); i += move.y {
			line := slope[i]

			if line[pos] == "#" {
				count++
			}

			pos = (pos + move.x) % len(line)

		}
		res = append(res, count)

	}
	return []int{res[1], prod(res)}
}
