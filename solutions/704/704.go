package _704

func search(nums []int, target int) int {
	a := 0
	b := len(nums)

	for a != b {
		c := (a + b) / 2

		value := nums[c]
		if value == target {
			return c
		}
		if c == a {
			break
		}

		if value > target {
			b = c
		} else {
			a = c
		}
	}

	return -1
}
