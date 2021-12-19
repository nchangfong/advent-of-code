package helpers

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadStrings(filename string) []string {
	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadInts(filename string) ([]int, error) {
	file, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, nil
}

type IntStack struct {
	data []int
}

func (s *IntStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *IntStack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	idx := len(s.data) - 1
	res := s.data[idx]
	s.data = s.data[:idx]
	return res, true
}

func (s *IntStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}
