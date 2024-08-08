package _2148

import (
	"slices"
)

func countElements(nums []int) int {
	ans := 0
	maxValue := slices.Max(nums)
	minValue := slices.Min(nums)

	for _, v := range nums {
		if v == maxValue || v == minValue {
			continue
		}

		ans++
	}

	return ans
}
