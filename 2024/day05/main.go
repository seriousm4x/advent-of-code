package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputBytes []byte

var orderingRules []string
var pageNums []string

func init() {
	splitted := strings.Split(string(inputBytes), "\n\n")
	orderingRules = strings.Split(splitted[0], "\n")
	pageNums = strings.Split(strings.TrimSpace(splitted[1]), "\n")
}

func solvePart1() {
	res := 0

	for _, page := range pageNums {
		ok := true
		for _, rule := range orderingRules {
			nums := strings.Split(rule, "|")

			if !strings.Contains(page, nums[0]) || !strings.Contains(page, nums[1]) {
				continue
			}

			pattern := regexp.MustCompile(fmt.Sprintf("%s,.*?%s(,|$)", nums[0], nums[1]))
			if !pattern.MatchString(page) {
				ok = false
				break
			}
		}
		if ok {
			splitted := strings.Split(page, ",")
			half := len(splitted) / 2
			num, err := strconv.Atoi(splitted[half])
			if err != nil {
				log.Fatalln(err)
			}
			res += num
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
