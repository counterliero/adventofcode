package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinFindsMinimum(t *testing.T) {
	assert.Equal(t, 1, Min(7, 16, 3, 28, 55, 1, 11, 22))
}

func TestMaxFindsMaximum(t *testing.T) {
	assert.Equal(t, 55, Max(7, 16, 3, 28, 55, 1, 11, 22))
}

func TestMinMaxGetsThemBoth(t *testing.T) {
	min, max := MinMax(7, 16, 3, 28, 55, 1, 11, 22)
	assert.Equal(t, 1, min)
	assert.Equal(t, 55, max)
}

func TestGetIntegersGetsValuesFromStringList(t *testing.T) {
	input := []string{"1", "2", "11", "9", "17"}
	expected := []int{1, 2, 11, 9, 17}
	assert.Equal(t, expected, GetIntegers(input))
}

func TestDivide(t *testing.T) {
	input := ReadLinesFromFile("input_test_div.txt")
	assert.Equal(t, 9, CheckSumDivide(input))
}
