package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

func newGrid(d []int, rowLen int) grid {
	return grid{
		d:      d,
		rowLen: rowLen,
	}
}

type grid struct {
	d      []int
	rowLen int
}

func (g *grid) String() string {
	var sb strings.Builder
	for i, v := range g.d {
		v := v
		if i%g.rowLen == 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprint(v))
		sb.WriteString(" ")

	}

	return sb.String()
}

func (g *grid) neighbours(i int) []int {
	neighbours := []int{i - 1, i + 1, i - g.rowLen, i - g.rowLen + 1, i - g.rowLen - 1, i + g.rowLen, i + g.rowLen + 1, i + g.rowLen - 1}
	var candidates []int
	max := len(g.d) - 1
	min := 0
	for _, v := range neighbours {
		if v > max || v < min {
			continue
		}
		// left column, ignore left values
		if i%g.rowLen == 0 && (v == i-1 || v == i+g.rowLen-1 || v == i-g.rowLen-1) {
			continue
		}
		// right column, ignore right values
		if i%(g.rowLen) == g.rowLen-1 && (v == i+1 || v == i-g.rowLen+1 || v == i+g.rowLen+1) && i != 0 {
			continue
		}
		candidates = append(candidates, v)
	}
	return candidates
}

func (g *grid) Step() int {
	var stack helpers.IntStack
	flashed := make(map[int]bool)
	// increment all
	for i := range g.d {
		g.d[i]++
		if g.d[i] > 9 {
			stack.Push(i)
		}
	}

	for {
		if stack.IsEmpty() {
			break
		}

		current, _ := stack.Pop()
		if flashed[current] {
			continue
		}

		flashed[current] = true

		next := g.neighbours(current)

		for _, v := range next {
			g.d[v]++
			_, ok := flashed[v]
			if g.d[v] > 9 && !ok {
				stack.Push(v)
			}
		}
	}

	for i := range g.d {
		if g.d[i] > 9 {
			g.d[i] = 0
		}
	}

	return len(flashed)
}

func Part1(filename string, steps int) int {
	lines := helpers.ReadStrings(filename)
	if len(lines) == 0 {
		return 0
	}
	var vals []int

	for i := range lines {
		for _, r := range lines[i] {
			num, _ := strconv.Atoi(fmt.Sprintf("%c", r))
			vals = append(vals, num)
		}
	}

	g := newGrid(vals, len(lines[0]))

	var flashes int

	for i := 0; i < steps; i++ {
		step := g.Step()
		flashes += step
	}

	return flashes
}

func Part2(filename string) int {
	lines := helpers.ReadStrings(filename)
	if len(lines) == 0 {
		return 0
	}
	var vals []int

	for i := range lines {
		for _, r := range lines[i] {
			num, _ := strconv.Atoi(fmt.Sprintf("%c", r))
			vals = append(vals, num)
		}
	}

	g := newGrid(vals, len(lines[0]))

	var i int
	for {
		i++
		step := g.Step()
		if step == len(vals) {
			break
		}
	}

	return i
}
