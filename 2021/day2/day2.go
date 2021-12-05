package day2

import (
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

func parseInstruction1(s string) (int, int) {
	ins := strings.Split(s, " ")
	direction := ins[0]
	magnitude, _ := strconv.Atoi(ins[1])
	var dx, dy int
	switch direction {
	case "forward":
		dx += magnitude
	case "down":
		dy += magnitude
	case "up":
		dy -= magnitude
	}

	return dx, dy
}

func parseInstruction2(s string, a int) (int, int, int) {
	ins := strings.Split(s, " ")
	direction := ins[0]
	magnitude, _ := strconv.Atoi(ins[1])
	var dx, dy, da int
	switch direction {
	case "forward":
		dx += magnitude
		dy += a * magnitude
	case "down":
		da = magnitude
	case "up":
		da = -magnitude
	}

	return dx, dy, da
}

func Part1(filename string) int {
	lines := helpers.ReadStrings(filename)
	var x, y int

	for _, line := range lines {
		dx, dy := parseInstruction1(line)
		x += dx
		y += dy

		if y < 0 {
			y = 0
		}
	}

	return x * y
}

func Part2(filename string) int {
	lines := helpers.ReadStrings(filename)
	var x, y, a int

	for _, line := range lines {
		dx, dy, da := parseInstruction2(line, a)
		x += dx
		y += dy
		a += da
	}

	return x * y
}
