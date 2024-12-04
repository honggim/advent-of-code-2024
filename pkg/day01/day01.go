package day01

import (
	"advent-of-code-2024/pkg/math"
	"errors"
	"sort"
)

func part1(lefts, rights []int) (int, error) {
	if len(lefts) != len(rights) {
		return -1, errors.New("unequal number of elements")
	}

	sum := 0

	sort.Ints(lefts)
	sort.Ints(rights)

	for i, left := range lefts {
		right := rights[i]
		sum += math.AbsDiff(left, right)
	}

	return sum, nil
}

func part2(lefts, rights []int) int {
	hash := make(map[int]int)

	for _, right := range rights {
		hash[right]++
	}

	sum := 0

	for _, left := range lefts {
		multiplier, ok := hash[left]
		if ok {
			sum += (left * multiplier)
		}
	}

	return sum
}
