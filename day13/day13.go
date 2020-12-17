package day13

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"
)

func readInput(filename string) []int {
	file, err := os.Open(path.Join("testdata", filename))

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")

	res := []int{time}

	for _, i := range buses {
		val, err := strconv.Atoi(i)
		if err != nil {
			val = -1
		}
		res = append(res, val)
	}

	return res
}

func part1(s []int) int {
	target := s[0]
	buses := s[1:]

	min := 0
	minIdx := 0

	for k, v := range buses {
		if v == -1 {
			continue
		}

		arrival := target + v - (target % v)

		if k == 0 {
			min = arrival
			minIdx = k
		} else if arrival < min {
			min = arrival
			minIdx = k
		}

	}

	return (min - target) * buses[minIdx]
}

// cong represents a congruence where x ≡ a mod n.
// since x is unknown, it's not included in the struct
type cong struct {
	a *big.Int
	n *big.Int
}

func positiveMod(a, b int) int {
	return (a%b + b) % b
}

func part2(s []int) big.Int {

	buses := s[1:]

	// system of congruences
	pairs := []cong{}
	prod := new(big.Int)
	prod.SetInt64(1)

	for k, v := range buses {
		if v == -1 {
			continue
		}
		bigK := new(big.Int).SetInt64(int64(k))
		n := new(big.Int).SetInt64(int64(v))
		// x + dt ≡ 0 mod n
		// in terms of x, x ≡ -dt mod n
		a := new(big.Int)
		a.Mod(new(big.Int).Neg(bigK), n)

		p := cong{a, n}
		pairs = append(pairs, p)
		prod.Mul(prod, n)
	}

	sum := new(big.Int).SetInt64(0)

	// solve system using general solution for chinese remainder theorem
	// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Existence_(direct_construction)
	for _, p := range pairs {
		q := new(big.Int)
		inv := new(big.Int)
		q.Div(prod, p.n)
		inv.ModInverse(q, p.n)
		q.Mul(q, p.a).Mul(q, inv)

		sum.Add(sum, q)
	}

	return *sum.Mod(sum, prod)
}
