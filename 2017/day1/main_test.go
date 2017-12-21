package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapchaSumValidatesFirstExample(t *testing.T) {
	result, _ := CaptchaSum("1122", 1)
	assert.Equal(t, 3, result)
}

func TestCapchaSumValidatesSecondExample(t *testing.T) {
	result, _ := CaptchaSum("1111", 1)
	assert.Equal(t, 4, result)
}

func TestCapchaSumValidatesThirdExample(t *testing.T) {
	result, _ := CaptchaSum("1234", 1)
	assert.Equal(t, 0, result)
}

func TestCapchaSumValidatesLastExample(t *testing.T) {
	result, _ := CaptchaSum("91212129", 1)
	assert.Equal(t, 9, result)
}

func TestCapchaSumValidatesSecondPart(t *testing.T) {
	input := "91212129"
	result, _ := CaptchaSum(input, len(input))
	assert.Equal(t, 27, result)
}

func TestCaptchaSumRaisesErrorWhenInputIsNotADigit(t *testing.T) {
	_, err := CaptchaSum("123foobar456", 1)
	assert.Error(t, err)
}
