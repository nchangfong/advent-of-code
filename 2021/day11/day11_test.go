package day11

import (
	"reflect"
	"sort"
	"testing"
)

func TestNeighbours(t *testing.T) {
	g := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// 0 1 2
	// 3 4 5
	// 6 7 8
	grid := newGrid(g, 3)

	cases := map[string]struct {
		input int
		want  []int
	}{
		"upper left": {
			input: 0,
			want:  []int{1, 3, 4},
		},
		"upper middle": {
			input: 1,
			want:  []int{0, 2, 3, 4, 5},
		},
		"upper right": {
			input: 2,
			want:  []int{1, 4, 5},
		},
		"middle left": {
			input: 3,
			want:  []int{0, 1, 4, 6, 7},
		},
		"middle": {
			input: 4,
			want:  []int{0, 1, 2, 3, 5, 6, 7, 8},
		},
		"middle right": {
			input: 5,
			want:  []int{1, 2, 4, 7, 8},
		},
		"bottom left": {
			input: 6,
			want:  []int{3, 4, 7},
		},
		"bottom middle": {
			input: 7,
			want:  []int{3, 4, 5, 6, 8},
		},
		"bottom right": {
			input: 8,
			want:  []int{4, 5, 7},
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := grid.neighbours(tc.input)
			sort.Ints(actual)
			sort.Ints(tc.want)

			if !reflect.DeepEqual(actual, tc.want) {
				t.Errorf("got %v, want %v", actual, tc.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	t.Logf("Day 11, Part 1: %d\n", Part1("input.txt", 100))
}
func TestPart2(t *testing.T) {
	t.Logf("Day 11, Part 2: %d\n", Part2("input.txt"))
}
