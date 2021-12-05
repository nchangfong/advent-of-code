package day14

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	op   string
	addr int
	arg  string
}

func readInput(reader io.Reader) []instruction {

	scanner := bufio.NewScanner(reader)
	res := []instruction{}
	for scanner.Scan() {
		var ins instruction
		line := scanner.Text()
		fields := strings.Split(line, " = ")
		if fields[0] == "mask" {
			ins.op = fields[0]
			ins.addr = -1
			ins.arg = fields[1]
		} else {
			subFields := strings.Split(fields[0], "[")
			ins.op = subFields[0]

			mem, err := strconv.Atoi(subFields[1][:len(subFields[1])-1])
			if err != nil {
				log.Fatal(err)
			}
			ins.addr = mem
			ins.arg = fields[1]
		}
		res = append(res, ins)
	}

	return res
}

func part1(ins []instruction) {
	// mem := make(map[int]int)

	for _, i := range ins {
		switch i.op {
		case "mask":

		}
	}
	// strconv.FormatInt(n, 2)
}
