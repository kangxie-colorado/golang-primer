"""
https://leetcode.com/problems/word-search-ii/


"""


from concurrent.futures import ThreadPoolExecutor
from typing import List


class TrieNode:
    def __init__(self, val="", terminal=False) -> None:
        self.val = val
        self.terminal = terminal
        self.children = [None]*26  # 26 chars

    def insert(self, word):
        curr = self
        for c in word:
            if not curr.children[ord(c)-ord('a')]:
                curr.children[ord(c)-ord('a')] = TrieNode(c)
            curr = curr.children[ord(c)-ord('a')]
        curr.terminal = True

    def search(self, word):
        curr = self
        for c in word:
            if not curr.children[ord(c)-ord('a')]:
                return False
            curr = curr.children[ord(c)-ord('a')]

        return curr.terminal


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        trie = TrieNode()
        for w in words:
            trie.insert(w)
            # trie.insert(w[::-1])
        m, n = len(board), len(board[0])
        res = []

        def backtrack(trieNode, run, r, c):
            if r < 0 or c < 0 or r >= m or c >= n or board[r][c] == "X" or trieNode is None:
                return

            save = board[r][c]
            # this char exists in the children nodes of current trieNode
            child = trieNode.children[ord(board[r][c]) - ord('a')]
            if child:
                if child.terminal:
                    res.append(run+child.val)
                    child.terminal = False

                board[r][c] = "X"
                backtrack(child, run+child.val, r+1, c)
                backtrack(child, run+child.val, r-1, c)
                backtrack(child, run+child.val, r, c+1)
                backtrack(child, run+child.val, r, c-1)
                board[r][c] = save

        for i in range(m):
            for j in range(n):
                backtrack(trie, "", i, j)
        return res


"""
huh... the solution I watched.. all pretty like what I did
I think they will TLE as well..


I am thinking let me try parallel computing

    def exist(self, board: List[List[str]], word: str) -> bool:
        m,n = len(board), len(board[0])
        def backtrack(i, r,c):
            if r<0 or c<0 or r>=m or c>=n or word[i]!=board[r][c]:
                return False
            if i==len(word)-1:
                return True

            save = board[r][c]
            board[r][c] = 'X'
            ret = backtrack(i+1, r+1,c) or \
                backtrack(i+1, r-1,c) or \
                backtrack(i+1, r,c+1) or\
                backtrack(i+1, r,c-1)
            board[r][c] = save
            return ret

        if len(set(word[:len(word)//2])) == 1:
            word = word[::-1]

        for i in range(m):
            for j in range(n):
                if backtrack(0, i,j):
                    return True
        return False


this is the word search I, with a hack to reverse the word if half of it is repeating chars
"""

res = []


def exist(board: List[List[str]], word: str) -> bool:
    m, n = len(board), len(board[0])
    visited = set()

    def backtrack(i, r, c):
        if r < 0 or c < 0 or r >= m or c >= n or board[r][c] != searchWord[i] or (r, c) in visited:
            return False
        if i == len(searchWord)-1:
            return True

        visited.add((r, c))
        ret = backtrack(i+1, r+1, c) or \
            backtrack(i+1, r-1, c) or \
            backtrack(i+1, r, c+1) or\
            backtrack(i+1, r, c-1)
        visited.remove((r, c))
        return ret

    searchWord = word
    if len(set(word[:len(word)//2])) < len(set(word[len(word)//2+1:])):
        searchWord = word[::-1]

    for i in range(m):
        for j in range(n):
            if backtrack(0, i, j):
                res.append(word)
                break


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        with ThreadPoolExecutor(max_workers=50) as executor:
            for word in words:
                executor.submit(exist, board, word)

        return res


"""
Runtime: 1016 ms, faster than 91.50% of Python3 online submissions for Word Search II.
Memory Usage: 15.8 MB, less than 65.29% of Python3 online submissions for Word Search II.


okay... 
"""


if __name__ == '__main__':
    print(Solution().findWords([["a", "b"], ["c", "d"]],
                               ["abd"]))
    board = [["o", "a", "a", "n"], ["e", "t", "a", "e"],
             ["i", "h", "k", "r"], ["i", "f", "l", "v"]]
    words = ["oath", "pea", "eat", "rain"]
    print(Solution().findWords(board, words))
