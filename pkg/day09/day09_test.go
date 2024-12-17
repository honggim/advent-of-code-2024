package day09

import (
	"advent-of-code-2024/pkg/parser"
	"testing"
)

func TestDay09(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		fn       func(string) int
		expected int
	}{
		// {
		// 	name:     "part 1: example",
		// 	filepath: "./example.txt",
		// 	fn:       part1,
		// 	expected: 1928,
		// },
		// {
		// 	name:     "part 1: what if more than single char",
		// 	filepath: "./example_30char.txt",
		// 	fn:       part1,
		// 	expected: 4006,
		// },
		// {
		// 	name:     "part 1: input",
		// 	filepath: "./input.txt",
		// 	fn:       part1,
		// 	expected: 6386640365805,
		// },
		{
			name:     "part 2: example",
			filepath: "./example.txt",
			// filepath: "./example_plus2.txt",
			fn:       part2,
			expected: 2858,
		},
		// {
		// 	name:     "part 2: example",
		// 	filepath: "./example_plus2.txt",
		// 	fn:       part2,
		// 	expected: 3261,
		// },
		{
			name:     "part 2: input",
			filepath: "./input.txt",
			fn:       part2,
			expected: 6423258376982,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := parser.Parse(tt.filepath)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}

			actual := tt.fn(lines[0]) // single line file
			if actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}
