"""
https://leetcode.com/problems/design-add-and-search-words-data-structure/

not much idea to deal with dots
but let me try trie a bit
"""

from utils import TrieNode


class TrieNode:
    def __init__(self, val="", terminal=False) -> None:
        self.val = val
        self.terminal = terminal
        self.children = [None]*26  # 26 chars


class WordDictionary:

    def __init__(self):
        self.root = TrieNode()
        self.words = set()
        self.maxLen = 0

    def addWord(self, word: str) -> None:
        self.words.add(word)
        self.maxLen = max(self.maxLen, len(word))

        curr = self.root
        for i, c in enumerate(word):
            if not curr.children[ord(c)-ord('a')]:
                curr.children[ord(c)-ord('a')] = TrieNode(c, i == len(word)-1)
            elif i == len(word)-1:
                curr.children[ord(c)-ord('a')].terminal = True

            curr = curr.children[ord(c)-ord('a')]

    def search(self, word: str) -> bool:
        if word in self.words:
            return True
        if len(word) > self.maxLen:
            return False

        return self.searchWithDots(self.root, word, 0)

    def searchWithDots(self, trieNode, word, i):
        curr = trieNode
        if i >= len(word):
            return False
        c = word[i]

        if c != '.':
            if not curr.children[ord(c)-ord('a')]:
                return False

            if i == len(word)-1 and curr.children[ord(c)-ord('a')].terminal:
                return True

            return self.searchWithDots(curr.children[ord(c)-ord('a')], word, i+1)
        else:
            if i == len(word) - 1:
                return any([child and child.terminal for child in curr.children])
            else:
                for child in curr.children:
                    if child and self.searchWithDots(child, word, i+1):
                        return True

        return False


""""
Runtime: 9441 ms, faster than 65.39% of Python3 online submissions for Design Add and Search Words Data Structure.
Memory Usage: 85.1 MB, less than 5.02% of Python3 online submissions for Design Add and Search Words Data Structure.

Runtime: 6318 ms, faster than 82.41% of Python3 online submissions for Design Add and Search Words Data Structure.
Memory Usage: 84.9 MB, less than 5.02% of Python3 online submissions for Design Add and Search Words Data Structure.

hhah.. with maxLen and map.. it passed
but I did use more memory... but given the case said 25 max chars and max 10^4 calls.. it would be okay
"""
