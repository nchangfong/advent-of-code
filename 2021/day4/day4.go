package day4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

type cell struct {
	value  int
	marked bool
}

type board struct {
	spaces []cell
}

// score accepts the last called number and multiplies it
// with the sums of the unmarked spaces on the board. The product
// is returned as an int.
func (b *board) score(called int) int {
	var sum int
	for i := range b.spaces {
		if !b.spaces[i].marked {
			sum += b.spaces[i].value
		}
	}
	return sum * called
}

func (b *board) add(n int) {
	b.spaces = append(b.spaces, cell{value: n})
}

// mark checks the board for values matching n
// and sets their flag to true.
func (b *board) mark(n int) bool {
	for i := range b.spaces {
		if b.spaces[i].value == n {
			b.spaces[i].marked = true
		}
	}

	return false
}

func checkRow(cs []cell) bool {
	for _, c := range cs {
		if !c.marked {
			return false
		}
	}
	return true
}

// check iterates over the columns and rows of the board to
// find a completed segment returning true if there are five consecutive
// marked spaces in a column or row.
func (b *board) check() bool {
	for i := 0; i < 25; i += 5 {
		if checkRow(b.spaces[i : i+5]) {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		var cs []cell
		for j := i; j < 25; j += 5 {
			cs = append(cs, b.spaces[j])
		}
		if checkRow(cs) {
			return true
		}
	}

	return false
}

func newBoard(s []string) board {
	var b board
	long := strings.Join(s, " ")

	nums := strings.Fields(long)

	for _, v := range nums {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		b.add(n)
	}
	return b
}

func Part1(filename string) int {
	lines := helpers.ReadStrings(filename)

	calls := strings.Split(lines[0], ",")
	var boards []board

	for i := 2; i < len(lines); i += 6 {
		b := newBoard(lines[i : i+5])
		boards = append(boards, b)
		if len(b.spaces) != 25 {
			fmt.Printf("error line %d", i)
		}
	}

	for _, c := range calls {
		call, _ := strconv.Atoi(c)
		for _, b := range boards {
			b.mark(call)
			if b.check() {
				return b.score(call)
			}
		}
	}

	return -1
}

func Part2(filename string) int {
	lines := helpers.ReadStrings(filename)

	calls := strings.Split(lines[0], ",")
	var boards []board

	for i := 2; i < len(lines); i += 6 {
		b := newBoard(lines[i : i+5])
		boards = append(boards, b)
	}

	var scores []int

	winners := make(map[int]bool)

	for _, c := range calls {
		call, _ := strconv.Atoi(c)
		for i, b := range boards {
			b.mark(call)
			if b.check() && !winners[i] {
				winners[i] = true
				scores = append(scores, b.score(call))
			}
		}
	}

	return scores[len(scores)-1]
}
