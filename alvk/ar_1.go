package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arraySchoolCalc(n int, arr []int) []int {
	if arr == nil {
		return nil
	}
	if n == 1 {
		return arr
	}
	left := 0
	right := 1
	for left < n && right < n {
		if arr[left] == 0 && arr[right] != 0 {
			t := arr[right]
			arr[left] = t
			arr[right] = 0
			left++
			right++
			continue
		}
		if arr[left] == 0 && arr[right] == 0 {
			right++
			continue
		}
		left++
		right++
	}
	return arr
}

func ArraySchool() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	arrStr, _ := reader.ReadString('\n')
	strNumbers := strings.Fields(arrStr)
	var arr []int
	for _, str := range strNumbers {
		num, _ := strconv.Atoi(str)
		arr = append(arr, num)
	}
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	arr = arraySchoolCalc(n, arr)
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}

func ArraySchoolTests() {
	fmt.Printf("%v", arraySchoolCalc(6, []int{0, 0, 6, 0, 9, 8}))
}
