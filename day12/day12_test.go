package day12

import "testing"

func TestReadInput(t *testing.T) {
	a := readInput("input")
	t.Log(part1(a))
	t.Log(part2(a))
}
