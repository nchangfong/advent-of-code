package day14

import (
	"fmt"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

func makeBigrams(s []string) map[string]int {
	res := make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		var sb strings.Builder
		sb.WriteString(s[i])
		sb.WriteString(s[i+1])
		key := sb.String()

		res[key] += 1
	}

	return res
}

// sumChars counts the occurrences of each character in the keyset of the input map.
func sumChars(m map[string]int) map[string]int {
	counts := make(map[string]int)
	for key, val := range m {
		key := key
		val := val
		chars := strings.Split(key, "")
		for _, c := range chars {
			counts[c] += val
		}
	}
	return counts
}

// generator accepts a bigram bg and a character c applying the rule ab -> ac, cb.
// The results of the rule are returned in a string slice.
func generator(bg string, c string) []string {
	var res []string
	chars := strings.Split(bg, "")

	var sb strings.Builder
	sb.WriteString(chars[0])
	sb.WriteString(c)
	res = append(res, sb.String())
	sb.Reset()
	sb.WriteString(c)
	sb.WriteString(chars[1])
	res = append(res, sb.String())

	return res
}

func Part1(filename string, n int) int {
	lines := helpers.ReadStrings(filename)

	state := strings.Split(lines[0], "")

	lines = lines[2:]

	bigrams := makeBigrams(state)

	rules := make(map[string]string)

	for _, line := range lines {
		line := line
		fields := strings.Fields(line)
		rules[fields[0]] = fields[2]
	}

	for i := 0; i < n; i++ {

		m := make(map[string]int)
		// a fixed copy of the input set bc following loop mutates the set, affecting number of iterations
		for key := range bigrams {
			m[key] = bigrams[key]
		}

		for k, count := range m {
			r, ok := rules[k]
			if !ok {
				fmt.Println("not found ", k)
				continue
			}

			bgs := generator(k, r)

			for _, v := range bgs {
				bigrams[v] += count
			}

			bigrams[k] -= count

		}
	}

	counts := sumChars(bigrams)
	first := state[0]
	last := state[len(state)-1]

	// reduce overlapping bigrams
	for key := range counts {
		if key == first || key == last {
			counts[key] = (counts[key] / 2) + 1
		} else {
			counts[key] = (counts[key] / 2)
		}
	}

	max := 0
	min := -1 // assumes values aren't below zero

	for _, v := range counts {
		if v > max {
			max = v
		}
		if min == -1 || v < min {
			min = v
		}
	}

	return max - min
}
