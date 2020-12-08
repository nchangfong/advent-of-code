package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type lookup struct {
	data  map[string]int
	index int
}

func (l *lookup) new() *lookup {
	l.data = make(map[string]int)
	l.index = 0

	return l
}

func (l *lookup) getIndex(s string) int {
	if l.data == nil {
		l.data = make(map[string]int)
	}
	_, ok := l.data[s]
	if !ok {
		l.data[s] = l.index
		l.index++
	}

	return l.data[s]
}

func readInput() int {
	file, err := os.Open(path.Join("testdata", "input"))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lookupTable := new(lookup)
	records := []string{}
	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	size := len(records)

	// adjacency matrix?
	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}

	for _, line := range records {
		fields := strings.Split(line, "bags contain")
		colour := fmt.Sprintf("%s", strings.TrimSpace(fields[0]))

		// add to lookup table
		node := lookupTable.getIndex(colour)

		fields = strings.Split(fields[1], ",")

		for _, v := range fields {
			items := strings.Split(strings.TrimSpace(v), " ")
			var capacity int
			if items[0] == "no" {
				continue
			}

			capacity, err = strconv.Atoi(items[0])

			if err != nil {
				log.Fatal(err)
			}

			c := fmt.Sprintf("%s %s", items[1], items[2])
			child := lookupTable.getIndex(c)
			a[node][child] = capacity
		}
	}

	count := 0

	startCol := lookupTable.getIndex("shiny gold")
	q := []int{startCol}
	visited := make(map[int]bool)
	for len(q) != 0 {
		col := q[0]
		for k := range a[:][col] {
			if a[k][col] != 0 && !visited[k] {
				q = append(q, k)
				visited[k] = true
				count++
			}
		}
		// "dequeue" first element
		q = q[1:]
	}

	return count
}

func run() []int {
	part1 := 0
	part2 := 0

	return []int{part1, part2}
}
