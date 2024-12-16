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
	fmt.Printf("Part 1: %d\n", MiddlePageNumbersSum("input.txt"))
	fmt.Printf("Part 2 %d\n", MiddlePageNumbersSumOfIncorectOrderedUpdates("input.txt"))
}

func MiddlePageNumbersSum(filePath string) int {
	var result int
	rules := getRules(filePath)
	updates := getUpdates(filePath)
	for _, pageNumbers := range updates {
		if arePagesInOrder(pageNumbers, rules) {
			result += pageNumbers[len(pageNumbers)/2]
		}
	}
	return result
}

func MiddlePageNumbersSumOfIncorectOrderedUpdates(filePath string) int {
	var result int
	rules := getRules(filePath)
	updates := getUpdates(filePath)
	for _, pageNumbers := range updates {
		if !arePagesInOrder(pageNumbers, rules) {
			reorderedNumbers := reorder(pageNumbers, rules)
			result += reorderedNumbers[len(reorderedNumbers)/2]
		}
	}
	return result
}

func arePagesInOrder(pageNumbers []int, rules map[int][]int) bool {
	for i := 1; i < len(pageNumbers); i++ {
		pageNumber := pageNumbers[i]

		if pageNumberOrderRules, ok := rules[pageNumber]; ok {
			for _, orderRule := range pageNumberOrderRules {
				if array.InArray(orderRule, pageNumbers[:i]) {
					return false
				}
			}
		}
	}

	return true
}

func reorder(pageNumbers []int, rules map[int][]int) []int {
	var result []int
	result = append(result, pageNumbers[0])
	for i := 1; i < len(pageNumbers); i++ {
		pageNumber := pageNumbers[i]
		shouldAppend := true
		for i, prevNumber := range result {
			prevNumberOrderRules, ok := rules[prevNumber]
			if !ok {
				tmp := slices.Concat(result[:i], []int{pageNumber})
				result = append(tmp, result[i:]...)
				shouldAppend = false
				break
			}
			if array.InArray(pageNumber, prevNumberOrderRules) {
				shouldAppend = true
				continue
			}
			if array.InArray(prevNumber, rules[pageNumber]) {
				tmp := slices.Concat(result[:i], []int{pageNumber})
				result = append(tmp, result[i:]...)
				shouldAppend = false
				break
			}
		}

		if shouldAppend {
			result = append(result, pageNumber)
		}
	}

	return result
}

func getRules(filePath string) map[int][]int {
	result := make(map[int][]int)
	lines := files.ReadFile(filePath)
	for _, values := range lines {
		if strings.Contains(values, ",") {
			break
		}

		splitLine := strings.Split(values, "|")
		leftNumber, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}
		rightNumber, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}

		if val, ok := result[leftNumber]; ok {
			result[leftNumber] = append(val, rightNumber)
		} else {
			result[leftNumber] = []int{rightNumber}
		}
	}

	return result
}

func getUpdates(filePath string) [][]int {
	var result [][]int
	lines := files.ReadFile(filePath)
	for _, values := range lines {
		if strings.Contains(values, "|") {
			continue
		}

		var lineResult []int
		splitLine := strings.Split(values, ",")
		for _, val := range splitLine {
			value, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}

			lineResult = append(lineResult, value)
		}

		result = append(result, lineResult)
	}

	return result
}
