"""
https://leetcode.com/problems/simplify-path/

typical stack application here
"""


class Solution:
    def simplifyPath(self, path: str) -> str:
        s = []

        paths = path.split('/')
        for p in paths:
            if p == '..':
                if len(s) > 0:
                    s.pop()
            elif p == '.' or p == '':
                ...
            else:
                s.append(p)

        return '/' + '/'.join(s) if len(s) > 0 else '/'


"""
Runtime: 48 ms, faster than 62.51% of Python3 online submissions for Simplify Path.
Memory Usage: 13.9 MB, less than 80.94% of Python3 online submissions for Simplify Path.

but using split is kind of like cheating...
but even without using split.. I don't want to deal with '/'..
I'll just still work around '/'
"""


class Solution:
    def simplifyPath(self, path: str) -> str:
        s = []
        i = 0
        lastSlash = 0
        path += '/'
        while i < len(path):
            if path[i] == '/':
                # deal with what is between lastSlah and thisOne
                if i - lastSlash > 1:
                    p = path[lastSlash+1:i]
                    if p == '..':
                        if len(s) > 0:
                            s.pop()
                    elif p == '.':
                        ...
                    else:
                        s.append(p)

                lastSlash = i
            i += 1
        # deal with last segment? no.. just add a '/' to make it generic
        return '/' + '/'.join(s) if len(s) > 0 else '/'


"""
Runtime: 55 ms, faster than 44.94% of Python3 online submissions for Simplify Path.
Memory Usage: 13.9 MB, less than 80.94% of Python3 online submissions for Simplify Path.

yeah.. so kind of don't store '/', which would actually complicates things a lot
"""


if __name__ == '__main__':
    print(Solution().simplifyPath("/home/"))
    print(Solution().simplifyPath("/../"))
    print(Solution().simplifyPath("/home//foo/"))
