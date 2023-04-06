package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	app := NewApp()

	inputFile, _ := os.Open("./text.txt")
	app.source = inputFile

	expectedOutput := []string{"a\r", "a\r", "a\r", "b\r", "c\r", "n\r", "s", "v\r", "z\r"}

	t.Run("Test 1", func(t *testing.T) {
		actual, err := app.Scan()
		assert.NoError(t, err)
		assert.Equal(t, expectedOutput, actual)
	})
}
