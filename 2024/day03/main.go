package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte
var lines []string

func init() {
	splitted := strings.Split(string(inputBytes), "\n")

	for _, line := range splitted {
		if line == "" {
			continue
		}

		lines = append(lines, line)
	}
}

func multiplyMatches(str1 string, str2 string) int {
	num1, err := strconv.Atoi(str1)
	if err != nil {
		log.Fatalln(err)
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		log.Fatalln(err)
	}

	return num1 * num2

}

func solvePart1() {
	res := 0
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)

		lineRes := 0

		for _, match := range matches {
			res += multiplyMatches(match[1], match[2])
		}

		res += lineRes
	}

	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0
	ok := true
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				ok = true
			} else if match[0] == "don't()" {
				ok = false
			} else {
				if ok {
					res += multiplyMatches(match[1], match[2])
				}
			}
		}
	}

	log.Println("Result of part 2 is:", res)
}

func main() {
	start := time.Now()
	solvePart1()
	solvePart2()
	log.Println("Operation took:", time.Since(start))
}
