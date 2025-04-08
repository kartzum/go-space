package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ArrayMax() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, t := range numsParts {
		v, _ := strconv.Atoi(t)
		nums = append(nums, v)
	}
	result := 0
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	fmt.Print(result)
}
