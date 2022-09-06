"""
https://leetcode.com/problems/my-calendar-i/

I just use bisect_left and bisect_right to find if any number in the bookings are
in my to-insert range... (kind of reverse looking)

and to simplify, end becomes end-1, because it is open (exclusive)
"""


from bisect import bisect_left, bisect_right
from sortedcontainers import SortedList


class MyCalendar:

    def __init__(self):
        self.bookings = []

    def book(self, start: int, end: int) -> bool:
        left = bisect_left(self.bookings, start)
        right = bisect_right(self.bookings, end-1)

        if left != right or left % 2 == 1:
            return False

        B = self.bookings
        self.bookings = B[:left] + [start, end-1] + B[left:]
        return True


"""
not bad at all..

Runtime: 355 ms, faster than 76.94% of Python3 online submissions for My Calendar I.
Memory Usage: 14.7 MB, less than 90.30% of Python3 online submissions for My Calendar I.
I did pass and actually pretty good
"""


class MyCalendar:

    def __init__(self):
        self.bookings = SortedList()

    def book(self, start: int, end: int) -> bool:
        left = SortedList.bisect_left(self.bookings, start)
        right = SortedList.bisect_right(self.bookings, end-1)

        if left != right or left % 2 == 1:
            return False

        self.bookings.append(start)
        self.bookings.append(end-1)
        return True


"""
Runtime: 449 ms, faster than 61.26% of Python3 online submissions for My Calendar I.
Memory Usage: 15 MB, less than 14.26% of Python3 online submissions for My Calendar I.
"""
