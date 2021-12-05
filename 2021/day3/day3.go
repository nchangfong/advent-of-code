package day3

import (
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

type trie struct {
	root *node
}

type node struct {
	children [2]*node
	count    int
}

func (t *trie) insert(s string) {
	if t.root == nil {
		t.root = &node{}
	}
	digits := strings.Split(s, "")

	current := t.root
	for _, d := range digits {
		i, _ := strconv.Atoi(d)
		current.count++
		if current.children[i] == nil {
			current.children[i] = &node{}
		}
		current = current.children[i]
	}
}

func (n *node) findNext() int {
	if n.children[0] == nil && n.children[1] == nil {
		return -1
	}

	if n.children[0] == nil && n.children[1] != nil {
		return 1
	} else if n.children[0] != nil && n.children[1] == nil {
		return 0
	}

	if n.children[0].count > n.children[1].count {
		return 0
	} else {
		return 1
	}
}

func (n *node) findLeast() int {
	if n.children[0] == nil && n.children[1] == nil {
		return -1
	}

	if n.children[0] == nil && n.children[1] != nil {
		return 1
	} else if n.children[0] != nil && n.children[1] == nil {
		return 0
	}

	if n.children[0].count <= n.children[1].count {
		return 0
	} else {
		return 1
	}
}

func Part1(filename string) uint {
	var gamma, epsilon uint
	input := helpers.ReadStrings(filename)
	full := strings.Join(input, "")
	width := len(input[0])

	counts := make([]int, width)

	for k, v := range full {
		if v == '1' {
			counts[k%width] += 1
		}
	}

	for _, b := range counts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if b > len(input)/2 {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	return gamma * epsilon
}

func Part2(filename string) uint {
	var ox, co uint
	input := helpers.ReadStrings(filename)
	t := trie{}
	for _, v := range input {
		t.insert(v)
	}

	current := t.root
	for {
		i := current.findNext()
		if i < 0 {
			break
		}
		ox = ox << 1
		if i == 1 {
			ox += 1
		}
		current = current.children[i]

	}

	current = t.root
	for {
		i := current.findLeast()
		if i < 0 {
			break
		}
		co = co << 1
		co += uint(i)
		current = current.children[i]
	}

	return ox * co
}
