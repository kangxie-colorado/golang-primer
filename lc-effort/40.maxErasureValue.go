// https://leetcode.com/problems/maximum-erasure-value/

package main

import (
	"fmt"
	"math"
)

/*
	tracking state: a hashmap map[int]int
	mapping each number to its occurance

	when a value >1, then the condition breaks.
*/

func _1_maximumUniqueSubarray(nums []int) int {
	m := make(map[int]int)
	i, j := 0, 0
	maxSum := 0
	sum := 0

	for j < len(nums) {
		m[nums[j]]++
		sum += nums[j]

		for m[nums[j]] > 1 {
			m[nums[i]]--
			sum -= nums[i]
			i++
		}

		maxSum = max(maxSum, sum)
		j++
	}

	return maxSum

}

// this problem is not so fit for non-shrinkable solution
// because the j changed and nums[j] changed... you would need to iterate thru the map..

func maxOfMap(m map[int]int) (int, int) {
	num, count := 0, math.MinInt
	for k, v := range m {
		if v > count {
			count = v
			num = k
		}
	}

	return num, count
}

func maximumUniqueSubarray(nums []int) int {
	// map the number to position
	m := make(map[int]int)

	maxSum := 0
	sum := 0

	for i, j := 0, 0; j < len(nums); {

		if _, found := m[nums[j]]; found {

			newI := m[nums[j]] + 1
			for k := i; k < newI; k++ {
				sum -= nums[k]
				delete(m, nums[k])
			}

			i = newI

		}

		sum += nums[j]

		m[nums[j]] = j
		j++
		maxSum = max(sum, maxSum)

	}

	return maxSum

}

/*
Wrong Answer
Details
Input
[1611,4668,7701,5885,7491,1645,665,7633,7705,652,1470,1798,7343,9641,3538,1700,5505,5328,6795,9011,6231,9845,2969,5281,3063,2579,1998,1897,3970,5543,6196,3198,4442,5399,1110,2982,7069,5187,1272,8090,8211,7403,9878,1053,6201,2280,1121,6661,194,6693,8951,6749,1932,4794,499,5833,9214,7789,8107,6662,4136,9677,6093,8142,2992,3958,695,9501,9714,1088,9748,5312,8055,2776,1064,2012,582,5621,375,2570,4152,8379,3942,906,367,1403,7409,1093,7324,1597,7517,5310,6005,182,1398,7989,4615,3721,6306,6272,2391,1585,8146,5000,6831,4793,5949,2088,4914,5322,7865,124,6917,1247,2366,4177,4516,6961,4741,8961,4089,8743,184,2045,2928,723,9573,4439,5809,7320,5954,2980,2113,4905,3337,2891,8409,433,2858,3255,483,166,1680,7881,6953,5030,7469,2715,8096,6786,8530,8212,8119,6144,8635,9274,3676,3273,4276,4450,1327,3602,8380,8704,9037,1283,3885,2531,438,2152,4689,4468,3931,7515,7965,7605,1902,180,857,3689,4065,1626,6315,8561,5932,8829,5223,8554,4200,7581,6526,7847,1208,5798,7642,3165,4198,4823,1820,6119,6621,9293,8616,576
View All
Output
1293564
Expected
1456646

possibly it misses counting some windows because the window

okay... maybe count the occurences is not right.
I see the guy tracking the position

- The main point here is for the subarray to contain unique elements for each index. Only the first subarrays starting from that index have unique elements.
- This can be solved using the two pointers technique

the hints..
after a few debugging
Runtime: 393 ms, faster than 16.67% of Go online submissions for Maximum Erasure Value.
Memory Usage: 9.4 MB, less than 41.67% of Go online submissions for Maximum Erasure Value.
Next challenges:

okay.. comparing to

class Solution {
public:
    int maximumUniqueSubarray(vector<int>& A) {
        int i = 0, ans = 0, N = A.size();
        unordered_map<int, int> m; // number -> index of last occurrence.
        vector<int> sum(N + 1);
        partial_sum(begin(A), end(A), begin(sum) + 1);
        for (int j = 0; j < N; ++j) {
            if (m.count(A[j])) i = max(i, m[A[j]] + 1);
            m[A[j]] = j;
            ans = max(ans, sum[j + 1] - sum[i]);
        }
        return ans;
    }
};

if (m.count(A[j])) i = max(i, m[A[j]] + 1);
	this is actually doing my map maintenance; taking care of multiple appearance in previous slots...

partial_sum(begin(A), end(A), begin(sum) + 1);
ans = max(ans, sum[j + 1] - sum[i]);
	also this, he maintains an extra array of sums...
	and use this to calculate a sum between i:j+1

	a bit too tricky

*/

func testmaximumUniqueSubarray() {
	fmt.Println(maximumUniqueSubarray([]int{449, 154, 934, 526, 429, 732, 784, 909, 884, 805, 635, 660, 742, 209, 742, 272, 669, 449, 766, 904, 698, 434, 280, 332, 876, 200, 333, 464, 12, 437, 269, 355, 622, 903, 262, 691, 768, 894, 929, 628, 867, 844, 208, 384, 644, 511, 908, 792, 56, 872, 275, 598, 633, 502, 894, 999, 788, 394, 309, 950, 159, 178, 403, 110, 670, 234, 119, 953, 267, 634, 330, 410, 137, 805, 317, 470, 563, 900, 545, 308, 531, 428, 526, 593, 638, 651, 320, 874, 810, 666, 180, 521, 452, 131, 201, 915, 502, 765, 17, 577, 821, 731, 925, 953, 111, 305, 705, 162, 994, 425, 17, 140, 700, 475, 772, 385, 922, 159, 840, 367, 276, 635, 696, 70, 744, 804, 63, 448, 435, 242, 507, 764, 373, 314, 140, 825, 34, 383, 151, 602, 745}))
}
