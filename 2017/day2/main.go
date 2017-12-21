package main

import (
	"strings"
	"fmt"
	"os"
)

func CheckSum(nums []string) int {
	checksum := 0
	for _, line := range nums {
		integers := GetIntegers(strings.Fields(line))
		min, max := MinMax(integers...)
		diff := max - min
		checksum += diff
	}
	return checksum
}

func CheckSumDivide(lines []string) int {
	checksum := 0
	for _, line := range lines {
		integers := GetIntegers(strings.Fields(line))
		checksum += Divide(integers)
	}
	return checksum
}

func main() {
	input := ReadLinesFromFile("input.txt")

	// Part 1
	fmt.Println(CheckSum(input))

	// Part 2
	res := CheckSumDivide(input)
	if res == -1 {
		fmt.Fprintln(os.Stderr, "ERROR: input is invalid")
		os.Exit(1)
	}
	fmt.Println(res)
}
