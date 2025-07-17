package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				t := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = t
			}
		}
	}
}

func BubbleSort2() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, t := range numsParts {
		v, _ := strconv.Atoi(t)
		nums = append(nums, v)
	}
	bubbleSort2(nums)
	for i, current := range nums {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(current)
	}
}

func BubbleSort2Test() {
	nums1 := []int{0, 0, 6, 0, 9, 8}
	bubbleSort2(nums1)
	fmt.Printf("%v", nums1)
}
