package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AbsDiff(x, y int) int {
	diff := x - y
	return Abs(diff)
}

func IsPositive(x, y int) bool {
	return (x - y) > 0
}
func IsNegative(x, y int) bool {
	return (x - y) < 0
}
