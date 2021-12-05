package day13

import "testing"

func TestReadInput(t *testing.T) {
	filename := "input"
	t.Log(part1(readInput(filename)))
	a := part2(readInput(filename))
	t.Log(a.String())
}
