package day04

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestDay04(t *testing.T) {
	tests := []struct {
		filepath string
		fn       func([][]rune) int
		expected int
	}{
		{
			filepath: "./example.txt",
			fn:       part1,
			expected: 18,
		},
		{
			filepath: "./example.txt",
			fn:       part2,
			expected: 9,
		},
		{
			filepath: "./input.txt",
			fn:       part1,
			expected: 2613,
		},
		{
			filepath: "./input.txt",
			fn:       part2,
			expected: 1905,
		},
	}

	for _, tt := range tests {
		lines, err := parser.Parse(tt.filepath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		puzzle := make([][]rune, len(lines))
		for i, line := range lines {
			puzzle[i] = []rune(line)
		}

		t.Run(tt.filepath, func(t *testing.T) {
			actual := tt.fn(puzzle)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}

}
