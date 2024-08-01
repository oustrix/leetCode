package _151

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(strings.TrimSpace(s), " ")

	answer := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			answer = append(answer, words[i])
		}
	}

	return strings.Join(answer, " ")
}
