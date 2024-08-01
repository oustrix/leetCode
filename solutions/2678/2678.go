package _2678

import (
	"strconv"
)

func countSeniors(details []string) int {
	res := 0
	for _, v := range details {
		age, _ := strconv.Atoi(v[11:13])

		if age > 60 {
			res++
		}
	}

	return res
}
