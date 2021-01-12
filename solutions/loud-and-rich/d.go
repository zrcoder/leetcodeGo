package rich

import "container/list"

/*
851. 喧闹和富有
https://leetcode-cn.com/problems/loud-and-rich

在一组 N 个人（编号为 0, 1, 2, ..., N-1）中，每个人都有不同数目的钱，以及不同程度的安静（quietness）。

为了方便起见，我们将编号为 x 的人简称为 "person x "。

如果能够肯定 person x 比 person y 更有钱的话，我们会说 richer[i] = [x, y] 。注意 richer 可能只是有效观察的一个子集。

另外，如果 person x 的安静程度为 q ，我们会说 quiet[x] = q 。

现在，返回答案 answer ，其中 answer[x] = y 的前提是，在所有拥有的钱不少于 person x 的人中，person y 是最安静的人（也就是安静值 quiet[y] 最小的人）。

示例：
输入：richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
输出：[5,5,2,5,4,5,6,7]
解释：
answer[0] = 5，
person 5 比 person 3 有更多的钱，person 3 比 person 1 有更多的钱，person 1 比 person 0 有更多的钱。
唯一较为安静（有较低的安静值 quiet[x]）的人是 person 7，
但是目前还不清楚他是否比 person 0 更有钱。

answer[7] = 7，
在所有拥有的钱肯定不少于 person 7 的人中(这可能包括 person 3，4，5，6 以及 7)，
最安静(有较低安静值 quiet[x])的人是 person 7。

其他的答案也可以用类似的推理来解释。
提示：

1 <= quiet.length = N <= 500
0 <= quiet[i] < N，所有 quiet[i] 都不相同。
0 <= richer.length <= N * (N-1) / 2
0 <= richer[i][j] < N
richer[i][0] != richer[i][1]
richer[i] 都是不同的。
对 richer 的观察在逻辑上是一致的。
*/
// BFS
// 要注意可能存在重复访问某个人的情况，比如 1，2 比 0富有，同时5比1和2都富有，那么在确定比0富有的人里谁更低调时，5会被重复访问
// 用一个viseted数组来记录某个人被访问过
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	richers := make([][]int, n)
	for _, v := range richer {
		x, y := v[0], v[1]
		richers[y] = append(richers[y], x)
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		q := list.New()
		q.PushBack(i)
		visited := make([]bool, n)
		visited[i] = true
		result[i] = i // 有可能没人比自己更有钱或更低调，自己也要参与比较
		for q.Len() > 0 {
			person := q.Remove(q.Front()).(int)
			if quiet[person] < quiet[result[i]] {
				result[i] = person
			}
			for _, v := range richers[person] {
				if visited[v] {
					continue
				}
				q.PushBack(v)
				visited[v] = true
			}
		}
	}
	return result
}

// dfs
func loudAndRichDfs(richer [][]int, quiet []int) []int {
	n := len(quiet)
	richers := make([][]int, n)
	for _, v := range richer {
		x, y := v[0], v[1]
		richers[y] = append(richers[y], x)
	}

	result := make([]int, len(quiet))
	for i := range result {
		result[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if result[i] != -1 {
			return result[i]
		}
		result[i] = i
		for _, v := range richers[i] {
			r := dfs(v)
			if quiet[r] < quiet[result[i]] {
				result[i] = r
			}
		}
		return result[i]
	}
	for i := 0; i < n; i++ {
		_ = dfs(i)
	}
	return result
}
