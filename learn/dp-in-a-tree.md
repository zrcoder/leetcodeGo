## 问题
LeetCode 上有两个问题，挺有意思：

[337. 打家劫舍 III](https://leetcode-cn.com/problems/house-robber-iii)

```
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.

示例 2:

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
```

[968. 监控二叉树](https://leetcode-cn.com/problems/binary-tree-cameras)

```
给定一个二叉树，我们在树的节点上安装摄像头。
节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
计算监控树的所有节点所需的最小摄像头数量。

示例 1：

     0
    / 
   📸   
  / \  
 0   0   

输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。

示例 2：

         0
        / 
       📸   
      / 
     0
    /
   📸
     \
      0

输入：[0,0,null,0,null,0,null,null,0]
输出：2
解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。

提示：
给定树的节点数的范围是 [1, 1000]。
每个节点的值都是 0。
```

## 分析

### 树上打劫
让我们从简单一点的第一个问题开始

从根节点开始，递归地做动态规划——这是我第一次抛开数组做动态规划~ <br>
对于当前节点，有两个情况：选或不选（偷或不偷），当然当前节点做了选择后会影响临近节点。 <br>
因为是从上到下，可以认为当前节点的选择只影响其孩子节点： <br>
如果选择当前节点，那么它的孩子节点就不能选 <br>
如果不选择当前节点，那么选不选它的孩子节点都行 <br>
第一版代码如下：

```go
func robTree1(root *TreeNode) int {
	var dfs func(node *TreeNode, selected bool) int
	dfs = func(node *TreeNode, selected bool) int {
		if node == nil {
			return 0
		}
		lNotSelected := dfs(node.Left, false)
		rNotSelected := dfs(node.Right, false)
		if selected {
			return node.Val + lNotSelected + rNotSelected
		}
		lSelected := dfs(node.Left, true)
		rSelected := dfs(node.Right, true)
		return max(lSelected, lNotSelected) + max(rSelected, rNotSelected)
	}
	return max(dfs(root, true), dfs(root, false))
}
```

这就是在树上边做了动态规划。 <br>
不过当前的写法会在 LeetCode 最后一个用例超时。 <br>
有一个轻巧的改进，将 dfs 函数多个参数改成多个返回值：

```go
func robTree(root *TreeNode) int {
	return max(dfs(root))
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lSelected, lNotSelected := dfs(node.Left)
	rSelected, rNotSelected := dfs(node.Right)
	selected := node.Val + lNotSelected + rNotSelected
	notSelected := max(lSelected, lNotSelected) + max(rSelected, rNotSelected)
	return selected, notSelected
}
```

这样就通过了所有用例。 <br>
细想想，一开始的写法有比较多的重复计算，对于同一个节点，会调用两次 dfs 函数。 <br>
但是后边的写法没有重复计算，对同一个节点只调用了一次 dfs 函数。 <br>
（也因为这样， 无需加备忘录优化）。 <br>

不但性能提升了，代码也更精简了。

时间复杂度 `O(n)`, 空间复杂度 `O(h)`, `n` 是节点总数，所有节点都遍历了； `h` 是 树的高度，递归栈的大小

[参考完整实现](../solutions/house-robber/d.go)

### 树上监控
这个问题稍微复杂一点。 <br>
对于一个节点，仅用是否安装了相机这一个状态没法得到结果，还需加一个状态：是否被监控。 <br>
这两个状态会有重合，安装了相机意味着同时被监控了；没装相机，也有可能被监控。 <br>

同样有第一版代码：

```go
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
				leftNotPlaceWatch+rightPlaceWatch) // 仅右子节点安装相机
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
```

果然，这个写法的战绩是：160 / 170 个通过测试用例，后边超时了。 <br>

同样修改多个函数入参为多个返回值

```
对于当前节点，可以返回下边三种情况下的结果
hasCam:					有相机
noCamWatchedByParent:	没相机，被父节点监控
noCamWatchedBySons:		没相机，被子节点监控
```

```go
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
```

AC了

时空复杂度都同树上打劫的那个问题。

[参考完整实现](../solutions/binary-tree-cameras/d.go)

## 小结
整体还是动态规划的思想，只是具体实现是在树上。 <br>
多状态的递归，将状态写到返回值，优于写到入参里，无论从可读性还是从性能。 <br>

## 延申
[[834] 树中距离之和](../solutions/sum-of-distances-in-tree/d.go)

