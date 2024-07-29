package _1768

func mergeAlternately(word1 string, word2 string) string {
	b := make([]uint8, 0, len(word1)+len(word2))

	if len(word1) > len(word2) {
		for i := range len(word1) {
			b = append(b, word1[i])

			if i < len(word2) {
				b = append(b, word2[i])
			}
		}
	} else {
		for i := range len(word2) {
			if i < len(word1) {
				b = append(b, word1[i])
			}

			b = append(b, word2[i])
		}
	}

	return string(b)
}
