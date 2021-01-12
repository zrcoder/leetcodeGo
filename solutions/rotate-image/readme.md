# [48. 旋转图像](https://leetcode-cn.com/problems/rotate-image)
```
给定一个 n × n 的二维矩阵表示一个图像。
将图像顺时针旋转 90 度。

说明：
你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。
请不要使用另一个矩阵来旋转图像。

示例 1:
给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

示例 2:
给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],
原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
```
## 分析
![](rotateMatrix.png)

```
如图,对于点 p1(r, c), 顺时针旋转90°后到达点 p2, 其坐标是(c, n-1-r)；
类似地，p2 旋转后到达 p3，p3 旋转后到达 p4，p4旋转后恰恰到达p1，这样完成一个循环：

p1 → p2
↑     ↓
p4 ← p3
```
```go
func rotate(matrix [][]int) {
	n := len(matrix)
	for r := 0; r < (n+1)/2; r++ {
		for c := 0; c < n/2; c++ {
			// p1, p4, p3, p2
			matrix[r][c], matrix[n-1-c][r], matrix[n-1-r][n-1-c], matrix[c][n-1-r] =
			// p4, p3, p2, p1
				matrix[n-1-c][r], matrix[n-1-r][n-1-c], matrix[c][n-1-r], matrix[r][c]
		}
	}
}
```
如果是逆时针旋转，也是类似的，直接看代码
```go
func rotateAnticlockwise(s [][]int) {
	n := len(s)
	for r := 0; r < (n+1)/2; r++ {
		for c := 0; c < n/2; c++ {
			s[r][c], s[n-1-c][r], s[n-1-r][n-1-c], s[c][n-1-r] =
				s[c][n-1-r], s[r][c], s[n-1-c][r], s[n-1-r][n-1-c]
		}
	}
}
```