package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/phgeorgiev/aoc-2024-go/pkg/array"
	"github.com/phgeorgiev/aoc-2024-go/pkg/files"
)

func main() {
	fmt.Printf("Part 1: %d\n", CountSafeReports("input.txt"))
	fmt.Printf("Part 2: %d\n", CountSafeReportsWithRemoval("input.txt"))
}

func CountSafeReports(filePath string) int {
	result := 0
	levels := getLevels(filePath)
	for _, values := range levels {
		if isReportSafe(values) {
			result++
		}
	}

	return result
}

func CountSafeReportsWithRemoval(filePath string) int {
	result := 0
	levels := getLevels(filePath)
	for _, values := range levels {
		safe := false
		for i := range values {
			safe = isReportSafe(remove(values, i))
			if safe {
				break
			}
		}
		if safe {
			result++
		}
	}

	return result
}

func isReportSafe(levels []int) bool {
	differences := zipWithNext(levels)
	flagIncrease, flagDecrease := false, false
	allowedDifferLevels := []int{-3, -2, -1, 1, 2, 3}
	for _, value := range differences {
		if !array.InArray(value, allowedDifferLevels) {
			return false
		}

		if value > 0 {
			flagIncrease = true
		} else {
			flagDecrease = true
		}

		if flagDecrease && flagIncrease {
			return false
		}
	}

	return true
}

func zipWithNext(levels []int) []int {
	var result []int
	i := 0
	for i < len(levels)-1 {
		result = append(result, levels[i]-levels[i+1])
		i++
	}

	return result
}

func remove(slice []int, s int) []int {
	return slices.Concat(slice[:s], slice[s+1:])
}

func getLevels(filePath string) [][]int {
	var result [][]int
	lines := files.ReadFile(filePath)
	for _, element := range lines {
		splitLine := strings.Split(element, " ")

		var levels []int
		for _, c := range splitLine {
			value, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			levels = append(levels, value)
		}

		result = append(result, levels)
	}

	return result
}
