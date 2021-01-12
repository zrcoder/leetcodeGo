package image_overlap

import "image"

/*
835. 图像重叠
https://leetcode-cn.com/problems/image-overlap/

给出两个图像 A 和 B ，A 和 B 为大小相同的二维正方形矩阵。（并且为二进制矩阵，只包含0和1）。

我们转换其中一个图像，向左，右，上，或下滑动任何数量的单位，并把它放在另一个图像的上面。
之后，该转换的重叠是指两个图像都具有 1 的位置的数目。

（请注意，转换不包括向任何方向旋转。）

最大可能的重叠是什么？

示例 1:

输入：A = [[1,1,0],
          [0,1,0],
          [0,1,0]]

     B = [[0,0,0],
          [0,1,1],
          [0,0,1]]

输出：3
解释: 将 A 向右移动一个单位，然后向下移动一个单位。

注意:
1 <= A.length = A[0].length = B.length = B[0].length <= 30
0 <= A[i][j], B[i][j] <= 1
*/

/*
逆向思维：

先统计出两个矩阵里1出现的位置，统计结果放入两个集合
嵌套遍历两个集合，每两个点求一下偏移量，并用一个map记录每个偏移量出现的次数
最后统计出现最多次数的偏移量即可

时间复杂度 O(n^2 + X*Y), 其中X、Y分别为A、B中1的个数，最大分别为n^2,则最坏情况下，时间复杂度为O(n^4)
空间复杂度 O(n^2)
*/
func largestOverlap(A [][]int, B [][]int) int {
	onesInA, onesInB := calOnes(A, B)
	offsetsCounter := countOffsets(onesInA, onesInB)
	return getMax(offsetsCounter)
}

func calOnes(A, B [][]int) (onesInA, onesInB []image.Point) {
	n := len(A)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if A[r][c] == 1 {
				onesInA = append(onesInA, image.Pt(r, c))
			}
			if B[r][c] == 1 {
				onesInB = append(onesInB, image.Pt(r, c))
			}
		}
	}
	return
}

func countOffsets(onesInA, onesInB []image.Point) map[image.Point]int {
	m := make(map[image.Point]int, len(onesInA)*len(onesInB))
	for _, a := range onesInA {
		for _, b := range onesInB {
			offset := a.Sub(b)
			m[offset]++
		}
	}
	return m
}

func getMax(m map[image.Point]int) int {
	max := 0
	for _, c := range m {
		if c > max {
			max = c
		}
	}
	return max
}
