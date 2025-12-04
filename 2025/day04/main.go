package main

import (
	_ "embed"
	"log"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte
var lines = strings.Split(string(inputBytes), "\n")

func rollIsAccessible(lines []string, lineIndex, charIndex int) bool {
	if lines[lineIndex][charIndex] != '@' {
		return false
	}

	neighbors := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {1, -1}, {-1, 1}, {1, 1},
	}

	rollsNearby := 0

	for _, direction := range neighbors {
		directionX, directionY := direction[0], direction[1]

		neighborY := lineIndex + directionY
		neighborX := charIndex + directionX

		if neighborY < 0 || neighborY >= len(lines) {
			continue
		}
		if neighborX < 0 || neighborX >= len(lines[neighborY]) {
			continue
		}

		if lines[neighborY][neighborX] == '@' {
			rollsNearby++
		}
	}

	return rollsNearby < 4
}

func solvePart1() {
	res := 0
	for lineIndex, line := range lines {
		if line == "" {
			continue
		}
		for charIndex := range line {
			if line[charIndex] != '@' {
				continue
			}
			if rollIsAccessible(lines, lineIndex, charIndex) {
				res++
			}
		}
	}
	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0
	log.Println("Result of part 2 is:", res)
}

func main() {
	start := time.Now()
	solvePart1()
	solvePart2()
	log.Println("Operation took:", time.Since(start))
}
