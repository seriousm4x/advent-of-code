package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte
var database = strings.Split(string(inputBytes), "\n\n")
var freshRanges [][]int

func init() {
	// parse fresh ranges
	for line := range strings.SplitSeq(database[0], "\n") {
		if line == "" {
			continue
		}
		lineSplitted := strings.Split(line, "-")
		startString := lineSplitted[0]
		endString := lineSplitted[1]

		start, err := strconv.Atoi(startString)
		if err != nil {
			log.Fatalln("Error converting startString to int:", err)
		}
		end, err := strconv.Atoi(endString)
		if err != nil {
			log.Fatalln("Error converting endString to int:", err)
		}

		freshRanges = append(freshRanges, []int{start, end})
	}
}

func solvePart1() {
	res := 0

	// check ingredient ID in fresh ranges
	for line := range strings.SplitSeq(database[1], "\n") {
		if line == "" {
			continue
		}
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("Error converting idString to int:", err)
		}

		for i := range freshRanges {
			if id >= freshRanges[i][0] && id <= freshRanges[i][1] {
				res++
				break
			}
		}
	}

	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0

	sortedRanges := make([][]int, len(freshRanges))
	copy(sortedRanges, freshRanges)

	sort.Slice(sortedRanges, func(i, j int) bool {
		return sortedRanges[i][0] < sortedRanges[j][0]
	})

	mergedFreshRanges := [][]int{}
	currentRange := sortedRanges[0]

	for i := 1; i < len(sortedRanges); i++ {
		nextRange := sortedRanges[i]
		if nextRange[0] <= currentRange[1]+1 {
			if nextRange[1] > currentRange[1] {
				currentRange[1] = nextRange[1]
			}
		} else {
			mergedFreshRanges = append(mergedFreshRanges, currentRange)
			currentRange = nextRange
		}
	}

	mergedFreshRanges = append(mergedFreshRanges, currentRange)

	for _, r := range mergedFreshRanges {
		res += r[1] - r[0] + 1
	}

	log.Println("Result of part 2 is:", res)
}

func main() {
	start := time.Now()
	solvePart1()
	solvePart2()
	log.Println("Operation took:", time.Since(start))
}
