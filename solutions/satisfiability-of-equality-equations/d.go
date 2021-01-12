package satisfiability_of_equality_equations

/*
990. 等式方程的可满足性
https://leetcode-cn.com/problems/satisfiability-of-equality-equations

给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，
并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。

只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

示例 1
输入：["a==b","b!=a"]
输出：false
解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。

示例 2：
输出：["b==a","a==b"]
输入：true
解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。

示例 3：
输入：["a==b","b==c","a==c"]
输出：true

示例 4：
输入：["a==b","b!=c","c==a"]
输出：false

示例 5：
输入：["c==c","b==d","x!=z"]
输出：true

提示：

1 <= equations.length <= 500
equations[i].length == 4
equations[i][0] 和 equations[i][3] 是小写字母
equations[i][1] 要么是 '='，要么是 '!'
equations[i][2] 是 '='
*/
/*
因所有字母都是小写，创建大小为26的并查集
先处理所有形如 x==y 的表达式， 将 == 两端的字母合并
再处理形如 x!=y 的表达式，在并查集中查看x和y是否属于同一个集合，如果是，出现了矛盾，返回false
遍历完所有形如 x!=y 的表达式后也没发现矛盾，返回true
*/
func equationsPossible(equations []string) bool {
	uf := NewUnionFind(26)
	for _, v := range equations {
		if v[1] == '=' {
			uf.Union(int(v[0]-'a'), int(v[3]-'a'))
		}
	}
	for _, v := range equations {
		if v[1] == '!' && uf.Find(int(v[0]-'a')) == uf.Find(int(v[3]-'a')) {
			return false
		}
	}
	return true
}

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}
func (uf UnionFind) Find(x int) int {
	for uf[x] != x {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}
func (uf UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY
}
