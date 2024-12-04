package day03

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filepath string
		fn       func(string) int
		expected int
	}{
		{
			filepath: "example_part1.txt",
			fn:       part1,
			expected: 161,
		},
		{
			filepath: "input.txt",
			fn:       part1,
			expected: 163931492,
		},
		{
			filepath: "example_part2.txt",
			fn:       part2,
			expected: 48,
		},
		{
			filepath: "input.txt",
			fn:       part2,
			expected: 76911921,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			input, err := parser.ParseIntoString(tt.filepath)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}

			actual := tt.fn(input)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}
