/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package path_sum_iv

/*
666. 路径和 IV https://leetcode-cn.com/problems/path-sum-iv/
对于一棵深度小于 5 的树，可以用一组三位十进制整数来表示。

对于每个整数：

百位上的数字表示这个结点的深度 D，1 <= D <= 4。
十位上的数字表示这个结点在当前层所在的位置 P， 1 <= P <= 8。位置编号与一棵满二叉树的位置编号相同。
个位上的数字表示这个结点的权值 V，0 <= V <= 9。
给定一个包含三位整数的升序数组，表示一棵深度小于 5 的二叉树，请你返回从根到所有叶子结点的路径之和。

样例 1:

输入: [113, 215, 221]
输出: 12
解释:
这棵树形状如下:
    3
   / \
  5   1

路径和 = (3 + 5) + (3 + 1) = 12.


样例 2:

输入: [113, 221]
输出: 4
解释:
这棵树形状如下:
    3
     \
      1

路径和 = (3 + 1) = 4.
*/
/*
方法一： 依据给定信息构建出常规表示的二叉树，再dfs求路径和即可
时空复杂度都是O(n), n是数组长度，也是节点总数

关键在构建：
遍历nums，对于每个节点，分解出其深度d、位置pos和值val，
1.从已经构建的根节点开始，需要向下走d-1步
2.每一步向左还是向右？ 根据pos和当前层节点排满的总数决定向左还是向右走
如果目标所在的层排满节点， 深度为d的层上应有2^(d-1)个节点，
如深度为4时，应有8个节点，根据题目描述，
pos的取值可能是1, 2, 3, 4, 5, 6, 7, 8里的一个
可见当pos<=4时，pos对应节点应该在整棵树左侧，反之在右侧；
即pos<=2^(d-1) / 2 则pos对应节点应该在整棵树左侧， 反之在右侧
3.到了新的一层，层节点的总数减半，pos要考虑大于层节点总数的情况取模
*/
type Node struct {
	Left, Right *Node
	Val         int
}

func pathSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	pathSum := 0
	var dfs func(t *Node)
	dfs = func(t *Node) {
		if t == nil {
			return
		}
		pathSum += t.Val
		if t.Left == nil && t.Right == nil { // t是叶子节点
			result += pathSum
		}
		dfs(t.Left)
		dfs(t.Right)
		pathSum -= t.Val
	}
	root := constructTree(nums)
	dfs(root)
	return result
}

func constructTree(nums []int) *Node {
	root := &Node{Val: nums[0] % 10}
	for i := 1; i < len(nums); i++ {
		depth, pos, val := nums[i]/100, nums[i]/10%10, nums[i]%10
		cur := root
		levelTotal := 1 << uint(depth-1) // 2^(depth-1)
		// 1.从已经构建的根节点开始，需要向下走d-1步
		for i := 0; i < depth-1; i++ {
			// 2.根据pos和当前层节点排满的总数决定向左还是向右走
			if pos <= levelTotal/2 { // 向左走
				if cur.Left == nil {
					cur.Left = &Node{Val: val}
					break
				}
				cur = cur.Left
			} else { // 向右走
				if cur.Right == nil {
					cur.Right = &Node{Val: val}
					break
				}
				cur = cur.Right
			}
			// 3.到了新的一层，层节点的总数减半，pos要考虑大于层节点总数的情况取模
			levelTotal /= 2
			pos = (pos-1)%levelTotal + 1
		}
	}
	return root
}

/*
方法二：
不用真的转换成常规表示的树，而是转化成哈希表或数组表示的树
时空复杂度依然都是O(n)

每个节点的前两位（深度+位置）唯一标识了该节点，且根据这个标识信息能很容易地推出左右孩子的标识（深度+位置）
假设某个节点的深度和位置分别是depth和pos，则其孩子结点的标识是：
left = 10 * (depth + 1) + 2 * pos - 1，
right = left + 1。
*/

// 方法二之哈希表模拟树
func pathSum1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nodes := make(map[int]int, len(nums))
	for _, v := range nums {
		nodes[v/10] = v % 10
	}
	result := 0
	pathSum := 0
	var dfs func(node int)
	dfs = func(node int) {
		val, ok := nodes[node]
		if !ok {
			return
		}
		pathSum += val
		depth, pos := node/10, node%10
		left := 10*(depth+1) + pos*2 - 1
		right := left + 1
		_, ok1 := nodes[left]
		_, ok2 := nodes[right]
		if !ok1 && !ok2 { // node是叶子节点
			result += pathSum
		}
		dfs(left)
		dfs(right)
		pathSum -= val
	}
	dfs(nums[0] / 10)
	return result
}

// 方法二之数组模拟树
/*
题目限定深度小于5层，那么最大的元素标识是，深度为4，位置8，即48;同时最小的节点标识是11
可以用一个大小48-11+1=38的数组来代替哈希表，这样读写更快
又因为给定树的值都是正数，初始化可以让数组所有元素为-1，代表空节点
*/
func pathSum2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	const (
		minId         = 11
		maxId         = 48
		maxLevelMinId = 41
		nilNode       = -1
	)
	nodes := make([]int, maxId-minId+1)
	for i := range nodes {
		nodes[i] = nilNode
	}
	for _, v := range nums {
		nodes[v/10-minId] = v % 10
	}
	result := 0
	pathSum := 0
	var dfs func(node int)
	dfs = func(node int) {
		val := nodes[node-minId]
		if val == nilNode {
			return
		}
		pathSum += val
		if node >= maxLevelMinId { // node是叶子节点
			result += pathSum
		} else {
			depth, pos := node/10, node%10
			left := 10*(depth+1) + pos*2 - 1
			right := left + 1
			if nodes[left-minId] == nilNode && nodes[right-minId] == nilNode { // node是叶子节点
				result += pathSum
			}
			dfs(left)
			dfs(right)
		}
		pathSum -= val
	}
	dfs(nums[0] / 10)
	return result
}
