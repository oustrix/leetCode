package _1508

import (
	"slices"
)

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, n*(n+1)/2)

	for i := range nums {
		prevSum := 0
		for j := i; j < len(nums); j++ {
			prevSum += nums[j]
			sums = append(sums, prevSum)
		}
	}
	slices.Sort(sums)

	sum := 0
	for _, v := range sums[left-1 : right] {
		sum += v
	}
	return sum % 1000000007
}
