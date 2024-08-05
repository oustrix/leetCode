package _2053

func KthDistinct(arr []string, k int) string {
	stringsCount := make(map[string]int, len(arr))

	for _, v := range arr {
		if _, ok := stringsCount[v]; !ok {
			stringsCount[v] = 1
		} else {
			stringsCount[v]++
		}
	}

	counter := 0
	for _, v := range arr {
		if stringsCount[v] == 1 {
			counter++
			if counter == k {
				return v
			}
		}
	}

	return ""
}
