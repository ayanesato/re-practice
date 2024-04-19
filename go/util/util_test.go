package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isLeapYear_assert(t *testing.T) {
	assert.False(t, false, isLeapYear(2023))
	assert.True(t, true, isLeapYear(2024))
	assert.False(t, false, isLeapYear(2025))
	assert.False(t, false, isLeapYear(1999))
	assert.True(t, true, isLeapYear(2000))
	assert.False(t, false, isLeapYear(2001))
	assert.False(t, false, isLeapYear(1899))
	assert.False(t, false, isLeapYear(1900))
	assert.False(t, false, isLeapYear(1901))
}

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{year: 2023}, false},
		{"test2", args{year: 2024}, true},
		{"test3", args{year: 2025}, false},
		{"test4", args{year: 1899}, false},
		{"test5", args{year: 1900}, false},
		{"test6", args{year: 1901}, false},
		{"test7", args{year: 1999}, false},
		{"test8", args{year: 2000}, true},
		{"test9", args{year: 2001}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
