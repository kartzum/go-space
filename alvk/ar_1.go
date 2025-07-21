package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arraySchoolCalc(n int, nums []int) {
	if nums == nil {
		return
	}
	if n == 1 {
		return
	}
	left := 0
	right := 1
	for left < n && right < n {
		if nums[left] == 0 && nums[right] != 0 {
			t := nums[right]
			nums[left] = t
			nums[right] = 0
			left++
			right++
			continue
		}
		if nums[left] == 0 && nums[right] == 0 {
			right++
			continue
		}
		left++
		right++
	}
}

func ArraySchool() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, str := range numsParts {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	arraySchoolCalc(n, nums)
	for i, v := range nums {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}

func ArraySchoolTest() {
	a1 := []int{0, 0, 6, 0, 9, 8}
	arraySchoolCalc(6, a1)
	fmt.Printf("%v", a1)
}
