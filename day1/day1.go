package day1

import (
	"bufio"
	"fmt"
	"os"
)

func Day1() {
	increase := getIncreaseCount("day1/input.txt")
	fmt.Println("There were ", increase, " increases")
}

func parseLine(line string) int {
	var res int
	_, err := fmt.Sscanf(line, "%d", &res)
	if err != nil {
		panic(err)
	}
	return res
}

func parseFile(filename string) []int {
	var res = []int{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, parseLine(scanner.Text()))
	}
	return res
}

func isPrevLessThan(prev, cur int) bool {
	return prev < cur
}

func getIncreaseCount(filename string) (increase int) {
	res := parseFile(filename)
	for i := range res {
		if i > 0 {
			if isPrevLessThan(res[i-1], res[i]) {
				increase += 1
			}
		}
	}
	return increase
}
