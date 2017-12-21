package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckSumCalcsStuff(t *testing.T) {
	input := ReadLinesFromFile("input_test.txt")
	assert.Equal(t, 136, CheckSum(input))
}
