package tasks

// 13. Roman to Integer

func romanToInt(s string) int {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0

	isPreviousI := false
	isPreviousX := false
	isPreviousC := false
	for _, v := range s {
		num += nums[v]
		switch v {
		case 'I':
			isPreviousI = true
			isPreviousX = false
			isPreviousC = false

		case 'V':
			if isPreviousI {
				num -= 2 * nums['I']
				isPreviousI = false
			}
			isPreviousX = false
			isPreviousC = false

		case 'X':
			if isPreviousI {
				num -= 2 * nums['I']
				isPreviousI = false
			}
			isPreviousX = true
			isPreviousC = false

		case 'L':
			if isPreviousX {
				num -= 2 * nums['X']
				isPreviousX = false
			}
			isPreviousI = false
			isPreviousC = false

		case 'C':
			if isPreviousX {
				num -= 2 * nums['X']
				isPreviousX = false
			}
			isPreviousC = true
			isPreviousI = false

		case 'D':
			fallthrough
		case 'M':
			if isPreviousC {
				num -= 2 * nums['C']
				isPreviousC = false
			}
			isPreviousX = false
			isPreviousI = false
		}

	}

	return num
}
