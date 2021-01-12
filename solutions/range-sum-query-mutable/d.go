package mutable

/*
307. 区域和检索 - 数组可修改
https://leetcode-cn.com/problems/range-sum-query-mutable

给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。

示例:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
说明:

数组仅可以在 update 函数下进行修改。
你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。
*/
/*
线段树，节点版

1.构建一棵满二叉树，叶子节点存储原始nums
2.根节点的索引是1
3.左子树节点的索引为偶数  右子树节点的索引为奇数
4.叶子节点的索引范围[n,2n-1]
*/
type NumArray1 struct {
	root *Node
}

func Constructor1(nums []int) NumArray1 {
	return NumArray1{root: buildTree(nums)}
}
func (na *NumArray1) Update1(i int, val int) {
	na.root.Update(i, val)
}
func (na *NumArray1) SumRange1(i int, j int) int {
	return na.root.SumRange(i, j)
}

type Node struct {
	left, right     *Node
	start, end, sum int
}

func buildTree(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	var help func(i, j int) *Node
	help = func(i, j int) *Node {
		if i == j {
			return &Node{start: i, end: i, sum: nums[i]}
		}
		mid := (i + j) / 2
		left := help(i, mid)
		right := help(mid+1, j)
		return &Node{
			left:  left,
			right: right,
			start: i,
			end:   j,
			sum:   left.sum + right.sum,
		}
	}
	return help(0, len(nums)-1)
}

func (n *Node) Update(i, val int) {
	if i == n.start && i == n.end {
		n.sum = val
		return
	}
	mid := (n.start + n.end) / 2
	if i <= mid {
		n.left.Update(i, val)
	} else {
		n.right.Update(i, val)
	}
	n.sum = n.left.sum + n.right.sum
}

func (n *Node) SumRange(i, j int) int {
	if i == n.start && j == n.end {
		return n.sum
	}
	mid := (n.start + n.end) / 2
	if j <= mid {
		return n.left.SumRange(i, j)
	}
	if i > mid {
		return n.right.SumRange(i, j)
	}
	return n.left.SumRange(i, mid) + n.right.SumRange(mid+1, j)
}

/*
数组版线段树
*/
type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{tree: nil}
	}
	n := len(nums)
	tree := make([]int, 2*n)
	for i, v := range nums {
		tree[i+n] = v
	}
	for i := n - 1; i >= 0; i-- {
		tree[i] = tree[2*i] + tree[2*i+1]
	}
	return NumArray{tree: tree, n: n}
}
func (na *NumArray) Update(i int, val int) {
	i += na.n
	na.tree[i] = val
	for i > 0 {
		left, right := i, i
		if i%2 == 0 {
			right = i + 1
		} else {
			left = i - 1
		}
		na.tree[i/2] = na.tree[left] + na.tree[right]
		i /= 2
	}
}
func (na *NumArray) SumRange(i int, j int) int {
	i += na.n
	j += na.n
	sum := 0
	for i <= j {
		//  线段树左子节点都是偶数下标2n，右子节点都是奇数下标2n+1.
		if i%2 == 1 { // 左指针指向了一个右端点
			sum += na.tree[i]
			i++
		}
		if j%2 == 0 { // 右指针指向了一个左端点
			sum += na.tree[j]
			j--
		}
		i /= 2
		j /= 2
	}
	return sum
}
