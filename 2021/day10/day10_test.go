package day10

import "testing"

func TestPart1(t *testing.T) {
	if got := Part1("example.txt"); got != 26397 {
		t.Errorf("got %d, wanted %d\n", got, 26397)
	}
	t.Logf("Day 10, Part 1: %d\n", Part1("input.txt"))
}
func TestPart2(t *testing.T) {
	if got := Part2("example.txt"); got != 288957 {
		t.Errorf("got %d, wanted %d\n", got, 288957)
	}
	t.Logf("Day 10, Part 2 %d\n", Part2("input.txt"))
}
