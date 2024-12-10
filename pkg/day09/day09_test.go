package day09

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestDay09(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		fn       func([]string) int
		expected int
	}{
		{
			name:     "part 1: example",
			filepath: "./example.txt",
			fn:       part1,
			expected: 1928,
		},
		{
			name:     "part 1: input",
			filepath: "./input.txt",
			fn:       part1,
			expected: 6386640365805,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := parser.Parse(tt.filepath)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}

			actual := tt.fn(lines)
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}
