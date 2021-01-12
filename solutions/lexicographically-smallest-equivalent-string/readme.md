## [1061. 按字典序排列最小的等效字符串](https://leetcode-cn.com/problems/lexicographically-smallest-equivalent-string)
```text
给出长度相同的两个字符串：A 和 B，其中 A[i] 和 B[i] 是一组等价字符。
举个例子，如果 A = "abc" 且 B = "cde"，那么就有 'a' == 'c', 'b' == 'd', 'c' == 'e'。

等价字符遵循任何等价关系的一般规则：
自反性：'a' == 'a'
对称性：'a' == 'b' 则必定有 'b' == 'a'
传递性：'a' == 'b' 且 'b' == 'c' 就表明 'a' == 'c'
例如，A 和 B 的等价信息和之前的例子一样，那么 S = "eed", "acd" 或 "aab"，这三个字符串都是等价的，而 "aab" 是 S 的按字典序最小的等价字符串

利用 A 和 B 的等价信息，找出并返回 S 的按字典序排列最小的等价字符串。

示例 1：
输入：A = "parker", B = "morris", S = "parser"
输出："makkek"
解释：根据 A 和 B 中的等价信息，我们可以将这些字符分为 [m,p], [a,o], [k,r,s], [e,i] 共 4 组。
每组中的字符都是等价的，并按字典序排列。所以答案是 "makkek"。

示例 2：
输入：A = "hello", B = "world", S = "hold"
输出："hdld"
解释：根据 A 和 B 中的等价信息，我们可以将这些字符分为 [h,w], [d,e,o], [l,r] 共 3 组。
所以只有 S 中的第二个字符 'o' 变成 'd'，最后答案为 "hdld"。

示例 3：
输入：A = "leetcode", B = "programs", S = "sourcecode"
输出："aauaaaaada"
解释：我们可以把 A 和 B 中的等价字符分为 [a,o,e,r,s,c], [l,p], [g,t] 和 [d,m] 共 4 组，
因此 S 中除了 'u' 和 'd' 之外的所有字母都转化成了 'a'，最后答案为 "aauaaaaada"。

提示：

字符串 A，B 和 S 仅有从 'a' 到 'z' 的小写英文字母组成。
字符串 A，B 和 S 的长度在 1 到 1000 之间。
字符串 A 和 B 长度相同。
```
## 分析
题目本身的示例已经说明了解决步骤。想想怎么细化实现：
```text
用一个哈希表记录A、B里所有字母对应的最小等价字母，如字母a=b，那么map里有 a:a, b:a两个键值对，
如果又有b=x, map里增加键值对 x:a，注意等价字符的传递性
从前向后遍历A、B，将A[i]和B[i]作为键插入哈希表，其值应该是表里已经有的等价字母中最小的，这里需要循环查找
最后遍历S，根据哈希表，将S里的字母一一替换成哈希表里最小的等价字母
```
```go
func smallestEquivalentString(A string, B string, S string) string {
	if len(A) == 0 || len(A) != len(B) {
		return ""
	}
	m := make(map[byte]byte, 26) // m记录A、B里所有字母对应的最小等价字母
	for i := range A {
		add(m, A[i], B[i])
	}
	r := []byte(S)
	for i, v := range r {
		r[i] = find(m, v)
	}
	return string(r)
}

func add(m map[byte]byte, k1, k2 byte) {
	c1 := find(m, k1)
	c2 := find(m, k2)
	c := min(c1, c2)
	m[c1] = c
	m[c2] = c
	m[k1] = c
	m[k2] = c
}

func find(m map[byte]byte, k byte) byte {
	for {
		c, ok := m[k]
		if !ok || c == k {
			break
		}
		k = c
	}
	return k
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}
```
蓦然发现，这就是并查集的思想啊，还可以进一步优化:
```text
add函数相当于union; 而且注意到对c1，c2的处理，根据题意，需要较小的字母做父节点，这里不应该按秩合并
find缺一点路径压缩的处理，不难实现；
另外因为都是小写字母，最多26个字母，其实可以用切片替换哈希表
```
优化后发现leetcode上的用例耗时从4ms降低到了0ms，<br>
主要是map改成切片得到的优化，路径压缩在这些用例里对性能的提升不明显<br>
最终实现：
```go
func smallestEquivalentString(A string, B string, S string) string {
	if len(A) == 0 || len(A) != len(B) {
		return ""
	}
	const maxLetters = 26
	const firstLetter = 'a'
	uf := MakeUnionFind(maxLetters)
	for i := range A {
		uf.Union(int(A[i]-firstLetter), int(B[i]-firstLetter))
	}
	r := []byte(S)
	for i, v := range r {
		r[i] = byte(uf.Find(int(v-firstLetter))) + firstLetter
	}
	return string(r)
}

type UnionFind []int

func MakeUnionFind(n int) UnionFind {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}

func (uf UnionFind) Union(v1, v2 int) {
	c1 := uf.Find(v1)
	c2 := uf.Find(v2)
	c := min(c1, c2)
	uf[c1], uf[c2] = c, c // 不要按秩合并，让较小的节点做父节点
	uf[v1], uf[v2] = c, c
}

func (uf UnionFind) Find(v int) int {
	for v != uf[v] {
		v, uf[v] = uf[v], uf[uf[v]] // 路径压缩
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
[参考实现](d.go)