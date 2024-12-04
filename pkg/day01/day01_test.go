package day01

import (
	"advent-of-code-2024/pkg/parser"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name          string
		filepath      string
		expectedPart1 int
		expectedPart2 int
	}{
		{
			name:          "example",
			filepath:      "./example.txt",
			expectedPart1: 11,
			expectedPart2: 31,
		},
		{
			name:          "input",
			filepath:      "./input.txt",
			expectedPart1: 1579939,
			expectedPart2: 20351745,
		},
	}

	const separator = "   "
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// parser.Parse()
			lines, err := parser.Parse(tt.filepath)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			lefts := make([]int, 0)
			rights := make([]int, 0)
			for _, line := range lines {
				left, right, err := parseLine(line, separator)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				lefts = append(lefts, left)
				rights = append(rights, right)
			}

			// part 1
			actual, err := part1(lefts, rights)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if actual != tt.expectedPart1 {
				t.Errorf("expected: %d, actual: %d", tt.expectedPart1, actual)
			}

			// part 2
			actual = part2(lefts, rights)
			if actual != tt.expectedPart2 {
				t.Errorf("expected: %d, actual: %d", tt.expectedPart2, actual)
			}
		})
	}
}

// assumes no malformed input per line
func parseLine(line, separator string) (int, int, error) {
	words := strings.Split(line, separator)

	left, _ := strconv.Atoi(words[0])
	right, _ := strconv.Atoi(words[1])

	return left, right, nil
}
