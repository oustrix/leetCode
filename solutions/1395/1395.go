package _1395

func numTeams(rating []int) int {
	result := 0

	for i := range rating {
		for j := i + 1; j < len(rating); j++ {
			if rating[i] < rating[j] {
				for k := j + 1; k < len(rating); k++ {
					if rating[j] < rating[k] {
						result++
					}
				}
			} else if rating[i] > rating[j] {
				for k := j + 1; k < len(rating); k++ {
					if rating[j] > rating[k] {
						result++
					}
				}
			}
		}
	}

	return result
}
