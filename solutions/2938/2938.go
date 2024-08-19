package _2938

func minimumSteps(s string) int64 {
	res := int64(0)
	lastZero := 0
	used := false
	for i, v := range []rune(s) {
		if v == '0' {
			if res != 0 || used {
				if i-lastZero == 1 {
					lastZero++
					continue
				}
				res += int64(i - lastZero - 1)
				lastZero++
			} else {
				res += int64(i - lastZero)
				used = true
			}
		}
	}

	return res
}
