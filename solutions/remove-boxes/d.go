package remove_boxes

/*
546. 移除盒子
给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。
每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
当你将所有盒子都去掉之后，求你能获得的最大积分和。

示例：

输入：boxes = [1,3,2,2,2,3,4,3,1]
输出：23
解释：
[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 分)
----> [1, 3, 3, 3, 1] (1*1=1 分)
----> [1, 1] (3*3=9 分)
----> [] (2*2=4 分)

提示：
1 <= boxes.length <= 100
1 <= boxes[i] <= 100
*/

/*
一行内的消消乐~

1. 记忆化搜索
从左侧开始消除（改成从右侧消除也行）
*/
func removeBoxes(boxes []int) int {
	n := len(boxes)
	memo := gen3d(n)
	var cal func(left, right, k int) int
	cal = func(left, right, k int) int {
		if left > right {
			return 0
		}
		if memo[left][right][k] != 0 {
			return memo[left][right][k]
		}
		for left < right && boxes[left] == boxes[left+1] {
			left++
			k++
		}
		memo[left][right][k] = cal(left+1, right, 0) + (k+1)*(k+1)
		for i := left + 1; i <= right; i++ {
			if boxes[i] == boxes[left] {
				memo[left][right][k] = max(memo[left][right][k], cal(i, right, k+1)+cal(left+1, i-1, 0))
			}
		}
		return memo[left][right][k]
	}
	return cal(0, n-1, 0)
}

/*
dp；从右侧消除（改成从左侧消除也行）

两个方法的时间复杂度都是 O(n^4)
实测dp耗时是记忆化递归的10倍多, k的枚举比记忆化递归要多
*/
func removeBoxes1(boxes []int) int {
	n := len(boxes)
	dp := gen3d(n + 1)
	for size := 1; size <= n; size++ {
		for left := 0; left+size-1 < n; left++ {
			right := left + size - 1
			for k := 0; k < n; k++ {
				tmp := (k + 1) * (k + 1)
				if left <= right-1 {
					tmp += dp[left][right-1][0]
				}
				dp[left][right][k] = max(dp[left][right][k], tmp)
				for i := left; i < right; i++ {
					if boxes[i] == boxes[right] {
						tmp = dp[left][i][k+1]
						if i+1 <= right-1 {
							tmp += dp[i+1][right-1][0]
						}
						dp[left][right][k] = max(dp[left][right][k], tmp)
					}
				}
			}
		}
	}
	return dp[0][n-1][0]
}

func gen3d(n int) [][][]int {
	r := make([][][]int, n)
	for i := range r {
		r[i] = make([][]int, n)
		for j := range r[i] {
			r[i][j] = make([]int, n)
		}
	}
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
