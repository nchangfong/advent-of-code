package day9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

var rowLen int

type grid []int

func (g grid) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for i, v := range g {
		if i%rowLen == 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprint(v))
		sb.WriteString(" ")

	}

	return sb.String()
}

func (g grid) getNeighbours(i int) []int {
	neighbours := []int{i - 1, i + 1, i - rowLen, i + rowLen}
	var candidates []int
	max := len(g) - 1
	min := 0
	for _, v := range neighbours {
		if v > max || v < min {
			continue
		}
		// first column, ignore left value
		if i%rowLen == 0 && v == i-1 {
			continue
		}
		// last column, ignore right value
		if i%(rowLen) == rowLen-1 && v == i+1 && i != 0 {
			continue
		}
		candidates = append(candidates, v)
	}
	return candidates
}

func (g grid) isLow(i int) bool {
	depth := g[i]

	candidates := g.getNeighbours(i)

	for _, c := range candidates {
		if g[c] <= depth {
			return false
		}
	}

	return true
}

func (g grid) search() int {
	visited := make(map[int]bool)
	var stack helpers.IntStack

	var results []int

	for i := 0; i < len(g); i++ {
		sum := 0
		stack.Push(i)
		for {
			if stack.IsEmpty() {
				break
			}
			n, _ := stack.Pop()

			if g[n] == 9 || visited[n] {
				visited[n] = true
				continue
			}

			visited[n] = true

			sum++

			next := g.getNeighbours(n)

			for j := range next {
				if visited[next[j]] {
					continue
				}

				stack.Push(next[j])
			}
		}

		// don't introduce 0s into result set
		if sum == 0 {
			continue
		}

		results = append(results, sum)
	}

	sort.Ints(results)

	product := 1
	// var limit int
	for i := len(results) - 1; i >= 0 && i >= len(results)-3; i-- {
		product *= results[i]
	}
	return product
}

func Part1(filename string) int {
	var grid grid
	lines := helpers.ReadStrings(filename)
	rowLen = len(strings.Split(lines[0], ""))
	for _, line := range lines {
		depthStrings := strings.Split(line, "")
		rowLen = len(depthStrings)
		var depths []int
		for _, d := range depthStrings {
			i, _ := strconv.Atoi(d)
			depths = append(depths, i)
		}
		grid = append(grid, depths...)
	}
	var sum int
	for i := range grid {
		if grid.isLow(i) {
			sum += grid[i] + 1
		}
	}

	return sum
}

func Part2(filename string) int {
	var grid grid
	lines := helpers.ReadStrings(filename)
	rowLen = len(strings.Split(lines[0], ""))
	for _, line := range lines {
		depthStrings := strings.Split(line, "")
		rowLen = len(depthStrings)
		var depths []int
		for _, d := range depthStrings {
			i, _ := strconv.Atoi(d)
			depths = append(depths, i)
		}
		grid = append(grid, depths...)
	}

	return grid.search()
}
