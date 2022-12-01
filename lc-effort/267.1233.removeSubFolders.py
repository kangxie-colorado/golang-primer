"""
https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

let me try the trie

"""


from typing import List


class TrieNode:
    def __init__(self, val, terminal, children) -> None:
        self.val = val
        self.terminal = terminal
        self.children = children


class Solution:
    def removeSubfolders(self, folders: List[str]) -> List[str]:
        root = TrieNode("", False, {})

        for f in folders:
            dirs = f.split("/")[1:]
            curr = root
            for i, d in enumerate(dirs):
                terminal = i == len(dirs)-1
                node = None
                if d in curr.children:
                    # path up to here exists
                    node = curr.children[d]
                    if terminal:
                        node.terminal = True
                else:
                    node = TrieNode(d, terminal, {})
                    curr.children[d] = node

                if not terminal:
                    curr = node

        res = []

        def f(node, path):
            if node.terminal:
                res.append(path)
                return

            for val, child in node.children.items():
                f(child, path+"/"+val)
        f(root, '')
        return res


"""
yep.. this can also be solved by sorting and merging

"""


class Solution:
    def removeSubfolders(self, folders: List[str]) -> List[str]:
        folders.sort()
        i = j = 0
        res = []
        while True:
            if i == j:
                res.append(folders[i])
                j += 1

            if j == len(folders):
                break

            f_i, f_j = folders[i], folders[j]
            # no need to test len==len case, because in that situation
            # it can only be the same string and we know the path is unique.
            if len(f_i) < len(f_j) and f_i == f_j[:len(f_i)] and f_j[len(f_i)] == '/':
                j += 1
            else:
                i = j

        return res


"""
Runtime: 243 ms, faster than 95.20% of Python3 online submissions for Remove Sub-Folders from the Filesystem.
Memory Usage: 29.9 MB, less than 94.80% of Python3 online submissions for Remove Sub-Folders from the Filesystem.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.removeSubfolders(["/a", "/a/b/c", "/a/b/d"]))
    folder = ["/a", "/a/b", "/c/d", "/c/d/e", "/c/f"]
    print(s.removeSubfolders(folder))
    folder = ["/a/b/c", "/a/b/ca", "/a/b/d"]
    print(s.removeSubfolders(folder))
    folder = ["/a/b/c", "/a/b/ca", "/a/b/d", '/a']
    print(s.removeSubfolders(folder))
