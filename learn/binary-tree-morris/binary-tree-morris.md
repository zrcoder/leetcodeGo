## Morris 遍历二叉树
这个算法是 KMP 算法作者之一 Morris 提出的，空间复杂度非常优秀，同时和常规实现的时间复杂度相当。

利用了树本身的空间，在遍历过程中会修改树结构，但是遍历完能恢复。

### 中序遍历
不失一般性，先给出如下二叉树定义：
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```
要保证整个算法只使用常数级的额外空间，难点在于：遍历到子节点时怎么返回父节点。<br>
Morris 算法用到了线索二叉树（threaded binary tree）的概念。<br>
中间过程利用叶子节点的右孩子空指针来指向其后继节点。<br>

具体过程：
```
用变量 cur 代表当前节点，一开始其值为根节点
1.如果 cur 左孩子为空，访问 cur 并 将 cur 指向其右孩子
2.否则，找到 cur 的前驱节点 node，实际为左子树最右叶子节点
    2.1 如果 node 的右指针为空，这说明是第一次来找前驱节点，现在将 node 的右指针指向 cur，cur 指向自己的左孩子。
    2.2 由于 2.1 的原因，node 的右指针可能是cur，如果是 cur，说明 cur 的左子树已经遍历完成，将 node 右指针重新置为 nil以恢复树结果，访问当前节点，cur 指向自己的右孩子。
```
见如下例子：

![](morris-inorder.png)

示例代码如下：
```go
func inorderTraversalMorris(root *TreeNode) {
	cur := root
	var node *TreeNode
	for cur != nil {
		if cur.Left == nil {
			visit(cur)
			cur = cur.Right
			continue
		}
		// 找 cur 的前驱
		node = cur.Left
		for node.Right != nil && node.Right != cur {
			node = node.Right
		}
		if node.Right == nil { // 还没线索化，建立线索
			node.Right = cur
			cur = cur.Left
		} else { // 已经线索化，访问节点并删除线索以恢复树的结构
			node.Right = nil
			visit(cur)
			cur = cur.Right
		}
	}
}
```
复杂度分析
```
空间复杂度为 O(1), 只用了两个指针 cur 和 node；
时间复杂度是 O(n), n 指节点总数
虽然找 cur 的前驱节点是一个循环，但总观整个遍历过程，每个节点最多被访问 3 次
```
### 前序遍历
和中序遍历非常类似，只是访问 cur 的时机稍有不同，直接看代码
```go
func preorderTraversalMorris(root *TreeNode) {
	cur := root
	var node *TreeNode
	for cur != nil {
		if cur.Left == nil {
			visit(cur)
			cur = cur.Right
			continue
		}
		// 找 cur 的前驱
		node = cur.Left
		for node.Right != nil && node.Right != cur {
			node = node.Right
		}
		if node.Right == nil { // 还没线索化，建立线索
			visit(cur)
			node.Right = cur
			cur = cur.Left
		} else { // 已经线索化，访问节点并删除线索以恢复树的结构
			node.Right = nil
			cur = cur.Right
		}
	}
}
```
可以看如下图示对比理解：

![](morris-preorder.png)

### 后序遍历
稍微复杂一些。<br>
要建立一个哨兵节点 dummy，令其左孩子是 root。<br>
并且还需要一个子过程，就是倒序输出 cur 左孩子到 cur 的前驱节点之间路径上所有节点。<br>
```go
func postorderTraversalMorris(root *TreeNode) {
	dummy := &TreeNode{Left: root}
	cur := dummy
	var node *TreeNode
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		// 找 cur 的前驱
		node = cur.Left
		for node.Right != nil && node.Right != cur {
			node = node.Right
		}
		if node.Right == nil { // 还没线索化，建立线索
			node.Right = cur
			cur = cur.Left
		} else { // 已经线索化，访问节点并删除线索以恢复树的结构
			node.Right = nil
			visitPath(cur.Left)
			cur = cur.Right
		}
	}
	dummy.Left = nil
}

func visitPath(node *TreeNode) {
	end := reversePath(node)
	for p := end; p != nil; p = p.Right {
		visit(p)
	}
	_ = reversePath(end)
}

func reversePath(node *TreeNode) * TreeNode {
	var prev *TreeNode
	for node != nil {
		prev, node, node.Right = node, node.Right, prev
	}
	return prev
}
```
图示：

![](morris-postorder.png)
### 参考
https://www.cnblogs.com/AnnieKim/archive/2013/06/15/MorrisTraversal.html <br>
https://zhuanlan.zhihu.com/p/101321696
