package day01

import (
	"errors"
	"sort"
)

func part1(left, right []int) (int, error) {
	if len(left) != len(right) {
		return -1, errors.New("unequal number of elements")
	}

	sum := 0

	sort.Ints(left)
	sort.Ints(right)

	for i, l := range left {
		r := right[i]
		sum += diff(l, r)
	}

	return sum, nil
}

func diff(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}

func part2(left, right []int) int {
	hash := make(map[int]int)

	for _, r := range right {
		hash[r]++
	}

	sum := 0

	for _, l := range left {
		multiplier, ok := hash[l]
		if ok {
			sum += (l * multiplier)
		}
	}

	return sum
}
