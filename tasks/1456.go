package tasks

// 1456. Maximum Number of Vowels in a Substring of Given Length

func maxVowels(s string, k int) int {

	n, maxCount := len(s), 0
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	currCount := 0
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currCount++
		}
	}

	maxCount = currCount

	for i := 1; i <= n-k; i++ {
		if vowels[s[i-1]] {
			currCount--
		}
		if vowels[s[i+k-1]] {
			currCount++
		}
		if currCount > maxCount {
			maxCount = currCount
		}
	}

	return maxCount
}
