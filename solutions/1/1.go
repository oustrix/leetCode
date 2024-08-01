package _1

import (
	"slices"
)

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		idx := slices.Index(nums[i+1:], target-v)

		if idx != -1 {
			return []int{i, i + idx + 1}
		}
	}

	return []int{}
}
