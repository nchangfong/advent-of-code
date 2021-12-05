package day4

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == ':' || r == ' '
}

// processPairs populates the map with the key value pairs
func processPairs(m map[string]string, p []string) {
	for i := 0; i < len(p); i += 2 {
		key := p[i]
		value := p[i+1]
		m[key] = value
	}
}

func validatePassport(m map[string]string) bool {
	var required = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, v := range required {
		if m[v] == "" {
			return false
		}
	}
	return true
}

func validateYear(min int, max int, y string) bool {
	year, err := strconv.Atoi(y)
	if err != nil || year < min || year > max {
		return false
	}
	return true
}

func validateHeight(h string) bool {
	s := []rune(h)
	if len(s) < 2 {
		return false
	}
	unit := string(s[len(s)-2:])
	value, err := strconv.Atoi(string(s[:len(s)-2]))
	if err != nil {
		return false
	}

	switch unit {
	case "cm":
		if value >= 150 && value <= 193 {
			return true
		}
	case "in":
		if value >= 59 && value <= 76 {
			return true
		}
	}

	return false
}

func validateColour(s string) bool {
	if len(s) == 0 {
		return false
	}
	if _, err := hex.DecodeString(string(s[1:])); err != nil || len(s) != 7 {
		return false
	}

	return true
}

func validatePID(value string) bool {
	set := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	return set[value]
}

func validateECL(value string) bool {
	if _, err := strconv.Atoi(value); err != nil || len(value) != 9 {
		return false
	}
	return true
}

func validatePassport2(m map[string]string) bool {
	var required = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	if len(m) == 0 || len(m) < 7 {
		return false
	}

	for _, v := range required {

		res := false
		value := m[v]
		switch v {
		case "byr":
			res = validateYear(1920, 2002, value)
		case "iyr":
			res = validateYear(2010, 2020, value)
		case "eyr":
			res = validateYear(2020, 2030, value)
		case "hgt":
			res = validateHeight(value)
		case "hcl":
			res = validateColour(value)
		case "ecl":
			res = validatePID(value)
		case "pid":
			res = validateECL(value)
		}
		if !res {
			return false
		}

	}
	return true
}

func run() []int {
	file, err := os.Open(filepath.Join("testdata", "input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1 := 0
	part2 := 0
	passport := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			pairs := strings.FieldsFunc(line, split)
			processPairs(passport, pairs)
		} else {
			if res := validatePassport(passport); res {
				part1++
			}

			if res := validatePassport2(passport); res {
				part2++
			}

			passport = make(map[string]string)
		}

	}

	return []int{part1, part2}
}
