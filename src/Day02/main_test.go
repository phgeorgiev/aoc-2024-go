package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		expected := 2
		actual := CountSafeReports("test_input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		expected := 4
		actual := CountSafeReportsWithRemoval("test_input.txt")
		assert.Equal(t, expected, actual)
	})
}
