package system

import (
	"container/heap"
)

/*
642. 设计搜索自动补全系统 https://leetcode-cn.com/problems/design-search-autocomplete-system
为搜索引擎设计一个搜索自动补全系统。用户会输入一条语句（最少包含一个字母，以特殊字符 '#' 结尾）。
除 '#' 以外用户输入的每个字符，返回历史中热度前三并以当前输入部分为前缀的句子。下面是详细规则：

一条句子的热度定义为历史上用户输入这个句子的总次数。
返回前三的句子需要按照热度从高到低排序（第一个是最热门的）。如果有多条热度相同的句子，请按照 ASCII 码的顺序输出（ASCII 码越小排名越前）。
如果满足条件的句子个数少于 3，将它们全部输出。
如果输入了特殊字符，意味着句子结束了，请返回一个空集合。
你的工作是实现以下功能：

构造函数：

AutocompleteSystem(String[] sentences, int[] times): 这是构造函数，输入的是历史数据。
Sentences 是之前输入过的所有句子，Times 是每条句子输入的次数，你的系统需要记录这些历史信息。

现在，用户输入一条新的句子，下面的函数会提供用户输入的下一个字符：

List<String> input(char c): 其中 c 是用户输入的下一个字符。
字符只会是小写英文字母（'a' 到 'z' ），空格（' '）和特殊字符（'#'）。输出历史热度前三的具有相同前缀的句子。



样例 ：
操作 ： AutocompleteSystem(["i love you", "island","ironman", "i love leetcode"], [5,3,2,2])
系统记录下所有的句子和出现的次数：
"i love you" : 5 次
"island" : 3 次
"ironman" : 2 次
"i love leetcode" : 2 次
现在，用户开始新的键入：


输入 ： input('i')
输出 ： ["i love you", "island","i love leetcode"]
解释 ：
有四个句子含有前缀 "i"。其中 "ironman" 和 "i love leetcode" 有相同的热度，
由于 ' ' 的 ASCII 码是 32 而 'r' 的 ASCII 码是 114，所以 "i love leetcode" 在 "ironman" 前面。
同时我们只输出前三的句子，所以 "ironman" 被舍弃。

输入 ： input(' ')
输出 ： ["i love you","i love leetcode"]
解释:
只有两个句子含有前缀 "i "。

输入 ： input('a')
输出 ： []
解释 ：
没有句子有前缀 "i a"。

输入 ： input('#')
输出 ： []
解释 ：

用户输入结束，"i a" 被存到系统中，后面的输入被认为是下一次搜索。



注释 ：

输入的句子以字母开头，以 '#' 结尾，两个字母之间最多只会出现一个空格。
即将搜索的句子总数不会超过 100。每条句子的长度（包括已经搜索的和即将搜索的）也不会超过 100。
即使只有一个字母，输出的时候请使用双引号而不是单引号。
请记住清零 AutocompleteSystem 类中的变量，因为静态变量、类变量会在多组测试数据中保存之前结果。详情请看这里。
*/
/*
前缀树的一个应用
*/
type Trie struct {
	links map[byte]*Trie
	times int
}

func (t *Trie) Insert(s string, times int) {
	n := t
	for i := range s {
		if n.links[s[i]] == nil {
			n.links[s[i]] = &Trie{links: map[byte]*Trie{}}
		}
		n = n.links[s[i]]
	}
	n.times += times
}

type AutocompleteSystem struct {
	trie     *Trie
	searched *Trie
	buf      []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	trie := &Trie{links: map[byte]*Trie{}}
	for i, v := range sentences {
		trie.Insert(v, times[i])
	}
	return AutocompleteSystem{trie: trie, searched: trie}
}

func (s *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		s.trie.Insert(string(s.buf), 1)
		s.buf = s.buf[:0]
		s.searched = s.trie
		return nil
	}
	s.buf = append(s.buf, c)
	if s.searched == nil || s.searched.links[c] == nil {
		s.searched = nil
		return nil
	}
	s.searched = s.searched.links[c]
	result := help(s.searched, c, string(s.buf[:len(s.buf)-1]))
	return result
}

func help(trie *Trie, curr byte, prefix string) []string {
	h := &Heap{} // 依据热度排列的小顶堆，热度相同则字典序大的在堆顶，最多维持3个元素
	var path []byte
	var dfs func(t *Trie, c byte)
	dfs = func(t *Trie, c byte) {
		path = append(path, c)
		if len(t.links) == 0 || t.times > 0 {
			item := HeapItem{s: prefix + string(path), t: t.times}
			heap.Push(h, item)
			if h.Len() > 3 {
				_ = heap.Pop(h)
			}
		}
		for k, v := range t.links {
			dfs(v, k)
		}
		path = path[:len(path)-1]
	}
	dfs(trie, curr)

	result := make([]string, h.Len())
	i := h.Len() - 1
	for h.Len() > 0 {
		result[i] = heap.Pop(h).(HeapItem).s
		i--
	}
	return result
}

type HeapItem struct {
	s string
	t int
}

type Heap struct {
	s []HeapItem
}

func (h *Heap) Len() int      { return len(h.s) }
func (h *Heap) Swap(i, j int) { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *Heap) Less(i, j int) bool {
	if h.s[i].t == h.s[j].t {
		return h.s[i].s > h.s[j].s
	}
	return h.s[i].t < h.s[j].t
}
func (h *Heap) Push(x interface{}) { h.s = append(h.s, x.(HeapItem)) }
func (h *Heap) Pop() interface{} {
	r := h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]
	return r
}
