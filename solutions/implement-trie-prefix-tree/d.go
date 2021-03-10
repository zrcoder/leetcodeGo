/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package implement_trie_prefix_tree

import "strings"

/*
208. 实现 Trie (前缀树) https://leetcode-cn.com/problems/implement-trie-prefix-tree
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。
*/
type Trie struct {
	links map[byte]*Trie // 这道题保证了输入的都是小写字母，links也可用大小为26的数组实现
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		links: map[byte]*Trie{},
	}
}

/** Inserts a word into the trie.

Time complexity : O(m), where m is the key length.

Space complexity : O(m)
In the worst case newly inserted key doesn't share a prefix with the the keys already inserted in the trie.
We have to add m new nodes, which takes us O(m) space.
*/
func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		if node.links[word[i]] == nil {
			node.links[word[i]] = &Trie{links: map[byte]*Trie{}}
		}
		node = node.links[word[i]]
	}
	node.IsEnd = true
}

/** Returns if the word is in the trie.

Time complexity : O(m) In each step of the algorithm we search for the next key character.
In the worst case the algorithm performs mm operations.

Space complexity : O(1)
*/
func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix.
Time complexity : O(m)

Space complexity : O(1)
*/
func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}

func (t *Trie) search(s string) *Trie {
	node := t
	for i := 0; i < len(s); i++ {
		if node.links[s[i]] == nil {
			return nil
		}
		node = node.links[s[i]]
	}
	return node
}

/*
Time complexity : O(m) In each step of the algorithm we search for the next key character.
In the worst case the algorithm performs m operations.

Space complexity : O(1)
*/
func (t *Trie) SearchLongestPrefixOf(word string) string {
	node := t
	i := 0
	for i < len(word) {
		ch := word[i]
		if node.links[ch] == nil || len(node.links) > 1 || node.IsEnd {
			return word[:i]
		}
		node = node.links[ch]
		i++
	}
	return word[:i]
}

/* 变体
677. 键值映射 https://leetcode-cn.com/problems/map-sum-pairs/

实现一个 MapSum 类里的两个方法，insert 和 sum。

对于方法 insert，你将得到一对（字符串，整数）的键值对。字符串表示键，整数表示值。
如果键已经存在，那么原来的键值对将被替代成新的键值对。

对于方法 sum，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。

示例 1:

输入: insert("apple", 3), 输出: Null
输入: sum("ap"), 输出: 3
输入: insert("app", 2), 输出: Null
输入: sum("ap"), 输出: 5
*/
// 上面的前缀树节点增加 val 属性
type MapSum struct {
	links map[byte]*MapSum
	val   int
	isEnd bool
}

func Constructor1() MapSum {
	return MapSum{links: map[byte]*MapSum{}}
}

func (s *MapSum) Insert(key string, val int) {
	node := s
	for i := range key {
		if node.links[key[i]] == nil {
			node.links[key[i]] = &MapSum{links: map[byte]*MapSum{}}
		}
		node = node.links[key[i]]
	}
	node.val = val
	node.isEnd = true
}

func (s *MapSum) Sum(prefix string) int {
	node := s
	for i := range prefix {
		if node.links[prefix[i]] == nil {
			return 0
		}
		node = node.links[prefix[i]]
	}
	result := 0
	var dfs func(t *MapSum)
	dfs = func(t *MapSum) {
		if t.isEnd {
			result += t.val
		}
		for _, v := range t.links {
			dfs(v)
		}
	}
	dfs(node)
	return result
	// 也可用bfs，如：
	//queue, total := []*MapSum{node}, 0
	//if node.isEnd {
	//	total += node.val
	//}
	//for len(queue) > 0 {
	//	node := queue[0]
	//	queue = queue[1:]
	//	for _, next := range node.links {
	//		if next.isEnd {
	//			total += next.val
	//		}
	//		queue = append(queue, next)
	//	}
	//}
	//return total
}

/* 应用
648. 单词替换 https://leetcode-cn.com/problems/replace-words

在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。
例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。
如果继承词有许多可以形成它的词根，则用最短的词根替换它。

你需要输出替换之后的句子。

示例：

输入：dict(词典) = ["cat", "bat", "rat"] sentence(句子) = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

提示：
输入只包含小写字母。
1 <= dict.length <= 1000
1 <= dict[i].length <= 100
1 <= 句中词语数 <= 1000
1 <= 句中词语长度 <= 1000
*/
func replaceWords(dict []string, sentence string) string {
	trie := &Trie{links: map[byte]*Trie{}}
	for _, v := range dict {
		trie.Insert(v)
	}
	res := strings.Split(sentence, " ")
	for i, v := range res {
		p := trie.SearchBreifPrifix(v)
		if p != "" {
			res[i] = p
		}
	}
	return strings.Join(res, " ")
}

func (t *Trie) SearchBreifPrifix(word string) string {
	i := 0
	node := t
	for i < len(word) {
		if node.IsEnd || node.links[word[i]] == nil {
			break
		}
		node = node.links[word[i]]
		i++
	}
	if node.IsEnd {
		return word[:i]
	}
	return ""
}

/*
211. 添加与搜索单词 - 数据结构设计 https://leetcode-cn.com/problems/add-and-search-word-data-structure-design
设计一个支持以下两种操作的数据结构：

void addWord(word)
bool search(word)
search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。

示例:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
说明:

你可以假设所有单词都是由小写字母 a-z 组成的。
*/
type WordDictionary struct {
	links map[byte]*WordDictionary
	isEnd bool
}

func Constructor2() WordDictionary {
	return WordDictionary{
		links: map[byte]*WordDictionary{},
	}
}

func (d *WordDictionary) AddWord(word string) {
	node := d
	for i := range word {
		if node.links[word[i]] == nil {
			node.links[word[i]] = &WordDictionary{links: map[byte]*WordDictionary{}}
		}
		node = node.links[word[i]]
	}
	node.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	node := d
	if len(word) == 0 {
		return node.isEnd
	}
	if len(node.links) == 0 || word[0] != '.' && node.links[word[0]] == nil {
		return false
	}
	if word[0] == '.' {
		for _, v := range node.links {
			if v.Search(word[1:]) {
				return true
			}
		}
		return false
	}
	node = node.links[word[0]]
	return node.Search(word[1:])
}

/*
212. 单词搜索 II https://leetcode-cn.com/problems/word-search-ii

给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

输出: ["eat","oath"]
说明:
你可以假设所有输入都由小写字母 a-z 组成。

提示:

你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
*/
func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}
	m, n := len(board), len(board[0])
	trie := &Trie{links: make(map[byte]*Trie, 0)}
	for _, v := range words {
		trie.Insert(v)
	}
	set := make(map[string]struct{}, 0)
	var path []byte
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(t *Trie, r, c int, seen [][]bool)
	dfs = func(t *Trie, r, c int, seen [][]bool) {
		if r < 0 || r >= m || c < 0 || c >= n || seen[r][c] {
			return
		}
		ch := board[r][c]
		if t.links[ch] == nil {
			return
		}
		path = append(path, ch)
		seen[r][c] = true
		if t.links[ch].IsEnd {
			set[string(path)] = struct{}{}
		}
		for _, d := range dirs {
			dfs(t.links[ch], r+d[0], c+d[1], seen)
		}
		// 回溯
		path = path[:len(path)-1]
		seen[r][c] = false
	}
	for r := range board {
		for c := range board[r] {
			dfs(trie, r, c, genSeen(m, n))
		}
	}
	result := make([]string, 0, len(set))
	for k := range set {
		result = append(result, k)
	}
	return result
}

func genSeen(m, n int) [][]bool {
	r := make([][]bool, m)
	for i := range r {
		r[i] = make([]bool, n)
	}
	return r
}
