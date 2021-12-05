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
