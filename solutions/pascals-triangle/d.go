/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package pascals_triangle

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
注意到第i行：
共i+1个元素；
两侧元素都是1；
其他每个元素都是上一行对齐的元素与对齐元素前一个元素的和，即f(i,j) = f(i-1, j-1) + f(i-1)(j)；
左右对称，即f(i，j) = f(i, i-j)

时间复杂度，O((1+2+...+n)/2) = O((1+n)*n/4) = O(n^2) // 每行只遍历一半数目
空间复杂度都是O(1+2+...+n) = O((1+n)*n/2) = O(n^2)
*/
func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	result := make([][]int, numRows)
	result[0], result[1] = []int{1}, []int{1, 1}
	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j <= (i+1)/2; j++ {
			value := result[i-1][j-1] + result[i-1][j]
			tmp[j], tmp[i-j] = value, value
		}
		result[i] = tmp
	}
	return result
}
