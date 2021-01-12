# 二维区域子矩阵和
给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)

## [区域固定，各个元素保证不会改变](https://leetcode-cn.com/problems/range-sum-query-2d-immutable)

```
示例:
给定 matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12

说明:
你可以假设矩阵不可变。
会多次调用 sumRegion 方法。
你可以假设 row1 ≤ row2 且 col1 ≤ col2。
```

* 前缀和技巧

```go
/*
1. 暴力法超出时间限制
2. 可以用一个二维数组，维护每一行的前缀和，在计算区域和的时候利用前缀和技巧在O(row2-row1+1)复杂度计算结果
3. 前缀和可以扩展到二维，遍历原矩阵，存储从左上角到每个点的子矩阵的和，不妨称作矩阵前缀和；
可以利用这个前缀和矩阵在常数时间得到结果：

O 。。。。。。。
。。A 。。B 。。
。。。。。。。。
。。C 。。D 。。
。。。。。。。。

sum(ABCD) = prefixSum(D) - prefixSum[B] - prefixSum[C] + prefixSum[A]
注意顶点是否包含的细节

prefixSum矩阵可以比原矩阵多一行一列，减少边界处理
*/

type NumMatrix struct {
	prefixSum [][] int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dp[r+1][c+1] = dp[r][c+1] + dp[r+1][c] - dp[r][c] + matrix[r][c]
		}
	}
	return NumMatrix{prefixSum: dp}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return m.prefixSum[row2+1][col2+1] - m.prefixSum[row1][col2+1] - m.prefixSum[row2+1][col1] + m.prefixSum[row1][col1]
}
```
## [假如矩阵可以修改](https://leetcode-cn.com/problems/range-sum-query-2d-mutable)
扩展上面的问题，假设增加一个api Update(row int, col int, val int) ，可以修改矩阵，要怎么做？

* 前缀和技巧,详细实现见[代码](d.go)

```
可以维护每一行的前缀和，在修改时更新对应行的前缀和即可；
计算子区域和的时候，从row1向row2遍历，每一行用前缀和相减的技巧得到每一行的和，逐行累加即可
时间复杂度：update是O(col2-col1+1),sumRegion是O(row2-row1+1)
空间复杂度O(m*n)
```

* 另有一个借助线段树的实现

```
type NumMatrix struct {
	root   *SegTreeNode
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	return NumMatrix{
		root:   buildTree(matrix, 0, 0, len(matrix)-1, len(matrix[0])-1),
		matrix: matrix,
	}
}

func (n *NumMatrix) Update(row int, col int, val int) {
	if n.root == nil {
		return
	}
	n.root.Update(row, col, val)
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if n.root == nil {
		return 0
	}
	return n.root.Query(row1, col1, row2, col2)
}
```

SegTree详细实现见[代码](segtree.go)

关于线段树的应用，另可参考[leetcode天际线问题](https://leetcode-cn.com/problems/the-skyline-problem/)
及[一个小伙的题解](https://leetcode-cn.com/problems/the-skyline-problem/solution/xian-duan-shu-he-sao-miao-xian-suan-fa-jie-jue-ci-/)
