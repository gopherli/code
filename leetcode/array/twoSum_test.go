package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	res := twoSumByBinary(numbers, target)
	fmt.Println(res)
}

// map法
func twoSumByMap(numbers []int, target int) []int {
	m := make(map[int]int)
	var num []int
	for k1, v1 := range numbers {
		m[v1] = k1 + 1
	}
	for k2, v2 := range numbers {
		if m[target-v2] > 0 {
			num = []int{m[target-v2], k2 + 1}
			if k2+1 < m[target-v2] {
				num = []int{k2 + 1, m[target-v2]}
			}
			return num
		}
	}
	return num
}

// 二分法
func twoSumByBinary(numbers []int, target int) []int {
	var num []int
	length := len(numbers)
	low := 0
	high := length
	for {
		mid := (low + high) / 2
		value := target - numbers[mid]
		if value == numbers[mid] {
			num = getNum(numbers, value, mid)
			return num
		}
		if value > numbers[mid] {
			low = mid
		}
		if value < numbers[mid] {
			high = mid
		}
	}
	return num
}

func getNum(numbers []int, value, mid int) []int {
	num := []int{}
	for k, v := range numbers {
		if value == v {
			num = []int{k + 1, mid + 1}
			if k > mid {
				num = []int{mid + 1, k + 1}
			}
		}
	}
	return num
}
