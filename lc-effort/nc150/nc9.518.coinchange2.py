"""
https://leetcode.com/problems/coin-change-ii/

memory is still fresh on this one
but lets do it 

choose or not-choose the coin for an amount
"""


from typing import List


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        dp = [0]*(amount+1)
        dp[0] = 1

        for c in coins:
            for w in range(1, amount+1):
                if w >= c:
                    dp[w] += dp[w-c]

        return dp[amount]


"""
this is regular DP
first loop over row(coins)
then columns(amount)

if you do amount before coins.. then it is combination sum4 problem..
for every coin for every weight.. implicitly it means order doesn't matter.. 
for every weight for every coin... it mean considering all coins for each weight.. duplicates (with different order are allowed)

someone really explained it good
https://leetcode.com/problems/coin-change-ii/discuss/176706/Beginner-Mistake:-Why-an-inner-loop-for-coins-doensn't-work-Java-Soln/306232

"To get the correct answer, the correct dp definition should be
 dp[i][j]="number of ways to get sum 'j' using 'first i' coins". 
 Now when we try to traverse the 2-d array row-wise 
 by keeping only previous row array(to reduce space complexity), 
 we preserve the above dp definition as 
 dp[j]="number of ways to get sum 'j' using 'previous /first i coins' " 
 but when we try to traverse the 2-d array column-wise by keeping only the previous column, 
 the meaning of dp array changes to dp[i]="number of ways to get sum 'j' using 'all' coins"."
"""


if __name__ == "__main__":
    s = Solution()
    print(s.change(5, [1, 2, 5]))
