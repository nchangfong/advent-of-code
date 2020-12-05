package day5

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"sort"
	"strings"
)

const rowMin int = 0
const rowMax int = 127
const colMin int = 0
const colMax int = 7

func readInput(filename string) [][]string {
	file, err := os.Open(path.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	res := [][]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "")
		res = append(res, tokens)
	}

	return res
}

func findRow(ticket []string) int {
	r := rowMax
	l := rowMin
	res := 0
	for _, s := range ticket {

		switch s {
		case "F":
			r = int(math.Floor(float64((r + l)) / float64(2.0)))
			res = r
		case "B":
			l = int(math.Ceil(float64((r + l)) / float64(2.0)))
			res = l
		default:
			return res
		}
	}
	return res
}

func findCol(ticket []string) int {
	r := colMax
	l := colMin
	res := 0
	for _, s := range ticket {

		switch s {
		case "L":
			r = int(math.Floor(float64((r + l)) / float64(2.0)))
			res = r
		case "R":
			l = int(math.Ceil(float64((r + l)) / float64(2.0)))
			res = l
		default:
			return res
		}
	}
	return res
}

func findSeat(ticket []string) int {
	return (findRow(ticket[:7]) * (colMax + 1)) + findCol(ticket[7:])
}

func run(filename string) []int {
	s := readInput(filename)
	ids := []int{}
	part1 := 0
	part2 := 0
	for _, ticket := range s {
		seat := findSeat(ticket)
		ids = append(ids, seat)
		if seat > part1 {
			part1 = seat
		}
	}

	sort.Ints(ids)
	for k, v := range ids {
		if k < len(ids)-1 && ids[k+1] == v+2 {
			part2 = v + 1
			break
		}
	}

	return []int{part1, part2}
}
