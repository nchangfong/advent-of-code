package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

type point struct {
	x int
	y int
}

func parsePoint(s string) point {
	tokens := strings.Split(s, ",")
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return point{x, y}
}

func parseLine(s string) (point, point) {
	tks := strings.Split(s, " -> ")
	p1 := parsePoint(tks[0])
	p2 := parsePoint(tks[1])

	return p1, p2

}

func colinear(p1, p2 point) bool {
	if p1.x == p2.x || p1.y == p2.y {
		return true
	}

	return false
}

type grid struct {
	plane []int
	width int
}

func newGrid(width int) grid {
	p := make([]int, width*width)
	return grid{plane: p, width: width}
}

func (g *grid) incCol(p1, p2 point) {
	x := p1.x
	start := p1.y
	end := p2.y
	if p1.y > p2.y {
		start = p2.y
		end = p1.y
	}

	for y := start; y <= end; y++ {
		g.plane[(g.width*y)+x]++
	}
}

func (g *grid) incLine(p1, p2 point) {

	start := p1.x
	end := p2.x
	if p1.x > p2.x {
		start = p2.x
		end = p1.x
	}

	m := (p2.y - p1.y) / (p2.x - p1.x)
	b := p2.y - m*p2.x

	for x := start; x <= end; x++ {
		y := m*x + b
		g.plane[(g.width*y)+x]++
	}
}

func (g *grid) draw(p1, p2 point) {
	if p1.x == p2.x {
		g.incCol(p1, p2)
	} else {
		g.incLine(p1, p2)
	}
}

func (g *grid) max() int {
	count := 0

	for _, v := range g.plane {
		if v > 1 {
			count++
		}
	}

	return count
}

func (g *grid) String() string {
	var sb strings.Builder
	for i := 0; i < len(g.plane); i++ {
		sb.WriteString(fmt.Sprint(g.plane[i]))
		if (i+1)%g.width == 0 && i > 0 {
			sb.WriteString("\r\n")
			continue
		}
		sb.WriteString(" ")
	}

	return sb.String()
}

func Part1(filename string, width int) int {
	lines := helpers.ReadStrings(filename)
	g := newGrid(width)
	for _, line := range lines {
		p1, p2 := parseLine(line)

		if !colinear(p1, p2) {
			continue
		}

		g.draw(p1, p2)
	}
	return g.max()

}

func Part2(filename string, width int) int {
	lines := helpers.ReadStrings(filename)
	g := newGrid(width)
	for _, line := range lines {
		p1, p2 := parseLine(line)

		g.draw(p1, p2)
	}
	return g.max()
}
