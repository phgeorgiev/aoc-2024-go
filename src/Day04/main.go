package main

import (
	"fmt"

	"github.com/phgeorgiev/aoc-2024-go/pkg/files"
)

func main() {
	fmt.Printf("Part 1: %d\n", CountWords("input.txt"))
	fmt.Printf("Part 2 %d\n", CountXShapeWords("input.txt"))
}

func CountWords(filePath string) int {
	var result int
	lines := files.ReadFile(filePath)
	for i := 0; i < len(lines[0]); i++ {
		for j := 0; j < len(lines); j++ {
			if lines[i][j] == 'X' {
				result += findWords(complex(float64(i), float64(j)), lines)
			}
		}
	}

	return result
}

func CountXShapeWords(filePath string) int {
	var result int
	lines := files.ReadFile(filePath)
	for i := 0; i < len(lines[0]); i++ {
		for j := 0; j < len(lines); j++ {
			if lines[i][j] == 'A' && findXShapeWords(complex(float64(i), float64(j)), lines) {
				result++
			}
		}
	}

	return result
}

func findWords(initialPosition complex128, grid []string) int {
	foundWords := 0
	directions := []complex128{
		complex(0, 1),
		complex(1, 1),
		complex(1, 0),
		complex(1, -1),
		complex(0, -1),
		complex(-1, -1),
		complex(-1, 0),
		complex(-1, 1),
	}
	searchTerm := "MAS"
	for _, direction := range directions {
		position := initialPosition
		for i := range searchTerm {
			position += direction
			x := int(real(position))
			y := int(imag(position))

			isOutsideTheGrid := 0 > x || x >= len(grid[0]) || 0 > y || y >= len(grid)
			if isOutsideTheGrid {
				break
			}

			value := grid[x][y]
			if value != searchTerm[i] {
				break
			}
			if i == len(searchTerm)-1 {
				foundWords++
			}
		}
	}

	return foundWords
}

func findXShapeWords(initialPosition complex128, grid []string) bool {
	foundWords := 0
	directions := []complex128{
		complex(1, 1),
		complex(1, -1),
		complex(-1, -1),
		complex(-1, 1),
	}
	searchTerm := "MS"
	for _, direction := range directions {
		position := initialPosition + direction
		for i := range searchTerm {
			x := int(real(position))
			y := int(imag(position))

			isOutsideTheGrid := 0 > x || x >= len(grid[0]) || 0 > y || y >= len(grid)
			if isOutsideTheGrid {
				break
			}

			value := grid[x][y]
			if value != searchTerm[i] {
				break
			}

			if i == len(searchTerm)-1 {
				foundWords++
			}

			position = initialPosition + (direction * -1)
		}
	}

	return foundWords == 2
}
