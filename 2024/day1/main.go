package main

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte
var listLeft []int
var listRight []int

func init() {
	lines := strings.Split(string(inputBytes), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		splitted := strings.Split(line, "   ")

		left, err := strconv.Atoi(strings.TrimSpace(splitted[0]))
		if err != nil {
			log.Fatalln(err)
		}

		right, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
		if err != nil {
			log.Fatalln(err)
		}

		listLeft = append(listLeft, left)
		listRight = append(listRight, right)
	}
}

func solvePart1() {
	slices.Sort(listLeft)
	slices.Sort(listRight)

	res := 0

	for i := 0; i < len(listLeft); i++ {
		diff := listLeft[i] - listRight[i]
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}

	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0
	counts := make(map[int]int, len(listRight))

	for _, numR := range listLeft {
		counts[numR]++
	}

	for _, numL := range listLeft {
		if count, exists := counts[numL]; exists {
			res += numL * count
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
