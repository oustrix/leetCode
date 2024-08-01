package _278

func isBadVersion(version int) bool {
	panic(version)
}

func firstBadVersion(n int) int {
	a := 0
	b := n
	for a != b {
		c := (b + a) / 2
		if c == a {
			break
		}

		if isBadVersion(c) {
			b = c
		} else {
			a = c
		}
	}

	return b
}
