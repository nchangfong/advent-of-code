package day5

import (
	"strings"
	"testing"
)

func TestFindRow(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"FBFBBFF", 44},
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
	}

	for _, c := range tests {
		input := strings.Split(c.input, "")
		actual := findRow(input)
		if actual != c.expected {
			t.Errorf("findRow(%s) = %d, expected %d", c.input, actual, c.expected)
		}
	}

}

func TestFindCol(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"RLR", 5},
		{"RRR", 7},
		{"RLL", 4},
	}

	for _, c := range tests {
		input := strings.Split(c.input, "")
		actual := findCol(input)
		if actual != c.expected {
			t.Errorf("findCol(%s) = %d, expected %d", c.input, actual, c.expected)
		}
	}
}

func TestFindSeat(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, c := range tests {
		input := strings.Split(c.input, "")
		actual := findSeat(input)
		if actual != c.expected {
			t.Errorf("findSeat(%s) = %d, expected %d", c.input, actual, c.expected)
		}
	}
}

func TestPart1(t *testing.T) {
	r := run("input")
	t.Logf("part1: %v, part2: %v", r[0], r[1])
}
