package find_mode_in_binary_search_tree

/*
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
中序遍历

可以使用一个哈希表暴力记录每个节点值的个数，这样空间复杂度较高，也没有利用 BST 有序的特性
把 BST 换成有序数组，就会比较简单：
先遍历一次数组，得到最大的元素个数；再遍历一次，统计个数是最大个数的元素即可
也可以只遍历一次

可以借 BST 的中序遍历，将 BST 转化成一个数组后求解，但更省空间的做法是直接在中序遍历过程里统计结果
*/

var pre *TreeNode
var res []int
var count, maxCount int

func findMode(root *TreeNode) []int {
	pre, res = nil, nil
	count, maxCount = 0, 0
	inorder(root)
	return res
}

// 递归式中序遍历代，可读性好
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	update(root)
	inorder(root.Right)
}

/*
morris 中序遍历
可读性没有递归式好
理论上的空间复杂度是常数级，优于递归式的 O(h)栈空间（h是树的高度）；
不过实测， LeetCode 用例两种方法的内存消耗都是 6.2 M
*/
func inorderMorris(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur)
			cur = cur.Right
			continue
		}
		node := cur.Left
		for node.Right != nil && node.Right != cur {
			node = node.Right
		}
		if node.Right == nil {
			node.Right = cur
			cur = cur.Left
		} else {
			node.Right = nil
			update(cur)
			cur = cur.Right
		}
	}
}
func update(cur *TreeNode) {
	if pre == nil || cur.Val == pre.Val {
		count++
	} else {
		count = 1
	}
	if count == maxCount {
		res = append(res, cur.Val)
	} else if count > maxCount {
		maxCount = count
		res = []int{cur.Val}
	}
	pre = cur
}
