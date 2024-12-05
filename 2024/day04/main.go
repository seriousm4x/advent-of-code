package main

import (
	_ "embed"
	"log"
	"reflect"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte
var input [][]string
var xmas = []string{"X", "M", "A", "S"}
var invxmas = []string{"S", "A", "M", "X"}

func init() {
	splitted := strings.Split(string(inputBytes), "\n")

	for _, line := range splitted {
		if line == "" {
			continue
		}

		input = append(input, strings.Split(line, ""))

	}
}

func solvePart1() {
	res := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			// Check horizontally order and inverted
			if x+len(xmas) <= len(input[y]) {
				if reflect.DeepEqual(input[y][x:x+len(xmas)], xmas) {
					res++
				} else if reflect.DeepEqual(input[y][x:x+len(xmas)], invxmas) {
					res++
				}
			}
		}
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			// Check vertically order and inverted
			if y+len(xmas) <= len(input) {
				verticalSlice := make([]string, len(xmas))
				for i := 0; i < len(xmas); i++ {
					verticalSlice[i] = input[y+i][x]
				}
				if reflect.DeepEqual(verticalSlice, xmas) {
					res++
				} else if reflect.DeepEqual(verticalSlice, invxmas) {
					res++
				}
			}
		}
	}

	// Check diagonally (top-left to bottom-right) order and inverted
	for y := 0; y < len(input)-len(xmas)+1; y++ {
		for x := 0; x < len(input[y])-len(xmas)+1; x++ {
			diagonalSlice := make([]string, len(xmas))
			for i := 0; i < len(xmas); i++ {
				diagonalSlice[i] = input[y+i][x+i]
			}
			if reflect.DeepEqual(diagonalSlice, xmas) {
				res++
			} else if reflect.DeepEqual(diagonalSlice, invxmas) {
				res++
			}
		}
	}

	// Check diagonally (top-right to bottom-left) order and inverted
	for y := 0; y < len(input)-len(xmas)+1; y++ {
		for x := len(xmas) - 1; x < len(input[y]); x++ {
			diagonalSlice := make([]string, len(xmas))
			for i := 0; i < len(xmas); i++ {
				diagonalSlice[i] = input[y+i][x-i]
			}
			if reflect.DeepEqual(diagonalSlice, xmas) {
				res++
			} else if reflect.DeepEqual(diagonalSlice, invxmas) {
				res++
			}
		}
	}

	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			if input[y][x] == "A" {
				if input[y-1][x-1] == "M" && input[y-1][x+1] == "S" && input[y+1][x+1] == "S" && input[y+1][x-1] == "M" {
					res++
				} else if input[y-1][x-1] == "S" && input[y-1][x+1] == "M" && input[y+1][x+1] == "M" && input[y+1][x-1] == "S" {
					res++
				} else if input[y-1][x-1] == "M" && input[y-1][x+1] == "M" && input[y+1][x+1] == "S" && input[y+1][x-1] == "S" {
					res++
				} else if input[y-1][x-1] == "S" && input[y-1][x+1] == "S" && input[y+1][x+1] == "M" && input[y+1][x-1] == "M" {
					res++
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
