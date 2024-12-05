package day05

import "fmt"

func part1(pairs, pages [][]int) int {
	hash := generateHash(pairs)
	// printHash(hash)

	sum := 0
	for i, page := range pages {
		if middle, ok := checkPage(hash, page); ok {
			// fmt.Println("i:", i, "middle:", middle)
			sum += middle
		}
	}

	return sum
}

func printHash(hash map[int]map[int]bool) {
	for k, v := range hash {
		fmt.Println("key:", k)
		fmt.Println("   val:", v)
	}
}

func generateHash(pairs [][]int) map[int]map[int]bool {
	hash := make(map[int]map[int]bool)
	for _, pair := range pairs {
		key := pair[0]
		val := pair[1]

		if vals, ok := hash[key]; ok {
			vals[val] = true
			hash[key] = vals
		} else {
			hash[key] = map[int]bool{val: true}
		}
	}

	return hash
}

func checkPage(
	hash map[int]map[int]bool,
	page []int,
) (int, bool) {
	for i := 1; i < len(page); i++ {
		key := page[i-1]
		val := page[i]

		mini, ok := hash[key]
		if !ok {
			return 0, false
		}

		if ok := mini[val]; !ok {
			return 0, false
		}

	}

	mid := page[len(page)/2]
	return mid, true
}
