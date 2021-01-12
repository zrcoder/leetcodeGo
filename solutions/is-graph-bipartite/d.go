/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package is_graph_bipartite

/*
给定一个无向图graph，当这个图为二分图时返回true。

如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。

graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。每个节点都是一个在0到graph.length-1之间的整数。
这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。


示例 1:
输入: [[1,3], [0,2], [1,3], [0,2]]
输出: true
解释:
无向图如下:
0----1
|    |
|    |
3----2
我们可以将节点分成两组: {0, 2} 和 {1, 3}。

示例 2:
输入: [[1,2,3], [0,2], [0,1,3], [0,2]]
输出: false
解释:
无向图如下:
0----1
| \  |
|  \ |
3----2
我们不能将节点分割成两个独立的子集。
注意:

graph 的长度范围为 [1, 100]。
graph[i] 中的元素的范围为 [0, graph.length - 1]。
graph[i] 不会包含 i 或者有重复的值。
图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-graph-bipartite
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
若节点i在集合A中，则与其相连的节点必定不在集合A中（而是在B中）；否则无法构成二分图

开辟一个与graph一样大的数组来记录每个节点的归属
0表示还没有决定归属；1表示归属于A；-1表示归属于B

遍历所有节点(0, 1, ..., n-1)，决定每个节点的归属
节点i如果还没有归属，则标记归属于A；
再递归遍历其相邻的节点，如果没归属则标记归属于B；如果有归属，预期和i的归属相反，否则出现了矛盾，可以断定不是二分图

最外层的循环，节点i如果未标记归属，可以随意标记个分组（以下示例程序统一标成了A）
这是因为前边的节点0,1,...,i-1及其邻居都标记过了，i还没有被标记，说明i和前边的节点不联通；
既然不联通，就无所谓归属于哪个组了
*/
func isBipartite(graph [][]int) bool {
	judge := make([]int, len(graph))

	var mark func(node, set int) bool // 这个内部函数也可以写到和isBipartite并列的位置，不过要增加judge和graph参数
	mark = func(node, set int) bool {
		judge[node] = set
		for _, neighbor := range graph[node] {
			if judge[neighbor] == 0 {
				if ok := mark(neighbor, -set); !ok {
					return false
				}
			} else if judge[neighbor] == set {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(graph); i++ {
		if judge[i] != 0 {
			continue
		}
		if ok := mark(i, 1); !ok {
			return false
		}
	}
	return true
}

/*
也可以用一个栈模拟实现递归
*/
func isBipartite1(graph [][]int) bool {
	judge := make([]int, len(graph))

	var mark func(node, set int) bool
	mark = func(node, set int) bool {
		judge[node] = set
		var stack []int
		stack = append(stack, node)
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, neighbor := range graph[top] {
				if judge[neighbor] == 0 {
					judge[neighbor] = -judge[top]
					stack = append(stack, neighbor)
				} else if judge[neighbor] == judge[top] {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < len(graph); i++ {
		if judge[i] != 0 {
			continue
		}
		if ok := mark(i, 1); !ok {
			return false
		}
	}
	return true
}
