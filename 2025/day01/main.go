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
var lines = strings.Split(string(inputBytes), "\n")

func solvePart1() {
	res := 0
	cur := 50

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		num %= 100

		switch direction {
		case "R":
			cur += num
		case "L":
			cur -= num
		default:
			log.Fatalf("Unknown rotation: %s", direction)
		}

		cur %= 100

		if cur == 0 {
			res += 1
		}
	}

	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0
	cur := 50

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		first := 0
		switch direction {
		case "R":
			first = (100 - cur) % 100
		case "L":
			first = cur % 100
		default:
			log.Fatalf("Unknown rotation: %s", direction)
		}

		if first == 0 {
			first = 100
		}

		if first <= num {
			res += 1 + (num-first)/100
		}

		switch direction {
		case "R":
			cur = (cur + num) % 100
		default:
			cur = (cur - num) % 100
			if cur < 0 {
				cur += 100
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
