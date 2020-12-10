package day8

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int
}

type state struct {
	acc int
	sp  int
}

func readInput(filename string) []instruction {

	file, err := os.Open(path.Join("testdata", filename))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	program := []instruction{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, " ")
		op := fields[0]
		arg, err := strconv.Atoi(fields[1])

		if err != nil {
			continue
		}

		program = append(program, instruction{op, arg})
	}

	return program

}

func execOne(p []instruction) int {
	acc := 0
	sp := 0
	visited := make(map[int]bool)

	for !visited[sp] && sp < len(p) {
		curr := p[sp]
		arg := curr.arg
		visited[sp] = true
		switch curr.op {
		case "acc":
			acc += arg
			sp++
		case "jmp":
			sp += arg
		case "nop":
			sp++
		}
	}

	return acc
}

func execTwo(acc int, sp int, p []instruction) []state {
	visited := make(map[int]bool)
	stk := []state{}
	curr := instruction{}

	for !visited[sp] && sp != len(p) {
		curr = p[sp]

		arg := curr.arg

		visited[sp] = true
		switch curr.op {
		case "acc":
			acc += arg
			sp++
		case "jmp":
			sp += arg
		case "nop":
			sp++
		}
		stk = append(stk, state{acc, sp})
	}

	return stk
}

func part2(prog []instruction) int {
	sp := 0
	acc := 0
	st := execTwo(acc, sp, prog)
	s := state{}
	for i := 1; sp != len(prog) && i <= len(st); i++ {
		s = st[len(st)-i]

		newProg := make([]instruction, len(prog))
		copy(newProg, prog)

		lastSP := s.sp
		lastOp := prog[lastSP].op

		switch lastOp {
		case "jmp":
			newProg[lastSP].op = "nop"
		case "nop":
			newProg[lastSP].op = "jmp"
		}
		newSt := execTwo(s.acc, s.sp, newProg)
		sp = newSt[len(newSt)-1].sp
		acc = newSt[len(newSt)-1].acc

	}

	return acc
}

func run(filename string) (int, int) {
	prog := readInput(filename)
	p1 := execOne(prog)
	p2 := part2(prog)

	return p1, p2
}
