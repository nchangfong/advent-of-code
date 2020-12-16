package day11

import "testing"

func TestPart1(t *testing.T) {
	filename := "input"
	t.Log(part1(readInput(filename)))
	t.Log(part2(readInput(filename)))
}
