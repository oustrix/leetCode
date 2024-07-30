package _605

func canPlaceFlowers(flowerbed []int, n int) bool {
	spots := 0

	prevEmpty := true
	for i, v := range flowerbed {
		if v == 1 {
			prevEmpty = false
			continue
		}

		if prevEmpty && i < len(flowerbed) {
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				spots++
				prevEmpty = false
			}
		} else {
			prevEmpty = true
		}
	}

	return n <= spots
}
