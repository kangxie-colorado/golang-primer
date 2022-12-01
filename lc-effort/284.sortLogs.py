from functools import cmp_to_key
from typing import List


from functools import cmp_to_key
from typing import List


class Solution:
    def reorderLogFiles(self, logs: List[str]) -> List[str]:
        def isDigital(c):
            return '0' <= c <= '9'

        def compare(item1, item2):

            id1, rest1 = item1.split(" ", 1)
            id2, rest2 = item2.split(" ", 1)

            isDigital1, isDigital2 = isDigital(rest1[0]), isDigital(rest2[0])

            if isDigital1 or isDigital2:
                if isDigital1 and isDigital2:
                    return 1
                elif isDigital1:
                    return 1
                else:
                    return -1
            else:
                if rest1 < rest2:
                    return -1
                elif rest1 > rest2:
                    return 1
                else:
                    if id1 < id2:
                        return -1
                    else:
                        return 1

        logs.sort(key=cmp_to_key(compare))
        return logs

    def reorderLogFiles(self, logs: List[str]) -> List[str]:
        def isDigital(c):
            return '0' <= c <= '9'

        def compare(item1, item2):

            id1, rest1 = item1.split(" ", 1)
            id2, rest2 = item2.split(" ", 1)

            isDigital1, isDigital2 = isDigital(rest1[0]), isDigital(rest2[0])

            if isDigital1 or isDigital2:
                if isDigital1 and isDigital2:
                    return -1
                elif isDigital1:
                    return 1
                else:
                    return -1
            else:
                if rest1 < rest2:
                    return -1
                elif rest1 > rest2:
                    return 1
                else:
                    if id1 < id2:
                        return 1
                    else:
                        return -1

        logs.sort(key=cmp_to_key(compare))
        return logs


class Solution:
    def reorderLogFiles(self, logs: List[str]) -> List[str]:
        def isDigit(c):
            return '0' <= c <= '9'

        def compare(x1, x2):

            id1, rest1 = x1.split(" ", 1)
            id2, rest2 = x2.split(" ", 1)

            isDigit1, isDigit2 = isDigit(rest1[0]), isDigit(rest2[0])
            if isDigit1 or isDigit2:
                if isDigit1:
                    # keep the order
                    return 1
                else:
                    return -1
            else:
                if rest1 == rest2:
                    return -1 if id1 < id2 else 1
                else:
                    return -1 if rest1 < rest2 else 1

        return sorted(logs, key=cmp_to_key(compare))


"""
okay.. I saved memory but I wasted time
if you filter out the digit logs, then sorting space is smaller...
"""


if __name__ == '__main__':

    s = Solution()
    print(s.reorderLogFiles(logs=[
          "dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"]))
