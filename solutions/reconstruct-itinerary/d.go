package itinerary

import "sort"

/*
332. 重新安排行程
https://leetcode-cn.com/problems/reconstruct-itinerary

给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。
所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。

说明:

如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
所有的机场都用三个大写字母表示（机场代码）。
假定所有机票至少存在一种合理的行程。

示例 1:
输入: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出: ["JFK", "MUC", "LHR", "SFO", "SJC"]

示例 2:
输入: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
解释: 另一种有效的行程是["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
*/
/*
主要难点在于图有环
这是数学里的欧拉“七桥问题”，即“一笔画”问题；类似问题：753. 破解保险箱

对于当前节点 cur， 可以 dfs 遍历每一个邻接节点，按照邻接节点升序做 dfs，最后将 cur 加入结果 res
这样最终结果就是 res 的逆序
*/
func findItinerary(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for len(m[curr]) > 0 {
			min := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(min)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		j := len(res) - 1 - i
		res[i], res[j] = res[j], res[i]
	}
	return res
}
