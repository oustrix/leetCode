// 1822. Sign of the Product of an Array

func arraySign(nums []int) int {
    negative := false
    for _, num := range nums {
        if num == 0 {
            return 0
        } else if num < 0 {
            negative = !negative
        }
    }

    if negative {
        return -1
    } else {
        return 1
    }
}
