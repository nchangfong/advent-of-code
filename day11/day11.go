package day11

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

const emptySeat string = "L"
const occupiedSeat string = "#"
const floor string = "."

func readInput(filename string) [][]string {
	file, err := os.Open(path.Join("testdata", filename))

	if err != nil {
		log.Fatal(err)
	}

	res := [][]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		res = append(res, row)
	}

	return res
}

func maxInt(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func minInt(l, r int) int {
	if l < r {
		return l
	}
	return r
}

func checkNeighbours(i int, j int, a [][]string) string {
	res := a[i][j]
	if res == floor {
		return floor
	}

	xMin := maxInt(j-1, 0)
	xMax := minInt(j+1, len(a[i])-1)
	yMin := maxInt(i-1, 0)
	yMax := minInt(i+1, len(a)-1)
	adj := 0

	for row := yMin; row <= yMax; row++ {
		for col := xMin; col <= xMax; col++ {
			if row == i && col == j {
				continue
			}
			if a[row][col] == occupiedSeat {
				adj++
			}
		}
	}
	if res == emptySeat && adj == 0 {
		res = occupiedSeat
	} else if res == occupiedSeat && adj >= 4 {
		res = emptySeat
	}
	return res
}

func count(t string, a [][]string) int {
	res := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == t {
				res++
			}
		}
	}
	return res
}

func part1(a [][]string) int {
	m := len(a)
	n := len(a[0])
	b := [][]string{}
	changed := false
	b = make([][]string, m)
	for i := 0; i < m; i++ {
		b[i] = make([]string, n)
		for j := 0; j < n; j++ {
			b[i][j] = checkNeighbours(i, j, a)
			if !changed && b[i][j] != a[i][j] {
				changed = true // if one is updated, need to continue
			}
		}
	}

	if changed {
		return (part1(b))
	}

	return count("#", a)
}
