// 9. Palindrome Number
package leetCode

import "strconv"

func isPalindrome(x int) bool {
	number := strconv.Itoa(x)
	for i := 0; i < len(number)/2; i++ {
		if number[i] != number[len(number)-1-i] {
			return false
		}
	}

	return true
}
