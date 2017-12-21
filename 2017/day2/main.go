package main

import (
	"strings"
	"fmt"
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
	fmt.Println(CheckSumDivide(input))
}
