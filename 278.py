# 278. First Bad Version

class Solution:
    def firstBadVersion(self, n: int) -> int:
        d = 0
        u = n
        m = 0
        while d <= u:
            m = ((d + u) / 2 + 0.5)
            if m == d or m == u:
                break
            if isBadVersion(m) == True:
                u = m
            else:
                d = m
        return int(m)
