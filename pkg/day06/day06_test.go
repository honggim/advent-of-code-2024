package day06

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestDay06(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		fn       func([][]rune, int, int) int
		expected int
	}{
		{
			name:     "part 1: example",
			filepath: "example.txt",
			fn:       part1,
			expected: 41,
		},
		{
			name:     "part 2: example",
			filepath: "example.txt",
			fn:       part2,
			expected: 6,
		},
		{
			name:     "part 1: input",
			filepath: "input.txt",
			fn:       part1,
			expected: 4890,
		},
		{
			name:     "part 2: input",
			filepath: "input.txt",
			fn:       part2,
			expected: 1995,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
