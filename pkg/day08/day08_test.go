package day08

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestDay8(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		fn       func([][]rune) int
		expected int
	}{
		{
			name:     "part 1: example",
			filepath: "./example.txt",
			fn:       part1,
			expected: 14,
		},
		{
			name:     "part 1: input",
			filepath: "./input.txt",
			fn:       part1,
			expected: 336,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			puzzle, err := parser.ParseIntoRuneMatrix(tt.filepath)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}

			actual := tt.fn(puzzle)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}
