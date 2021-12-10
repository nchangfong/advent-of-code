package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nchangfong/aoc2020/internal/helpers"
)

var m map[string]int
var ref map[string]string

func init() {
	m = map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	ref = map[string]string{
		"35556667":  "a",
		"456667":    "b",
		"23455667":  "c",
		"4555667":   "d",
		"5667":      "e",
		"234556667": "f",
		"5556667":   "g",
	}

}

func makeKey(in []string, lookup map[string]string) map[string]string {
	m := make(map[string][]string)

	for _, word := range in {
		tokens := strings.Split(word, "")
		for _, v := range tokens {
			m[v] = append(m[v], fmt.Sprint(len(word)))
		}
	}

	output := make(map[string]string)
	// fmt.Println(m)
	for k := range m {
		sort.Strings(m[k])
		output[k] = lookup[strings.Join(m[k], "")]
	}

	return output
}

func decode(input []string, key map[string]string) []int {

	var res []int
	for _, v := range input {
		tokens := strings.Split(v, " ")

		var inter []string
		for _, j := range tokens {
			s := strings.Split(j, "")
			for _, q := range s {
				inter = append(inter, key[q])
			}
		}

		sort.Strings(inter)

		res = append(res, m[strings.Join(inter, "")])
	}

	return res
}

func Part1(filename string) int {

	var key map[string]string

	lines := helpers.ReadStrings(filename)
	var score int
	for _, line := range lines {
		fields := strings.Split(line, "|")

		input := strings.Fields(fields[0])
		output := strings.Fields(fields[1])
		key = makeKey(input, ref)

		for _, v := range decode(output, key) {
			switch v {
			case 1, 4, 7, 8:
				score++
			}
		}
	}
	return score
}

func Part2(filename string) int {

	var key map[string]string

	lines := helpers.ReadStrings(filename)
	var score int
	for _, line := range lines {
		fields := strings.Split(line, "|")

		input := strings.Fields(fields[0])
		output := strings.Fields(fields[1])
		key = makeKey(input, ref)

		b, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(decode(output, key))), ""), "[]"))
		score += b
	}
	return score
}
