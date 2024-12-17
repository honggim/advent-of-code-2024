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
	nums := convertToIntArray(line)
	files := converToPartitions(nums)
	// printPartitions(files)

	// go backward
	for index := 1; len(files)-index > 0; index++ {
		srcIndex := len(files) - index

		srcFile := files[srcIndex]
		if srcFile.empty || srcFile.moved {
			// fmt.Printf("ignoring file: %+v\n", srcFile)
			// ignore these files
			continue
		}

		// go forward
		for dstIndex, dstFile := range files {
			if dstIndex >= srcIndex {
				// done checking
				break
			}
			if !dstFile.empty {
				continue
			}
			if dstFile.moved {
				continue
			}
			if dstFile.length < srcFile.length {
				continue
			}
			if dstFile.length == srcFile.length {
				// swap
				files[srcIndex].moved = true
				files[dstIndex], files[srcIndex] = files[srcIndex], files[dstIndex]
				// printPartitions(files)
				break // dst loop
			}
			if dstFile.length > srcFile.length {
				diff := dstFile.length - srcFile.length

				// swap
				files[srcIndex].empty = true

				files[dstIndex].empty = false
				files[dstIndex].moved = true
				files[dstIndex].length = srcFile.length
				files[dstIndex].id = srcFile.id

				// insert new empty partition
				newFiles := append(files[:dstIndex+1], files[dstIndex:]...)
				newFiles[dstIndex+1] = partition{
					empty:  true,
					moved:  false,
					length: diff,
				}
				files = newFiles

				// printPartitions(files)
				break // dst loop
			}
		}
	}

	return checksumPartitions(files)
}

type partition struct {
	id     int
	length int
	empty  bool
	moved  bool
}

func converToPartitions(lengths []int) []partition {
	files := make([]partition, 0)

	for index, length := range lengths {
		file := partition{
			id:     index / 2,
			length: length,
			empty:  index%2 == 1, // odd indexes are empty
			moved:  false,
		}

		files = append(files, file)
	}

	return files
}

func printPartitions(files []partition) {
	for _, file := range files {
		if file.empty {
			fmt.Print(strings.Repeat(emptyFile, file.length))
		} else {
			fmt.Print(strings.Repeat(strconv.Itoa(file.id), file.length))
		}
	}
	fmt.Println("")
}

func checksumPartitions(files []partition) int {
	sum := 0
	index := 0
	for _, file := range files {
		if file.empty {
			index += file.length
			continue
		}
		for i := 0; i < file.length; i++ {
			sum += (index + i) * file.id
		}
		index += file.length
	}

	return sum
}
