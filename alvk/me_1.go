package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func majorityElement(nums []int) int {
	freq := make(map[int]int)
	n := len(nums)

	for _, num := range nums {
		freq[num]++
		if freq[num] > n/2 {
			return num
		}
	}

	return -1
}

func MajorityElement() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, t := range numsParts {
		v, _ := strconv.Atoi(t)
		nums = append(nums, v)
	}
	result := majorityElement(nums)
	fmt.Print(result)
}

func MajorityElementTest() {
	result := majorityElement([]int{7, 7, 8, 8, 8, 8, 9})
	fmt.Print(result)
}
