package uniquebst

/*
96. 不同的二叉搜索树
https://leetcode-cn.com/problems/unique-binary-search-trees

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

/*
对于数字1，2，3，。。。，n，每个数字都可以作为根节点构造一棵BST
假设f(n)表示n个节点能构建成的BST的个数，g(i, n)表示以i为根节点能构造的BST的个数（1<=i<=n）
则
f(n) = g(1,n) + g(2,n) + ... + g(n, n)	(1)
g(i, n) = f(i-1) * f(n-i)				(2)
*/
// 以下解法超时
func numTrees1(n int) int {
	if n <= 1 {
		return 1
	}
	r := 0
	for i := 1; i <= n; i++ {
		r += numTrees1(i-1) * numTrees1(n-i)
	}
	return r
}

// 给以上解法加上备忘录, 时空复杂度同下边dp的解法
func numTrees2(n int) int {
	if n <= 1 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 1
	var help func(int) int
	help = func(n int) int {
		if memo[n] > 0 {
			return memo[n]
		}
		for i := 1; i <= n; i++ {
			memo[n] += help(i-1) * help(n-i)
		}
		return memo[n]
	}
	return help(n)
}

/*
基于以上解法进一步分析：
(2)代入(1)得到：
f(n) = f(0)*f(n-1) + f(1) * f(n-2) + ... + f(n-1) * f(0)
可以自底向上dp

时间复杂度O(n^2), 空间复杂度O(n)
*/
func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for N := 2; N <= n; N++ {
		for j := 1; j <= N; j++ {
			dp[N] += dp[j-1] * dp[N-j]
		}
	}
	return dp[n]
}

// 卡塔兰数列（catalan），见百度百科：https://baike.baidu.com/item/catalan/7605685?fr=aladdin，
// f(n+1) = f(n) * 2 * (2*i + 1) / (i + 2),  f(0) = 1
// 有递推公式，可以将时间复杂度将为O(n), 空间复杂度将为常数级

func numTrees0(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result = result * 2 * (2*i + 1) / (i + 2)
	}
	return result
}

/*
95. 不同的二叉搜索树 II
https://leetcode-cn.com/problems/unique-binary-search-trees-ii

给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。

示例：
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return help(1, n)
}

func help(low, hi int) []*TreeNode {
	if low > hi {
		return []*TreeNode{nil} // nil 树加入结果，方便后边循环
	}
	var result []*TreeNode
	for i := low; i <= hi; i++ { // i作为根节点
		lefts := help(low, i-1)
		rights := help(i+1, hi)
		for _, left := range lefts {
			for _, right := range rights {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				result = append(result, root)
			}
		}
	}
	return result
}
