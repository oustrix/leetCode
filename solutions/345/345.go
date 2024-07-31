package _345

import (
	"slices"
)

func reverseVowels(s string) string {
	str := []rune(s)

	vo := "aeiouAEIOU"
	vowels := []rune(vo)

	strOnlyVowels := make([]rune, 0, len(str))
	for _, v := range str {
		if slices.Contains(vowels, v) {
			strOnlyVowels = append(strOnlyVowels, v)
		}
	}
	slices.Reverse(strOnlyVowels)

	count := 0
	answer := make([]rune, 0, len(str))
	for _, v := range str {
		if slices.Contains(vowels, v) {
			answer = append(answer, strOnlyVowels[count])
			count++
		} else {
			answer = append(answer, v)
		}
	}

	return string(answer)
}
