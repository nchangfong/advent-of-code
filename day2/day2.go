package day2

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

func split(r rune) bool {
	return r == ' ' || r == ':' || r == '-'
}

func validatePass(min int, max int, target string, pass string) bool {
	var m map[rune]int
	m = make(map[rune]int)
	conv := []rune(target)
	t := conv[0]
	for i, w := 0, 0; i < len(pass); i += w {
		c, width := utf8.DecodeRuneInString(pass[i:])
		m[c]++
		w = width
	}

	return m[t] >= min && m[t] <= max
}

func validatePass2(min int, max int, target string, pass string) bool {
	return (pass[min-1] == target[0]) != (pass[max-1] == target[0])
}

func passwordMatches(filename string) []int {
	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1 := 0
	part2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, split)
		min, _ := strconv.Atoi(fields[0])
		max, _ := strconv.Atoi(fields[1])
		target, pass := fields[2], fields[3]

		if validatePass(min, max, target, pass) {
			part1++
		}

		if validatePass2(min, max, target, pass) {
			part2++
		}
	}

	return []int{part1, part2}
}
