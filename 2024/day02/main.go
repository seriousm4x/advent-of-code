package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte
var numbers [][]int

func init() {
	lines := strings.Split(string(inputBytes), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		splitted := strings.Split(line, " ")
		var levels []int

		for _, n := range splitted {
			num, err := strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				log.Fatalln(err)
			}
			levels = append(levels, num)
		}

		numbers = append(numbers, levels)
	}
}

func isSafe(level []int) bool {
	isIncreasing := level[0] < level[1]
	var diff int

	for i := 1; i < len(level); i++ {
		if isIncreasing {
			diff = level[i] - level[i-1]
		} else {
			diff = level[i-1] - level[i]
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func solvePart1() {
	res := 0

	for _, level := range numbers {
		if isSafe(level) {
			res++
			continue
		}
	}

	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0

	for _, level := range numbers {
		if isSafe(level) {
			res++
			continue
		}

		for i := 0; i < len(level); i++ {
			temp := make([]int, 0, len(level)-1)
			temp = append(temp, level[:i]...)
			temp = append(temp, level[i+1:]...)
			if isSafe(temp) {
				res++
				break
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
