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

	if len(lines) == 0 {
		log.Println("No input provided")
		return
	}

	for ranges := range strings.SplitSeq(lines[0], ",") {
		ids := strings.Split(ranges, "-")

		firstID, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatal("Error parsing first ID:", err)
		}

		lastID, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatal("Error parsing last ID:", err)
		}

		for id := firstID; id <= lastID; id++ {
			s := strconv.Itoa(id)
			l := len(s)
			if s[:l/2] == s[l/2:l] {
				res += id
			}
		}
	}
	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0

	ranges := strings.SplitSeq(lines[0], ",")
	for r := range ranges {
		ids := strings.Split(r, "-")

		firstID, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatal("Error parsing first ID:", err)
		}

		lastID, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatal("Error parsing last ID:", err)
		}

		for id := firstID; id <= lastID; id++ {
			s := strconv.FormatInt(int64(id), 10)
			l := len(s)

			invalid := false

			for size := 1; size*2 <= l && !invalid; size++ {
				if l%size != 0 {
					continue
				}

				chunk := s[:size]
				ok := true

				for i := size; i < l; i += size {
					if s[i:i+size] != chunk {
						ok = false
						break
					}
				}

				if ok {
					invalid = true
				}
			}

			if invalid {
				res += id
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
