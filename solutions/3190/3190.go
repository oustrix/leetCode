package _3190

func minimumOperations(nums []int) int {
	res := 0
	for _, v := range nums {
		if v%3 != 0 {
			res++
		}
	}

	return res
}
