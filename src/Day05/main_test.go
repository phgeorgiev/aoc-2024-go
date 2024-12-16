package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		expected := 143
		actual := MiddlePageNumbersSum("test_input.txt")
		assert.Equal(t, expected, actual)
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		expected := 123
		actual := MiddlePageNumbersSumOfIncorectOrderedUpdates("test_input.txt")
		assert.Equal(t, expected, actual)
	})
}
