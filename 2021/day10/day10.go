package day10

import (
	"sort"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

func firstIncorrect(s string) (rune, []rune) {
	expect := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
		'(': ')',
	}

	var stack helpers.RuneStack
	var firstInc rune
	var complete []rune

	for _, val := range s {
		_, ok := expect[val]
		if ok {
			stack.Push(val)
			continue
		}

		current, ok := stack.Pop()
		if !ok {
			continue
		}

		if expect[current] != val {
			firstInc = val
			break
		}
	}

	if firstInc != 0 {
		return firstInc, complete
	}

	for {
		if stack.IsEmpty() {
			break
		}

		r, _ := stack.Pop()

		complete = append(complete, expect[r])
	}

	return firstInc, complete

}

func Part1(filename string) int {
	var scores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var sum int
	lines := helpers.ReadStrings(filename)
	incorrect := make(map[rune]int)
	for _, line := range lines {
		inc, _ := firstIncorrect(line)
		if inc == 0 {
			continue
		}
		incorrect[inc] += scores[inc]
	}

	for _, v := range incorrect {
		sum += v
	}

	return sum
}

func scoreRunes(complete []rune) int {
	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var score int
	for _, r := range complete {
		score = (score * 5) + scores[r]
	}
	return score
}

func Part2(filename string) int {
	lines := helpers.ReadStrings(filename)
	var scores []int
	for _, line := range lines {
		inc, complete := firstIncorrect(line)
		if inc != 0 {
			continue
		}
		scores = append(scores, scoreRunes(complete))
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}
