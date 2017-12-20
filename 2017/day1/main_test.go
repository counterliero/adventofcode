package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapchaSumValidatesFirstExample(t *testing.T) {
	result, _ := CaptchaSum("1122", 1)
	assert.Equal(t, result, 3)
}

func TestCapchaSumValidatesSecondExample(t *testing.T) {
	result, _ := CaptchaSum("1111", 1)
	assert.Equal(t, result, 4)
}

func TestCapchaSumValidatesThirdExample(t *testing.T) {
	result, _ := CaptchaSum("1234", 1)
	assert.Equal(t, result, 0)
}

func TestCapchaSumValidatesLastExample(t *testing.T) {
	result, _ := CaptchaSum("91212129", 1)
	assert.Equal(t, result, 9)
}

func TestCapchaSumValidatesSecondPart(t *testing.T) {
	input := "91212129"
	result, _ := CaptchaSum(input, len(input))
	assert.Equal(t, result, 27)
}

func TestCaptchaSumRaisesErrorWhenInputIsNotADigit(t *testing.T) {
	_, err := CaptchaSum("123foobar456", 1)
	assert.Error(t, err)
}
