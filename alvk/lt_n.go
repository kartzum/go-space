package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func listNCalc(nums []int) (int, int) {
	if nums == nil {
		return 0, 0
	}
	if len(nums) == 1 {
		return 0, 0
	}
	l := 0
	r := 1
	s := 0
	f := 0
	current := int(^uint(0) >> 1)
	for l < len(nums) && r < len(nums) {
		dif := absInt(nums[l] - nums[r])
		if dif > current {
			if dif != 0 && dif < current {
				current = dif
				s = l
				f = r
			}
			l++
		} else {
			if dif != 0 && dif < current {
				current = dif
				s = l
				f = r
			}
			r++
		}
	}
	return nums[s], nums[f]
}

func ListNCalcTest() {
	fmt.Println(listNCalc([]int{1, 3, 4, 11}))
	fmt.Println(listNCalc([]int{1, 3, 5, 11, 12}))
}

func ListNCalc() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, t := range numsParts {
		v, _ := strconv.Atoi(t)
		nums = append(nums, v)
	}
	fmt.Println(listNCalc(nums))
}
