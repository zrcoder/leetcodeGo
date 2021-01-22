/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package flood_fill

import "container/list"

/*
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，
接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

示例 1:

输入:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析:
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
注意:

image 和 image[0] 的长度在范围 [1, 50] 内。
给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flood-fill
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
意思是将与给定点相连且颜色相同的点都改变颜色

DFS递归，时空复杂度都是O(N)
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return nil
	}
	oldColor := image[sr][sc]
	var fill func(r, c int)
	fill = func(r, c int) { // DFS
		if r < 0 || r == len(image) || c < 0 || c == len(image[0]) ||
			newColor == image[r][c] || image[r][c] != oldColor {
			return
		}
		image[r][c] = newColor
		fill(r-1, c)
		fill(r+1, c)
		fill(r, c-1)
		fill(r, c+1)
	}
	fill(sr, sc)
	return image
}

/*
BFS迭代。时空复杂度都是O(N)
*/
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return nil
	}
	oldColor := image[sr][sc]
	if newColor == oldColor {
		return image
	}
	directories := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := list.New() // 记录坐标
	queue.PushBack([]int{sr, sc})
	for queue.Len() > 0 {
		rc := queue.Remove(queue.Front()).([]int)
		image[rc[0]][rc[1]] = newColor // 着色
		for _, v := range directories {
			r, c := rc[0]+v[0], rc[1]+v[1] // 四个方向
			if r >= 0 && r < len(image) && c >= 0 && c < len(image[0]) &&
				newColor != image[r][c] && image[r][c] == oldColor {
				queue.PushBack([]int{r, c})
			}
		}
	}
	return image
}
