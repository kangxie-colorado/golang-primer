// leetcode 1504
// now I am sensing the company is tighening up
// let me go the extra 3 feet
// coding in this computer but submit in my home computer
// coding here because I can use two screens and switch up between work/lc quickly
// extra 3 feet but nothing will stop me
// the link I will just remember the problem number
// and I can share code by airplay..
// test case.. will be a little troublesome..
// let me see

package main

/*

I have read some discussions and watached some videos about monotnoical stacks
let me code some up

1st, the brute force one, I think even that has some smart trick in it
the idea is: for any left-top corner, count every possible 1s rect the intuition is when you encounter a 1, add the count by 1...
which is not really obvious but you can kind of see it when you walk thru the row
1,1,1,1

so walking by, you got 1, 2, 3, 4 (each time +1)
now adding 2nd row
1,1,1,1
1,1,1,1

walking by first row, 1,2,3,4
walking the 2nd, 5,6,7,8 <--- is that right? actually yes... the key is to to anchor the top-left corner

if you run into 0.. then you cannot really pass this column
1,1,0 .. 1,2
1,1,1 .. 3,4

*/

func countSubmat(pos point, mat [][]int) int {
	bound := len(mat[0])
	count := 0
	for x := pos.x; x < len(mat); x++ {
		for y := pos.y; y < bound; y++ {
			if mat[x][y] == 1 {
				count += 1
			} else {
				bound = y
			}
		}
	}

	return count
}

func numSubmat(mat [][]int) int {
	count := 0

	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 1 {
				count += countSubmat(point{x, y}, mat)
			}
		}
	}

	return count
}
