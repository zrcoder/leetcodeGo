package binary_tree_cameras

import "math"

/*
968. 监控二叉树
https://leetcode-cn.com/problems/binary-tree-cameras

给定一个二叉树，我们在树的节点上安装摄像头。
节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
计算监控树的所有节点所需的最小摄像头数量。

示例 1：
输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。

示例 2：
输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。

提示：
给定树的节点数的范围是 [1, 1000]。
每个节点的值都是 0。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
自上而下递归决策，初步尝试
类似 337. 打家劫舍 III，更复杂些

对于一个节点，仅用是否安装了相机这一个状态没法得到结果，还需加一个状态：是否被监控
这两个状态会有重合，安装了相机意味着同时被监控了
详见如下代码

不过这个写法的战绩是：160 / 170 个通过测试用例，后边超时了
*/
func minCameraCover0(root *TreeNode) int {
	var help func(*TreeNode, bool, bool) int
	// placeCam，是否在 node 处安装相机；
	// watched，node 是否被父节点或自身监控(递归过程是自上而下，对于当前节点，只知道父节点或自身是否监控自己，并不知道子节点的情况)
	help = func(node *TreeNode, placeCam, watched bool) int {
		if node == nil {
			if placeCam {
				return math.MaxInt32
			}
			return 0
		}

		leftPlaceWatch := help(node.Left, true, true)
		rightPlaceWatch := help(node.Right, true, true)

		if placeCam {
			leftNotPlaceWatch := help(node.Left, false, true)
			rightNotPlaceWatch := help(node.Right, false, true)
			return 1 + min(
				leftNotPlaceWatch+rightNotPlaceWatch, // 两个子节点都不安装相机
				leftPlaceWatch+rightNotPlaceWatch,    // 仅左子节点安装相机
				leftNotPlaceWatch+rightPlaceWatch)    // 仅右子节点安装相机
			// 两个子节点都装相机的情况不用考虑
		}
		leftNotPlaceNotWatch := help(node.Left, false, false)
		rightNotPlaceNotWatch := help(node.Right, false, false)
		res := min(
			leftPlaceWatch+rightPlaceWatch,       // 两个子节点都安装相机
			leftPlaceWatch+rightNotPlaceNotWatch, // 左装右不装
			leftNotPlaceNotWatch+rightPlaceWatch) // 右装左不装
		if watched {
			res = min(res, leftNotPlaceNotWatch+rightNotPlaceNotWatch) // 左右都不装，当前节点是被其父节点监控的
		}
		return res

	}
	return min(help(root, true, true), help(root, false, false))
}

/*
自上而下决策，优化

初步尝试中对左右子树反复调用 help，传的参数许多是重复的

原本分三次调用求一个节点三种状态下的结果，可以合并到一次调用里做，只需要把 help 函数的入参减少，返回值个数增大~

对于当前节点，可以返回下边三种情况下的结果
hasCam:					有相机
noCamWatchedByParent:	没相机，被父节点监控
noCamWatchedBySons:		没相机，被子节点监控

时间复杂度 O(n)，其中 n 是节点总数
空间复杂度 O(h)，h 是树高，递归栈的大小
*/
func minCameraCover(root *TreeNode) int {
	// 返回三种情况下分别所需最少相机数
	var help func(node *TreeNode) (int, int, int)
	help = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return math.MaxInt32, 0, 0
		}
		lHasCam, lNoCamWatchedByParent, lNoCamWatchedBySons := help(node.Left)
		rHasCam, rNoCamWatchedByParent, rNoCamWatchedBySons := help(node.Right)

		hasCam := 1 + min(
			lNoCamWatchedByParent+rNoCamWatchedByParent,
			lHasCam+rNoCamWatchedByParent,
			lNoCamWatchedByParent+rHasCam)

		noCamWatchedByParent := min(
			lHasCam+rHasCam,
			lHasCam+rNoCamWatchedBySons,
			lNoCamWatchedBySons+rHasCam,
			lNoCamWatchedBySons+rNoCamWatchedBySons)

		noCamWatchedBySons := min(
			lHasCam+rHasCam,
			lHasCam+rNoCamWatchedBySons,
			lNoCamWatchedBySons+rHasCam)
		return hasCam, noCamWatchedByParent, noCamWatchedBySons
	}
	hasCam, _, noCamWatchedBySons := help(root)
	return min(hasCam, noCamWatchedBySons)
}

func min(s ...int) int {
	r := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < r {
			r = s[i]
		}
	}
	return r
}

/*
自下而上贪心决策

先考虑一棵满二叉树，因为越向下节点越多，为了尽量少地安装相机，可以使用贪心策略从下向上层层决策。
叶子节点不需要装相机				——记为状态0
向上一层，叶子节点的父节点需要装相机	——记为状态1；如果不是满二叉树，这里只需要有一个子节点是叶子节点即可
再上一层，子节点是状态1，不需要装相机	——记为状态2；如果不是满二叉树，这里只需要有一个子节点是状态1即可
再上一层，节点所有子节点都是状态2，当前节点回归到状态0
...
如下所示满二叉树各个节点状态：

       1
     0   0
   2       2
 1   1   1   1
0 0 0 0 0 0 0 0

如果不是满二叉树，以上贪心策略依然正确，不会有越向下节点越少的树，最极端情况是树退化为一条链表，每一层都只有一个节点。

最后有个情况要特别考虑，即 root 最后是状态0；这意味着root本来就是一个叶子节点，或者如下所示
     0
   2
 1   1
0 0 0 0
显然这种情况，需要在 root 节点安装相机，结果+1即可

时空复杂度同上
*/

var res int

func minCameraCover1(root *TreeNode) int {
	res = 0
	if dfs(root) == 0 {
		res++
	}
	return res
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 2
	}
	left, right := dfs(node.Left), dfs(node.Right)
	if left == 0 || right == 0 {
		res++
		return 1
	}
	if left == 1 || right == 1 {
		return 2
	}
	return 0
}
