package day14

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestInput(t *testing.T) {
	file, err := os.Open(path.Join("testdata", "input"))

	check(err)
	defer file.Close()

	fmt.Println(readInput(file))

}

func set() {

}

func clear(int, string) {

}

func TestPart1(t *testing.T) {
	file, err := os.Open(path.Join("testdata", "input"))

	check(err)
	defer file.Close()

	input := readInput(file)

	part1(input)

}
