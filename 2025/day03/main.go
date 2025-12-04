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
var lines = strings.Split(string(inputBytes), "\n")

func sliceToSum(s []int) int {
	sumStr := ""
	for i := range s {
		sumStr += strconv.Itoa(s[i])
	}
	sum, err := strconv.Atoi(sumStr)
	if err != nil {
		log.Fatal(err)
	}
	return sum
}

func solvePart1() {
	res := 0

	for _, bank := range lines {
		if bank == "" {
			continue
		}
		batteriesInBank := strings.Split(bank, "")
		var highestBattery int
		var secondHighestBattery int
		for i, batteryStr := range batteriesInBank {
			batteryInt, err := strconv.Atoi(batteryStr)
			if err != nil {
				log.Fatal(err)
			}
			if batteryInt > highestBattery && i+1 != len(batteriesInBank) {
				highestBattery = batteryInt
				secondHighestBattery = 0
			} else if batteryInt > secondHighestBattery {
				secondHighestBattery = batteryInt
			}
		}
		res += sliceToSum([]int{highestBattery, secondHighestBattery})
	}
	log.Println("Result of part 1 is:", res)
}

func solvePart2() {
	res := 0

	for _, bank := range lines {
		if bank == "" {
			continue
		}
		batteriesInBank := strings.Split(bank, "")
		highestBatteries := make([]int, 12)
		for _, joltageStr := range batteriesInBank {
			joltageInt, err := strconv.Atoi(joltageStr)
			if err != nil {
				log.Fatal(err)
			}

			highestCombination := highestBatteries
			for i := range highestBatteries {
				tryOurNum := slices.Delete(slices.Clone(highestBatteries), i, i+1)
				tryOurNum = append(tryOurNum, joltageInt)
				// log.Println("max:", highestCombination, "try:", tryOurNum, "current num:", joltageInt, "del index:", i)
				if sliceToSum(tryOurNum) > sliceToSum(highestCombination) {
					// log.Println("higher")
					highestCombination = tryOurNum
				}
			}
			if sliceToSum(highestCombination) > sliceToSum(highestBatteries) {
				highestBatteries = highestCombination
			}
		}
		// log.Println("highest number is:", highestBatteries)
		res += sliceToSum(highestBatteries)
	}

	log.Println("Result of part 2 is:", res)
}

func main() {
	start := time.Now()
	solvePart1()
	solvePart2()
	log.Println("Operation took:", time.Since(start))
}
