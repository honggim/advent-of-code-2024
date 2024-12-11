package day09

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(line string) int {
	// 2333133121414131402
	nums := convertToIntArray(line)

	// 00...111...2...333.44.5555.6666.777.888899
	disk := convertToDisk(nums)
	fmt.Println(strings.Join(disk, ""))

	// 0099811188827773336446555566..............
	defragDisk(disk)
	// fmt.Println(strings.Join(disk, ""))

	return checksum(disk)
}

func convertToIntArray(line string) []int {
	tokens := strings.Split(line, "")
	nums := make([]int, len(tokens))

	for i, token := range tokens {
		nums[i], _ = strconv.Atoi(token)
	}

	return nums
}

const emptyFile = "."

func convertToDisk(nums []int) []string {
	output := make([]string, 0)

	for index, num := range nums {
		val := emptyFile
		if index%2 == 0 {
			val = strconv.Itoa(index / 2)
		}
		output = append(output, getDisk(val, num)...)
	}

	return output
}

func getDisk(index string, num int) []string {
	output := make([]string, num)

	for i := 0; i < num; i++ {
		output[i] = index
	}

	return output
}

func defragDisk(disk []string) {
	first := 0
	last := len(disk) - 1

	for first < last {
		if disk[first] == emptyFile && disk[last] != emptyFile {
			disk[first], disk[last] = disk[last], disk[first]
		}

		if disk[first] != emptyFile {
			first++
		}

		if disk[last] == emptyFile {
			last--
		}
	}
}

func checksum(disk []string) int {
	sum := 0

	for i, data := range disk {
		if data == emptyFile {
			continue
		}
		num, _ := strconv.Atoi(data)
		sum += num * i
	}

	return sum
}

func part2(line string) int {
	files := convertToPartitions(line)
	// printPartitions(files)

	disk, replacementIndex := partitionsToString(files)
	// fmt.Println(disk)

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		replacementIndex -= file.size

		if file.val == emptyFile || file.occurences == 0 {
			continue
		}

		freeSpace := strings.Repeat(emptyFile, file.size)
		freeIndex := strings.Index(disk, freeSpace)
		if freeIndex == -1 {
			continue
		}
		if freeIndex < replacementIndex {
			disk = replaceAtIndex(disk, replacementIndex, freeSpace)

			replacement := strings.Repeat(file.val, file.occurences)
			disk = replaceAtIndex(disk, freeIndex, replacement)
			// fmt.Println(disk)
		}
	}

	fmt.Println(disk)
	stuff := strings.Split(disk, "")
	return checksum(stuff)

}

func replaceAtIndex(s string, index int, replacement string) string {
	return s[:index] + replacement + s[index+len(replacement):]
}

type partition struct {
	val        string
	size       int
	occurences int
	moved      bool
}

func convertToPartitions(line string) []partition {
	output := make([]partition, 0)
	nums := strings.Split(line, "")

	for index, num := range nums {
		val := emptyFile
		if index%2 == 0 {
			val = strconv.Itoa(index / 2)
		}
		occurrences, _ := strconv.Atoi(num)
		p := partition{
			val:        val,
			size:       occurrences * len(val),
			occurences: occurrences,
		}

		output = append(output, p)
	}

	return output
}

func printPartitions(files []partition) {
	for _, file := range files {
		for i := 0; i < file.occurences; i++ {
			fmt.Printf(file.val)
		}
	}
	fmt.Println("")
}

func partitionsToString(disk []partition) (string, int) {
	output := ""
	chars := 0

	for _, file := range disk {
		output += strings.Repeat(file.val, file.occurences)
		chars += file.size
	}

	return output, chars
}
