/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package campus_bikes_ii

import "math"

/*
在由 2D 网格表示的校园里有 n 位工人（worker）和 m 辆自行车（bike），n <= m。所有工人和自行车的位置都用网格上的 2D 坐标表示。

我们为每一位工人分配一辆专属自行车，使每个工人与其分配到的自行车之间的曼哈顿距离最小化。

p1 和 p2 之间的曼哈顿距离为 Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|。

返回每个工人与分配到的自行车之间的曼哈顿距离的最小可能总和。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/campus-bikes-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
0-1背包的解法：
工人数等价于物品数
自行车数等价于容量
距离等价于价值
*/
func assignBikes(workers [][]int, bikes [][]int) int {
	dp := make(map[int]int, 0)
	dp[0] = 0
	for _, worker := range workers {
		newDp := make(map[int]int, 0)
		for j, bike := range bikes {
			for key, value := range dp {
				if key&(1<<uint(j)) != 0 {
					continue
				}
				lastKey := key | (1 << uint(j))
				newV := value + dist(worker[0], worker[1], bike[0], bike[1])
				if _, ok := newDp[lastKey]; !ok || newDp[lastKey] > newV {
					newDp[lastKey] = newV
				}
			}
		}
		dp = newDp
	}
	min := math.MaxInt32
	for _, v := range dp {
		if v < min {
			min = v
		}
	}
	return min
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func assignBikes1(workers [][]int, bikes [][]int) int {
	n, m := uint8(len(workers)), uint8(len(bikes))
	dp := make([]int, (1<<m)-1)
	var cal func(uint8, uint16) int
	cal = func(workerId uint8, bikeTaken uint16) int {
		if workerId == n {
			return 0
		}
		if dp[bikeTaken] != 0 {
			return dp[bikeTaken]
		}

		t1 := math.MaxInt32
		for i := uint8(0); i < m; i++ {
			if (bikeTaken & (1 << i)) != 0 {
				continue
			}
			w, b := workers[workerId], bikes[i]
			t2 := cal(workerId+1, bikeTaken|(1<<i)) + dist(w[0], w[1], b[0], b[1])
			if t2 < t1 {
				t1 = t2
			}
		}
		dp[bikeTaken] = t1
		return t1
	}
	return cal(0, 0)
}
