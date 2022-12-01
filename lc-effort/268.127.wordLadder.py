"""
https://leetcode.com/problems/word-ladder/

I have no idea so I looked at the related topic
BFS.. then I see the brute force way
"""


from collections import defaultdict
from typing import List


class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        wordSet = set(wordList + [beginWord, ])
        adjSets = defaultdict(set)

        for word in wordSet:
            for i, c in enumerate(word):
                for replaceChar in "abcdefghijklmnopqrstuvwxyz":
                    if replaceChar == c:
                        continue
                    replaceWord = word[:i]+replaceChar+word[i+1:]
                    if replaceWord in wordSet:
                        adjSets[word].add(replaceWord)

        q = [(beginWord, 1)]
        seen = set()
        while q:
            w, step = q[0]
            q = q[1:]
            if w in seen:
                continue
            seen.add(w)
            if w == endWord:
                return step

            for nei in adjSets[w]:
                q.append((nei, step+1))
        return 0


if __name__ == "__main__":
    s = Solution()
    print(s.ladderLength(beginWord="hit", endWord="cog",
          wordList=["hot", "dot", "dog", "lot", "log", "cog"]))

    print(s.ladderLength(beginWord="hit", endWord="cog",
          wordList=["hot", "dot", "dog", "lot", "log"]))
