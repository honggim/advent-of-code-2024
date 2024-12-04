package day04

import "fmt"

var (
	X = 'X'
	M = 'M'
	A = 'A'
	S = 'S'
)

// p(uzzle)
func part1(p [][]rune) int {
	sum := 0

	rows := len(p)
	cols := len(p[0])
	const min = 3 // minimal distance from edge

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if p[r][c] == X {
				// N
				if r >= min && verifyMAS(p[r-1][c], p[r-2][c], p[r-3][c]) {
					debug(fmt.Sprintf("r: %d, c: %d, N\n", r, c))
					sum++
				}

				// NE
				if r >= min && c < cols-min && verifyMAS(p[r-1][c+1], p[r-2][c+2], p[r-3][c+3]) {
					debug(fmt.Sprintf("r: %d, c: %d, NE\n", r, c))
					sum++
				}

				// E
				if c < cols-min && verifyMAS(p[r][c+1], p[r][c+2], p[r][c+3]) {
					debug(fmt.Sprintf("r: %d, c: %d, E\n", r, c))
					sum++
				}

				// SE
				if r < rows-min && c < cols-min && verifyMAS(p[r+1][c+1], p[r+2][c+2], p[r+3][c+3]) {
					debug(fmt.Sprintf("r: %d, c: %d, SE\n", r, c))
					sum++
				}

				// S
				if r < rows-min && verifyMAS(p[r+1][c], p[r+2][c], p[r+3][c]) {
					debug(fmt.Sprintf("r: %d, c: %d, S\n", r, c))
					sum++
				}

				// SW
				if r < rows-min && c >= min && verifyMAS(p[r+1][c-1], p[r+2][c-2], p[r+3][c-3]) {
					debug(fmt.Sprintf("r: %d, c: %d, SW\n", r, c))
					sum++
				}

				// W
				if c >= min && verifyMAS(p[r][c-1], p[r][c-2], p[r][c-3]) {
					debug(fmt.Sprintf("r: %d, c: %d, W\n", r, c))
					sum++
				}

				// NW
				if r >= min && c >= min && verifyMAS(p[r-1][c-1], p[r-2][c-2], p[r-3][c-3]) {
					debug(fmt.Sprintf("r: %d, c: %d, NW\n", r, c))
					sum++
				}
			}
		}
	}

	return sum
}

const print = false

func debug(msg string) {
	if print {
		fmt.Println(msg)
	}
}

func verifyMAS(maybeM, maybeA, maybeS rune) bool {
	return maybeM == M && maybeA == A && maybeS == S
}

// p(uzzle)
func part2(p [][]rune) int {
	sum := 0
	rows := len(p)
	cols := len(p[0])

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if p[r][c] == A {
				nw := p[r-1][c-1]
				ne := p[r-1][c+1]
				sw := p[r+1][c-1]
				se := p[r+1][c+1]

				// N: MM, S: SS
				if nw == M && ne == M && sw == S && se == S {
					sum++
				}

				// N: SS, S: MM
				if nw == S && ne == S && sw == M && se == M {
					sum++
				}

				// W: MM, E: SS
				if nw == M && ne == S && sw == M && se == S {
					sum++
				}

				// W: SS, E: MM
				if nw == S && ne == M && sw == S && se == M {
					sum++
				}
			}
		}
	}

	return sum
}
