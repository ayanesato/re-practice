package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isLeapYear(t *testing.T) {
	test1 := isLeapYear(2023)
	test2 := isLeapYear(2024)
	test3 := isLeapYear(2025)
	test4 := isLeapYear(1999)
	test5 := isLeapYear(2000)
	test6 := isLeapYear(2001)
	test7 := isLeapYear(1899)
	test8 := isLeapYear(1900)
	test9 := isLeapYear(1901)
	assert.Equal(t, false, test1)
	assert.Equal(t, true, test2)
	assert.Equal(t, false, test3)
	assert.Equal(t, false, test4)
	assert.Equal(t, true, test5)
	assert.Equal(t, false, test6)
	assert.Equal(t, false, test7)
	assert.Equal(t, false, test8)
	assert.Equal(t, false, test9)
}
