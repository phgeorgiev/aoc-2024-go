package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		expected := 11
		actual := TotalDistance("test_input.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		expected := 31
		actual := SimilarityScore("test_input.txt")
		assert.Equal(t, expected, actual)
	})

}
