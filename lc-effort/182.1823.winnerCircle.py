"""
https://leetcode.com/problems/find-the-winner-of-the-circular-game/

solved with brute force... by simulating the array/list

"""


class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        A = [i for i in range(1, n+1)]
        start = 0
        while len(A) > 1:
            out = (start+k-1) % len(A)
            A = A[:out]+A[out+1:]
            start = out

        return A[0]


"""
now thinking linear time and constant space

with n=5,k=2..
I can see the elemet getting out is 
2 4 6(1) 8(3)
..

hmm... cannot find out how to

reading these posts.. brilliant.. 
f(n,k) = (f(n-1,k)+k) %n # 0-based index...
https://leetcode.com/problems/find-the-winner-of-the-circular-game/discuss/1601186/C%2B%2B-oror-3-Approach-oror-Easy-Understanding

https://leetcode.com/problems/find-the-winner-of-the-circular-game/discuss/1157717/Java-full-solution-and-explanation

if I can draw a few cases.. from n=1 to n=5 (with k=2)
it is true... but hell... how can I even find out this..

I don't know what is the math behind it...

how to think of this
if k=1... it will of course be the last one standing..
if k=2... there is n-1, then adding n.. it will not go from n-1 to n.. n will always be skipped 
so it kind of will not make f(n-1) results change... ???
hellllll.... anyway... good to learn, not to dwell too much
"""
