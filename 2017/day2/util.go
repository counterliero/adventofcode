package main

import (
	"strconv"
	"os"
	"fmt"
	"bufio"
)

func Min(all...int) int {
	min := 0xffffffff

	for _, num := range all {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(all...int) (max int) {
	for _, num := range all {
		if num > max {
			max = num
		}
	}
	return max
}

func MinMax(all...int) (int, int) {
	return Min(all...), Max(all...)
}

func GetIntegers(values []string) (numbers []int) {
	for _, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Fprintf(os.Stderr,"ERROR: %s", err)
			os.Exit(1)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func ReadLinesFromFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return lines
}

func Divide(nums []int) (int) {
	for _, a := range nums {
		for _, b := range nums {
			if a == b {
				continue
			}
			if a/b > 0 && a%b == 0 {
				return a/b
			}
		}
	}
	panic("That did not work")
}
