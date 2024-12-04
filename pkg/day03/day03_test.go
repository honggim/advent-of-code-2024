package day03

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filepath string
		expected int
	}{
		{
			filepath: "example.txt",
			expected: 161,
		},
		{
			filepath: "input.txt",
			expected: 163931492,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			input, err := parser.ParseIntoString(tt.filepath)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}

			actual := part1(input)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}
