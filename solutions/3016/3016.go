package _3016

import (
	"cmp"
	"slices"
)

func minimumPushes(word string) int {
	letters := make([]int, 26)
	for _, l := range word {
		letters[l-97]++
	}

	letters = slices.DeleteFunc(letters,
		func(n int) bool {
			return n == 0
		},
	)
	slices.SortFunc(letters, func(a, b int) int {
		return -cmp.Compare(a, b)
	})

	sum := 0
	for i, v := range letters {
		if i < 8 {
			sum += v
		} else if i < 16 {
			sum += v * 2
		} else if i < 24 {
			sum += v * 3
		} else {
			sum += v * 4
		}
	}

	return sum
}
