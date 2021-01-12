package sum

/*
931. 下降路径最小和 https://leetcode-cn.com/problems/minimum-falling-path-sum/

给定一个方形整数数组 A，我们想要得到通过 A 的下降路径的最小和。

下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。
在下一行选择的元素和当前行所选元素最多相隔一列。



示例：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：12
解释：
可能的下降路径有：
[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
和最小的下降路径是 [1,4,7]，所以答案是 12。



提示：

1 <= A.length == A[0].length <= 100
-100 <= A[i][j] <= 100


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-falling-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
自顶向下递归会超时，可自底向上dp
*/
func minFallingPathSum(A [][]int) int {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}
	m, n := len(A), len(A[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for c := 0; c < n; c++ {
		dp[m-1][c] = A[m-1][c]
	}
	for r := m - 2; r >= 0; r-- {
		for c := 0; c < n; c++ {
			t := dp[r+1][c]
			if c > 0 && dp[r+1][c-1] < t {
				t = dp[r+1][c-1]
			}
			if c < n-1 && dp[r+1][c+1] < t {
				t = dp[r+1][c+1]
			}
			dp[r][c] = A[r][c] + t
		}
	}
	result := dp[0][0]
	for i := 1; i < n; i++ {
		if dp[0][i] < result {
			result = dp[0][i]
		}
	}
	return result
}
