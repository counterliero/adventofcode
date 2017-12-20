package main

import (
	"strconv"
	"fmt"
)

func CaptchaSum(input string, step int) (int, error) {
	sum := 0

	for i := range input {
		if input[i] == input[(i+step)%len(input)] {
			_, err := strconv.Atoi(string(input[i]))
			if err != nil {
				return sum, err
			}

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
