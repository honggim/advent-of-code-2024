package day02

func part1(levels [][]int) int {
	sum := 0
	for _, level := range levels {
		if evaluate(level) {
			sum++
		}
	}
	return sum
}

func evaluate(level []int) bool {

	var xor bool
	for i := 0; i < len(level)-1; i++ {
		if i == 0 {
			xor = isPositive(level[i], level[i+1])
		}
		if xor != isPositive(level[i], level[i+1]) {
			return false
		}

		abs := absDiff(level[i], level[i+1])
		if abs == 0 || abs > 3 {
			return false
		}
	}

	return true
}

// TODO:
func absDiff(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}

func isPositive(x, y int) bool {
	return (x - y) >= 0
}
