// https://leetcode.com/tag/binary-search/discuss/786126/Python-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems

package main

/*
another good read
https://leetcode.com/problems/search-insert-position/discuss/249092/Come-on-forget-the-binary-search-patterntemplate!-Try-understand-it!

well... not really inspiring...
*/

/*
A good read
	- When to exit the loop? Should we use left < right or left <= right as the while loop condition?
	- How to initialize the boundary variable left and right?
	- How to update the boundary? How to choose the appropriate combination from left = mid, left = mid + 1 and right = mid, right = mid - 1?

left, right must cover all possible values, sometime +1, sometime to the biggest allowed
also transition the problem to a different problem

As for the question "When can we use binary search?",
my answer is that, If we can discover some kind of monotonicity, for example, if condition(k) is True then condition(k + 1) is True, then we can consider binary search.


*/

/*
1011. Capacity To Ship Packages Within D Days [Medium]

Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
Output: 15
Explanation:
A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
1st day: 1, 2, 3, 4, 5
2nd day: 6, 7
3rd day: 8
4th day: 9
5th day: 10

at first sight, it is really hard to see this can be solved by binary search
what the hell? what the fuck?

some kind of monotonicity -- then this..

yes, if the ship is biggest, the days needed is of course less
and we don't even need to concern its true relationship.. we just need to fact of monotonicity..

and then we can search
min=1, max=500

given the capacity, then it is easy to see, how many days are needed, because it ships in order
then I am looking for a min-capacity that can be shipped in 5 days.

I want to code this up

*/

/*
https://leetcode.com/problems/koko-eating-bananas/
875. Koko Eating Bananas [Medium]
Koko loves to eat bananas. There are N piles of bananas, the i-th pile has piles[i] bananas. The guards have gone and will come back in H hours. Koko can decide her bananas-per-hour eating speed of K. Each hour, she chooses some pile of bananas, and eats K bananas from that pile. If the pile has less than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.

Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back. Return the minimum integer K such that she can eat all the bananas within H hours.

Example :

Input: piles = [3,6,7,11], H = 8
Output: 4
Input: piles = [30,11,23,4,20], H = 5
Output: 30
Input: piles = [30,11,23,4,20], H = 6
Output: 23

again.. if I am not told this is a search problem I don't even know...

so the monotonicity? where is it? the more koko eats per hour, the quicker it can finish  (less hours)
but it can finish one pile at most at a time... so the speed is capped at max

apparently the slowest  koko eats would be 1.. and cannot do 0.. that is not eating.
so with these information available, I would know to use binary search
*/

/*
1482. Minimum Number of Days to Make m Bouquets [Medium]
Given an integer array bloomDay, an integer m and an integer k. We need to make m bouquets. To make a bouquet, you need to use k adjacent flowers from the garden. The garden consists of n flowers, the ith flower will bloom in the bloomDay[i] and then can be used in exactly one bouquet. Return the minimum number of days you need to wait to be able to make m bouquets from the garden. If it is impossible to make m bouquets return -1.

Examples:

Input: bloomDay = [1,10,3,10,2], m = 3, k = 1
Output: 3
Explanation: Let's see what happened in the first three days. x means flower bloomed and _ means flower didn't bloom in the garden.
We need 3 bouquets each should contain 1 flower.
After day 1: [x, _, _, _, _]   // we can only make one bouquet.
After day 2: [x, _, _, _, x]   // we can only make two bouquets.
After day 3: [x, _, x, _, x]   // we can make 3 bouquets. The answer is 3.
Input: bloomDay = [1,10,3,10,2], m = 3, k = 2
Output: -1
Explanation: We need 3 bouquets each has 2 flowers, that means we need 6 flowers. We only have 5 flowers so it is impossible to get the needed bouquets and we return -1.

damn.. even I know this is a Binary Search problem I see no easy way to tackle it
let alone if I don't know... it is hard to see

so ask is the minimum number of days..
so the longer you wait, the more blooms you will have..
apparently when there are not enough blooms, it will be -1

but at the max days, all flowers will be bloom, so if there is enough, then it will be enough... otherwise, -1, so -1 is only when there is not enough
so a little different is there is not a min*max thing going on
it is just waiting a bit long, the adjacent number of flowers will appear bigger

*/
