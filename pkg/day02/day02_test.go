package day02

import (
	"advent-of-code-2024/pkg/parser"
	"strconv"
	"strings"
	"testing"
)

func TestDay02(t *testing.T) {
	const separator = " "
	tests := []struct {
		name          string
		filepath      string
		expectedPart1 int
		expectedPart2 int
	}{
		{
			name:          "example",
			filepath:      "./example.txt",
			expectedPart1: 2,
			expectedPart2: 4,
		},
		{
			name:          "input",
			filepath:      "./input.txt",
			expectedPart1: 598,
			expectedPart2: 634,
		},
	}

	for _, tt := range tests {
		levels := make([][]int, 0)

		lines, err := parser.Parse(tt.filepath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		for _, line := range lines {
			level, err := parseLine(line, separator)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			levels = append(levels, level)
		}

		t.Run(tt.name, func(t *testing.T) {
			actual := part1(levels)
			if actual != tt.expectedPart1 {
				t.Errorf("expected: %d, actual: %d", tt.expectedPart1, actual)
			}
		})

		//TODO: part 2
		t.Run(tt.name, func(t *testing.T) {
			actual := part2(levels)
			if actual != tt.expectedPart2 {
				t.Errorf("expected: %d, actual: %d", tt.expectedPart2, actual)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{1, 2, 3, 4},
			expected: [][]int{
				{2, 3, 4},
				{1, 3, 4},
				{1, 2, 4},
				{1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		actual := permutations(tt.input)
		if len(actual) != len(tt.expected) {
			t.Errorf("mismatched len, expected: %d, actual: %d", len(tt.expected), len(actual))
		}

		for i := 0; i < len(tt.expected); i++ {
			if len(actual[i]) != len(tt.expected[i]) {
				t.Errorf("mismatched len at [%d], expected: %d, actual: %d", i, len(tt.expected[i]), len(actual[i]))
			}
			for j := 0; j < len(tt.expected[i]); j++ {
				if actual[i][j] != tt.expected[i][j] {
					t.Errorf("at [%d][%d], expected: %d, actual: %d", i, j, actual[i][j], tt.expected[i][j])
				}
			}
		}
	}
}

func parseLine(line, separator string) ([]int, error) {
	words := strings.Split(line, separator)

	nums := make([]int, len(words))
	for i, word := range words {
		num, _ := strconv.Atoi(word)
		nums[i] = num
	}

	return nums, nil
}
