# 1. Two Sum

class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        for i, value in enumerate(nums):
            for j in range(i + 1, len(nums)):
                if value + nums[j] == target:
                    return [i, j]