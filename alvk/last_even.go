package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lastEvenCalc(nums []int) int {
	var result int = -1
	for _, v := range nums {
		if v%2 == 0 {
			result = v
		}
	}
	return result
}

func LastEven() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, str := range numsParts {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	fmt.Println(lastEvenCalc(nums))
}

func LastEvenTest() {
	fmt.Println(lastEvenCalc([]int{7, 3, 4, 1, 10, 11, 12, 19, 21}))
}
