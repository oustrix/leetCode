# 704. Binary Search
import math


class Solution:
    def search(self, nums: list[int], target: int) -> int:
        s = 0
        e = len(nums)
        iter = math.log2(e)
        print(iter)
        for i in range(0, int(math.floor(math.log2(e))) + 1):
            m = (s + e) // 2
            if nums[m] == target:
                return m
            elif nums[m] > target:
                e = m
            else:
                s = m
        return -1