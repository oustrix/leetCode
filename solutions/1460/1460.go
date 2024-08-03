package _1460

import (
	"slices"
)

func canBeEqual(target []int, arr []int) bool {
	slices.Sort(target)
	slices.Sort(arr)

	return slices.Equal(target, arr)
}
