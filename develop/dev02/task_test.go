package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	// - "a4bc2d5e" => "aaaabccddddde"
	// - "abcd" => "abcd"
	// - "45" => "" (некорректная строка)
	// - "" => ""

	var testsOk = []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
	}

	var testsWrong = []struct {
		input    string
		expected string
	}{
		{"45", ""},
	}

	for i, x := range testsOk {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			actual, err := Unpack(x.input)
			assert.NoError(t, err)
			assert.Equal(t, x.expected, actual)
		})
	}

	for i, x := range testsWrong {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			actual, err := Unpack(x.input)
			assert.Error(t, err)
			assert.Equal(t, x.expected, actual)
		})
	}
}
