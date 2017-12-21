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

func TestDivideDoesThings(t *testing.T) {
	input := []int{5, 9, 2, 8}
	assert.Equal(t, 4, Divide(input))
}

func TestDivideDeterminesWhenThereIsNoResult(t *testing.T) {
	input := []int{5, 4}
	assert.Equal(t, -1, Divide(input))
}
