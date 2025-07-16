package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exponentialSearch(arr []int, x int) (int, int) {
	n := len(arr)

	/*if arr[0] == x {
		return 0, 0
	}*/
	if arr[0] > x {
		return -1, -1
	}

	index := 1
	for index < n && arr[index] <= x {
		index *= 2
	}

	left := index / 2
	right := 0
	if index > n-1 {
		right = n - 1
	} else {
		right = index
	}

	return left, right
}

func ExponentialSearch() {
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
	left, right := exponentialSearch(nums, target)
	fmt.Println(left, right)
}

func ExponentialSearchTest() {
	fmt.Println(exponentialSearch([]int{8, 11, 12, 16, 18, 21, 33, 36, 48, 54, 63}, 16))
}
