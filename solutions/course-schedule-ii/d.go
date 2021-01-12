/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package course_schedule_ii

import "container/list"

// bfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	degree := make([]int, numCourses)
	nexts := make([][]int, numCourses)
	for _, req := range prerequisites {
		degree[req[0]]++
		nexts[req[1]] = append(nexts[req[1]], req[0])
	}
	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue.PushBack(i)
		}
	}
	res := make([]int, 0, numCourses)
	for queue.Len() > 0 {
		course := queue.Remove(queue.Front()).(int)
		res = append(res, course)
		for _, next := range nexts[course] {
			degree[next]--
			if degree[next] == 0 {
				queue.PushBack(next)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return nil
}

// dfs
func findOrder11(numCourses int, prerequisites [][]int) []int {
	dependency := make([][]int, numCourses)
	for _, req := range prerequisites {
		dependency[req[0]] = append(dependency[req[0]], req[1])
	}
	flags := make([]int, numCourses)
	var result []int
	var dfs func(course int) bool
	dfs = func(course int) bool {
		// 1.节点已经被访问过
		if flags[course] == 1 { // 之前 dfs 访问该点无环
			return true
		}
		if flags[course] == -1 { // 之前 dfs 访问该点有环
			return false
		}
		// 2.节点还没有被访问过，先假设有环
		flags[course] = -1
		for _, neighbor := range dependency[course] {
			if !dfs(neighbor) { // 真的有环
				return false
			}
		}
		// 所有相邻节点都 dfs 搜索过了，可以确定当次 dfs 无环，加入结果
		flags[course] = 1
		result = append(result, course)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return nil
		}
	}
	return result
}
