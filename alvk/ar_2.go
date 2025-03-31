package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ArrayRemove() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	numsStr, _ := reader.ReadString('\n')
	gStr, _ := reader.ReadString('\n')
	g, _ := strconv.Atoi(strings.TrimSpace(gStr))
	numsParts := strings.Fields(numsStr)
	var nums []int
	for _, t := range numsParts {
		v, _ := strconv.Atoi(t)
		if v != g {
			nums = append(nums, v)
		}
	}
	for i, current := range nums {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(current)
	}
}
