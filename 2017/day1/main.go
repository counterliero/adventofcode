package main

import (
	"strconv"
	"fmt"
	"os"
)

func CaptchaSum(input string, step int) (int, error) {
	sum := 0

	_, err := strconv.Atoi(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: %s is not a digit", input)
		return sum, err
	}

	for i := range input {
		if input[i] == input[(i+1)%len(input)] {
			sum += int(input[i] - '0')
		}
	}

	return sum, nil
}

func main() {
	longListofNumers:= "2248558274598352939674972954641755895841499719997985122794292687271422835368955117812935354833177978374291536134322914153833468825481971481366443"

	fmt.Println(CaptchaSum(longListofNumers, 1))

	fmt.Println(CaptchaSum(longListofNumers, len(longListofNumers)/2))
}
