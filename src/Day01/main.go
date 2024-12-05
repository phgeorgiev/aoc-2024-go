package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/phgeorgiev/aoc-2024-go/pkg/files"
)

func main() {
	fmt.Printf("Part 1: %d\n", TotalDistance("input.txt"))
	fmt.Printf("Part 2 %d\n", SimilarityScore("input.txt"))
}

func TotalDistance(filePath string) int {
	left, right := parseList(filePath)
	slices.Sort(left)
	slices.Sort(right)

	var result int
	for i := range left {
		result += int(math.Abs(float64(left[i] - right[i])))
	}

	return result
}

func SimilarityScore(filePath string) int {
	left, right := parseList(filePath)

	leftListSet := createMap(left)
	rightListSet := createMap(right)

	var result int
	for value, timesLeft := range leftListSet {
		if timesRight, ok := rightListSet[value]; ok {
			result += timesLeft * (value * timesRight)
		}
	}

	return result
}

func parseList(filePath string) ([]int, []int) {
	lines := files.ReadFile(filePath)

	var left []int
	var right []int
	for _, element := range lines {
		splitLine := strings.Split(element, "   ")

		leftValue, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}

		rightValue, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	return left, right
}

func createMap(list []int) map[int]int {
	var set = make(map[int]int)
	for _, value := range list {
		set[value]++
	}

	return set
}
