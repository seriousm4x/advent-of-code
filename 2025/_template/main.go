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

func solvePart1() {
	res := 0
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
