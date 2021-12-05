package day1

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func find2(filename string, total int) int {
	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var m map[int]bool
	m = make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		m[i] = true
		diff := total - i
		if m[diff] {
			return diff * i
		}
	}

	return 0
}

func find3(filename string, total int) int {
	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int]bool)
	l := make([]int, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		m[i] = true
		l = append(l, i)
	}

	for i, first := range l[1 : len(l)-1] {
		for _, second := range l[i+1:] {
			diff := total - first - second
			if m[diff] {
				return diff * first * second
			}
		}
	}

	return 0
}
