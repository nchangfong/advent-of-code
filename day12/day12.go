package day12

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
)

type command struct {
	cmd string
	arg int
}

func readInput(filename string) []command {
	file, err := os.Open(path.Join("testdata", filename))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	res := []command{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cmd := line[:1]
		arg, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, command{cmd, arg})
	}

	return res
}

func positiveMod(a, b int) int {
	return (a%b + b) % b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func rotate(card string, dir string, deg int) string {
	newDeg := 0
	cardinalRef := map[string]int{
		"N": 0,
		"E": 90,
		"S": 180,
		"W": 270,
	}

	degRef := map[int]string{
		0:   "N",
		90:  "E",
		180: "S",
		270: "W",
	}

	if dir == "R" {
		newDeg = positiveMod(cardinalRef[card]+deg, 360)
	} else if dir == "L" {
		newDeg = positiveMod(cardinalRef[card]-deg, 360)
	}

	return degRef[newDeg]
}

func move(dir string, arg int, x int, y int) (int, int) {
	switch dir {
	case "N":
		y += arg
	case "S":
		y -= arg
	case "E":
		x += arg
	case "W":
		x -= arg
	}

	return x, y
}

func part1(cmds []command) int {
	dir := "E"
	x := 0
	y := 0

	for _, v := range cmds {
		cmd := v.cmd
		arg := v.arg
		switch cmd {
		case "N":
			y += arg
		case "S":
			y -= arg
		case "E":
			x += arg
		case "W":
			x -= arg
		case "F":
			x, y = move(dir, arg, x, y)
		default:
			dir = rotate(dir, cmd, arg)
		}
	}

	return abs(x) + abs(y)
}

func rotate2(x1, y1, x2, y2, deg int) (int, int) {

	rx := x1 - x2
	ry := y1 - y2

	deg = positiveMod(deg, 360)

	for i := 0; i < deg/90; i++ {
		tmp := rx
		rx = ry
		ry = -tmp
	}

	return rx + x2, ry + y2
}

func part2(cmds []command) int {

	x := 0
	y := 0

	wx := 10
	wy := 1

	for _, v := range cmds {
		cmd := v.cmd
		arg := v.arg
		switch cmd {
		case "N":
			wy += arg
		case "S":
			wy -= arg
		case "E":
			wx += arg
		case "W":
			wx -= arg
		case "F":
			dx := wx - x
			dy := wy - y

			x += dx * arg
			y += dy * arg

			wx = x + dx
			wy = y + dy
		case "L":
			wx, wy = rotate2(wx, wy, x, y, -arg)
		case "R":
			wx, wy = rotate2(wx, wy, x, y, arg)
		}
	}

	return abs(x) + abs(y)
}
