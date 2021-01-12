/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package open_the_lock

func openLock(deadends []string, target string) int {
	const initial = "0000"
	if target == initial {
		return 0
	}
	visited := make(map[string]bool)
	for _, v := range deadends {
		visited[v] = true
	}
	if visited[initial] {
		return -1
	}

	start := map[string]bool{initial: true}
	end := map[string]bool{target: true}
	return bfs(start, end, visited, 0) // count代表步数，从0开始
}

// 模拟双向搜索。 count为步数
func bfs(start, end, visited map[string]bool, count int) int {
	if len(start) > len(end) { // 从已经遍历少的那一端开始， 维持两端搜索的数量相当，能明显优化搜索步数
		return bfs(end, start, visited, count)
	}
	if len(start) <= 0 {
		return -1
	}

	nextStatus := make(map[string]bool) //存储start端下一步需要处理的状态
	for s := range start {
		if _, ok := end[s]; ok { // end队列也有，说明从初始到目标状态的一个通路形成了
			return count
		}
		visited[s] = true
		b := []byte(s)
		for i, c := range b {
			for d := -1; d <= 1; d += 2 { // 1 或 -1， b的每一位要加1或减1，以得到下一个状态
				b[i] = byte((int(c-'0')+d+10)%10) + '0'
				next := string(b)
				b[i] = c // 复原状态
				if visited[next] {
					continue
				}
				nextStatus[next] = true
			}
		}
	}
	count++
	return bfs(nextStatus, end, visited, count)
}
