package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckSumCalcsStuff(t *testing.T) {
	input := ReadLinesFromFile("input_test.txt")
	assert.Equal(t, 136, CheckSum(input))
}

func TestCheckSumDivideCalcsStuff(t *testing.T) {
	input := ReadLinesFromFile("input_test_div.txt")
	assert.Equal(t, 9, CheckSumDivide(input))
}
