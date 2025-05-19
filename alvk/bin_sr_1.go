package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		value := nums[middle]
		if value == target {
			return middle
		}
		if target < value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func BinarySearch() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, str := range numsParts {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	tStr, _ := reader.ReadString('\n')
	target, _ := strconv.Atoi(strings.Trim(tStr, "\n"))
	fmt.Println(binarySearch(nums, target))
}

func BinarySearchTest() {
	fmt.Println(binarySearch([]int{1, 2, 3, 5, 10}, 5))
	fmt.Println(binarySearch([]int{5, 7, 9, 11, 13}, 6))
}
