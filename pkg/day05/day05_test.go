package day05

import (
	"advent-of-code-2024/pkg/parser"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDay05(t *testing.T) {
	tests := []struct {
		filepath string
		expected int
	}{
		{
			filepath: "./example.txt",
			expected: 143,
		},
		{
			filepath: "./input.txt",
			expected: 5651,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			lines, err := parser.Parse(tt.filepath)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			pairs, pages := parseLines(lines)

			actual := part1(pairs, pages)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}

func parseLines(lines []string) ([][]int, [][]int) {
	pairs := make([][]int, 0)
	pages := make([][]int, 0)

	index := 0

	// pairs
	for i, line := range lines {
		if line == "" {
			index = i + 1
			break
		}
		pair := makePair(line)
		pairs = append(pairs, pair)
	}

	// pages
	for i := index; i < len(lines); i++ {
		page := makePage(lines[i])
		pages = append(pages, page)
	}

	return pairs, pages
}

func makePair(line string) []int {
	tokens := strings.Split(line, "|")
	key, _ := strconv.Atoi(tokens[0])
	val, _ := strconv.Atoi(tokens[1])

	return []int{key, val}
}

func makePage(line string) []int {
	tokens := strings.Split(line, ",")
	page := make([]int, len(tokens))

	for i, token := range tokens {
		num, _ := strconv.Atoi(token)
		page[i] = num
	}

	return page
}

func TestHelpers(t *testing.T) {
	// t.Skip()

	// makePair
	actualPair := makePair("97|45")
	expectedPair := []int{97, 45}

	if actualPair[0] != expectedPair[0] &&
		actualPair[1] != expectedPair[1] {
		t.Errorf("expected: %v, actual: %v", expectedPair, actualPair)
	}

	// generateHash
	pairs := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
	}

	hash := generateHash(pairs)
	if hash == nil {
		t.Errorf("hash is nil")
		for k, v := range hash {
			fmt.Println("key:", k)
			fmt.Println("   values:", v)
		}
	}
}
