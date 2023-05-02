// 9. Palindrome Number
package leetCode

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	length := (int)(math.Log10(float64(x)) + 1)
	for i := 0; i < length/2; i++ {
		f := (int)(math.Pow(10, (float64)(length-i*2-1)))
		first := x / f
		last := x % 10

		if first != last {
			return false
		}

		x -= first * f
		x -= last
		x /= 10
	}

	return true
}
