/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package game

/*
174. 地下城游戏
https://leetcode-cn.com/problems/dungeon-game

一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。
我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。

骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。

有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；
其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。

为了尽快到达公主，骑士决定每次只向右或向下移动一步。

编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。

例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。

-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)


说明:

骑士的健康点数没有上限。

任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
*/
/*
动态规划，时空复杂度都是O(m*n)
这个骑士健康值小于1就会死去，从终点公主的房间逆向考虑就行：
如果公主房间健康值dp[m-1][n-1]是负数，那么勇士在进入之前至少要-dungeon[m-1][n-1]+1的健康值；如果是正数，则进入前健康值为1就行
设dp[i][j]代表进入房间（i，j）前最少的健康值
显然，如果dungeon[i][j] >= 0, dp[i][j] = 1;
否则，需要在扣完dungeon[i][j]后剩余下一个房间的最低要求;两个房间，选要求低的;则dp[i][j] = -dungeon[i][j] + min(dp[i+1][j], dp[i][j+1])
综合来看，dp[i][j] = max(1, -dungeon[i][j] + min(dp[i+1][j], dp[i][j+1])；对于边界情况，最后一个房间、最后一行、最后一列稍有不同
*/
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for r := 0; r < m; r++ {
		dp[r] = make([]int, n)
	}
	// 确定边界情况，同时也确定了初始值
	dp[m-1][n-1] = max(1, -dungeon[m-1][n-1]+1) // 最后一个房间
	for c := n - 2; c >= 0; c-- {               // 最后一行
		dp[m-1][c] = max(1, -dungeon[m-1][c]+dp[m-1][c+1])
	}
	for r := m - 2; r >= 0; r-- { // 最后一列
		dp[r][n-1] = max(1, -dungeon[r][n-1]+dp[r+1][n-1])
	}
	// 从右下向起点遍历，计算每个房间的dp值
	for r := m - 2; r >= 0; r-- {
		for c := n - 2; c >= 0; c-- {
			dp[r][c] = max(1, -dungeon[r][c]+min(dp[r+1][c], dp[r][c+1]))
		}
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
