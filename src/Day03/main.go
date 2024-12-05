package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/phgeorgiev/aoc-2024-go/pkg/files"
)

func main() {
	fmt.Printf("Part 1: %d\n", MultiplyResults("input.txt"))
	fmt.Printf("Part 2 %d\n", MultiplyResultsWithAllInstructions("input.txt"))
}

func MultiplyResults(filePath string) int {
	var result int

	lines := files.ReadFile(filePath)
	for _, line := range lines {
		var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			result += num1 * num2
		}
	}

	return result
}

func MultiplyResultsWithAllInstructions(filePath string) int {
	var result int

	lines := files.ReadFile(filePath)
	enable := true
	for _, line := range lines {
		var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			switch match[0] {
			case "do()":
				enable = true
			case "don't()":
				enable = false
			default:
				{
					if enable {
						num1, _ := strconv.Atoi(match[1])
						num2, _ := strconv.Atoi(match[2])

						result += num1 * num2
					}
				}
			}
		}
	}

	return result
}
