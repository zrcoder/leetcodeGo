/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package serialize_and_deserialize_binary_tree

import (
	"bytes"
	"container/list"
	"strconv"
	"strings"
)

/*
297. 二叉树的序列化与反序列化 https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// DFS 先序
func (c *Codec) serialize(root *TreeNode) string {
	buf := bytes.NewBuffer(nil)
	var preorder func(*TreeNode)
	preorder = func(n *TreeNode) {
		if n == nil {
			buf.WriteString("#,")
			return
		}
		buf.WriteString(strconv.Itoa(n.Val))
		buf.WriteString(",")
		preorder(n.Left)
		preorder(n.Right)
	}
	preorder(root)
	return buf.String()
}

func (c *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	index := 0
	var help func() *TreeNode
	help = func() *TreeNode {
		if index == len(nodes) {
			return nil
		}
		val, err := strconv.Atoi(nodes[index])
		index++
		if err != nil { // nodes[index] == "#"
			return nil
		}
		root := &TreeNode{Val: val}
		root.Left = help()
		root.Right = help()
		return root
	}
	return help()
}

// BFS
func (c *Codec) serializeBfs(root *TreeNode) string {
	queue := list.New()
	queue.PushBack(root)
	b := strings.Builder{}
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node == nil {
			b.WriteString("#,")
			continue
		}
		b.WriteString(strconv.Itoa(node.Val))
		b.WriteByte(',')
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
	return b.String()
}

func (c *Codec) deserializeBfs(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := 0
	val, err := strconv.Atoi(values[index])
	if err != nil { // values[0] == "#"
		return nil
	}
	root := &TreeNode{Val: val}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		index++
		if values[index] != "#" {
			val, _ = strconv.Atoi(values[index])
			left := &TreeNode{Val: val}
			node.Left = left
			queue.PushBack(left)
		}
		index++
		if values[index] != "#" {
			val, _ = strconv.Atoi(values[index])
			right := &TreeNode{Val: val}
			node.Right = right
			queue.PushBack(right)
		}
	}
	return root
}
