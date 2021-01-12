package recover_a_tree_from_preorder_traversal

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

/*
1028. 从先序遍历还原二叉树
我们从二叉树的根节点 root 开始进行深度优先搜索。

在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。
（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。

如果节点只有一个子节点，那么保证该子节点为左子节点。

给出遍历输出 S，还原树并返回其根节点 root。


示例 1：


输入："1-2--3--4-5--6--7"
输出：[1,2,5,3,4,6,7]

示例 2：

输入："1-2--3---4-5--6---7"
输出：[1,2,5,3,null,6,null,4,null,7]

示例 3：

输入："1-401--349---90--88"
输出：[1,401,null,349,88,90]


提示：
原始树中的节点数介于 1 和 1000 之间。
每个节点的值介于 1 和 10 ^ 9 之间。
*/
/*
非常精妙的一个问题
可以先复习下[113]路径和相关问题

遍历s，可以根据约束规范统计出一个个节点的深度和值
遍历过程中，对于当前节点，怎么放到已经构建的树的合适的位置？

这里需要逆向思考。
对于当前节点N和上一个遍历过的节点M，只有两种情况：
1. N是M的左孩子； 2.N是根节点到M的路径上某个节点X的右孩子——且可以确定X不是M
这两种情况又有共性，可以发现1中N的深度肯定是M的深度加1；2中也类似，N的深度是X的深度加1
这样根据深度就可以大体确定当前节点N的父节点了
可以再借助一个切片path保存路径，就能精确定位了：
如果path的大小正好是N的深度depth， 这意味着情况1，path最后一个元素是N的父节点M
否则只可能N的深度小于path的大小，对应情况2，path的第depth-1个元素即节点X；同时X之后的元素都从path移除
*/
func recoverFromPreorder(s string) *TreeNode {
	var path []*TreeNode
	for i := 0; i < len(s); {
		depth := 0
		for i < len(s) && s[i] == '-' {
			depth++
			i++
		}
		value := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			value = value*10 + int(s[i]-'0')
			i++
		}
		node := &TreeNode{Val: value}
		if len(path) == 0 {
			path = append(path, node)
			continue
		}
		parent := path[depth-1]
		if depth == len(path) {
			parent.Left = node
		} else if depth < len(path) {
			parent.Right = node
			path = path[:depth]
		}
		path = append(path, node)
	}
	if len(path) == 0 {
		return nil
	}
	return path[0]
}
