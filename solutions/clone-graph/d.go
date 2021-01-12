/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package clone_graph

import "container/list"

/*
133. 克隆图
https://leetcode-cn.com/problems/clone-graph

给你无向 连通 图中一个结点的引用，请你返回该图的 深拷贝（克隆）。

图中的每个结点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

class Node {
    public int val;
    public List<Node> neighbors;
}


测试用例格式：
简单起见，每个结点的值都和它的索引相同。例如，第一个结点值为 1（val = 1），第二个结点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中结点的邻居集。
给定结点将始终是图中的第一个结点（值为 1）。你必须将 给定结点的拷贝 作为对克隆图的引用返回。

示例 1：
输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
输出：[[2,4],[1,3],[2,4],[1,3]]
解释：
图中有 4 个结点。
结点 1 的值是 1，它有两个邻居：结点 2 和 4 。
结点 2 的值是 2，它有两个邻居：结点 1 和 3 。
结点 3 的值是 3，它有两个邻居：结点 2 和 4 。
结点 4 的值是 4，它有两个邻居：结点 1 和 3 。
示例 2：
输入：adjList = [[]]
输出：[[]]
解释：输入包含一个空列表。该图仅仅只有一个值为 1 的结点，它没有任何邻居。
示例 3：
输入：adjList = []
输出：[]
解释：这个图是空的，它不含任何结点。
示例 4：
输入：adjList = [[2],[1]]
输出：[[2],[1]]
提示：
结点数不超过 100 。
每个结点值 Node.val 都是唯一的，1 <= Node.val <= 100。
无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
由于图是无向的，如果结点 p 是结点 q 的邻居，那么结点 q 也必须是结点 p 的邻居。
图是连通图，你可以从给定结点访问到所有结点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/clone-graph
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Node struct {
	Val       int
	Neighbors []*Node
}

/*
把图看成一棵树会比较好理解
要注意和树不同的是，孩子结点可能反过来成为父结点，递归或遍历会形成无限循环
需要用额外的数据结构如哈希表来记录结点是否访问过
直观的做法是哈希表定义为键为*Node值为bool，这样不太好处理；想想，改成值为原图中结点，键为新图中结点，会比较好。
*/
/*
递归DFS

假设所有结点个数为n，
时间复杂度O(n)，每个结点处理一次，栈调用时间复杂度O(H),H为图的最大深度，综合复杂度O(n)
空间复杂度O(n)，哈希表需要O(n)，栈需要O(H)
*/
func cloneGraph1(node *Node) *Node {
	return help(node, make(map[*Node]*Node, 0))
}

func help(node *Node, seen map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if seen[node] != nil {
		return seen[node]
	}
	cloned := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
	seen[node] = cloned
	for i, v := range node.Neighbors {
		cloned.Neighbors[i] = help(v, seen)
	}
	return cloned
}

/*
迭代BFS，存放临时结点的容器可随意选用，这里选list
时间复杂度O(n)，每个结点处理一次
空间复杂度O(n)，哈希表需要O(n), BFS使用的容器需要O(W)，其中W是图的宽度， 综合复杂度O(n)
*/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(node)
	seen := make(map[*Node]*Node, 0)
	seen[node] = &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(*Node)
		for i, v := range n.Neighbors {
			if _, ok := seen[v]; !ok {
				queue.PushBack(v)
				seen[v] = &Node{Val: v.Val, Neighbors: make([]*Node, len(v.Neighbors))}
			}
			seen[n].Neighbors[i] = seen[v]
		}
	}
	return seen[node]
}
