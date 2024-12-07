package day06

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestDay06(t *testing.T) {
	tests := []struct {
		filepath string
		fn       func([][]rune, int, int) int
		expected int
	}{
		{
			filepath: "example.txt",
			fn:       part1,
			expected: 41,
		},
		{
			filepath: "input.txt",
			fn:       part1,
			expected: 4890,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			lines, _ := parser.Parse(tt.filepath)
			puzzle, x, y := parseLines(lines)

			actual := tt.fn(puzzle, x, y)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}

func parseLines(lines []string) ([][]rune, int, int) {
	puzzle := make([][]rune, len(lines))

	x, y := -1, -1
	for yy, line := range lines {
		puzzle[yy] = []rune(line)

		if x == -1 && y == -1 {
			for xx, char := range puzzle[yy] {
				if char == '^' {
					y = yy
					x = xx
				}
			}
		}
	}

	return puzzle, x, y
}
