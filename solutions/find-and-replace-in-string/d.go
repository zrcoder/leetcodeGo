package find_and_replace_in_string

import (
	"sort"
	"strings"
)

/*
833. 字符串中的查找与替换
https://leetcode-cn.com/problems/find-and-replace-in-string/

对于某些字符串 S，我们将执行一些替换操作，用新的字母组替换原有的字母组（不一定大小相同）。

每个替换操作具有 3 个参数：起始索引 i，源字 x 和目标字 y。
规则是如果 x 从原始字符串 S 中的位置 i 开始，那么我们将用 y 替换出现的 x。如果没有，我们什么都不做。

举个例子，如果我们有 S = “abcd” 并且我们有一些替换操作 i = 2，x = “cd”，y = “ffff”，那么因为 “cd” 从原始字符串 S 中的位置 2 开始，我们将用 “ffff” 替换它。

再来看 S = “abcd” 上的另一个例子，如果我们有替换操作 i = 0，x = “ab”，y = “eee”，
以及另一个替换操作 i = 2，x = “ec”，y = “ffff”，那么第二个操作将不执行任何操作，因为原始字符串中 S[2] = 'c'，与 x[0] = 'e' 不匹配。

所有这些操作同时发生。保证在替换时不会有任何重叠： S = "abc", indexes = [0, 1], sources = ["ab","bc"] 不是有效的测试用例。


示例 1：
输入：S = "abcd", indexes = [0,2], sources = ["a","cd"], targets = ["eee","ffff"]
输出："eeebffff"
解释：
"a" 从 S 中的索引 0 开始，所以它被替换为 "eee"。
"cd" 从 S 中的索引 2 开始，所以它被替换为 "ffff"。

示例 2：
输入：S = "abcd", indexes = [0,2], sources = ["ab","ec"], targets = ["eee","ffff"]
输出："eeecd"
解释：
"ab" 从 S 中的索引 0 开始，所以它被替换为 "eee"。
"ec" 没有从原始的 S 中的索引 2 开始，所以它没有被替换。


提示：
0 <= indexes.length = sources.length = targets.length <= 100
0 < indexes[i] < S.length <= 1000
给定输入中的所有字符都是小写字母。
*/
/*
借助strings.Builder或bytes.buffer
需要先把替换数据按照索引排序
*/
type ReplaceItem struct {
	index  int
	source string
	target string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	replaceInfo := make([]ReplaceItem, len(indexes))
	for i := 0; i < len(indexes); i++ {
		replaceInfo[i] = ReplaceItem{indexes[i], sources[i], targets[i]}
	}
	sort.Slice(replaceInfo, func(i, j int) bool {
		return replaceInfo[i].index < replaceInfo[j].index
	})
	index := 0
	buf := strings.Builder{}
	buf.Grow(100000)
	for _, info := range replaceInfo {
		if index < info.index {
			buf.WriteString(S[index:info.index])
			index = info.index
		}
		end := info.index + len(info.source)
		if end <= len(S) && S[info.index:end] == info.source {
			buf.WriteString(info.target)
			index = end
		}
	}
	return buf.String() + S[index:]
}

// 可先统计哪些位置可替换
func findReplaceString1(S string, indexes []int, sources []string, targets []string) string {
	matched := make([]int, len(S))
	for i := range matched {
		matched[i] = -1
	}
	for i, index := range indexes {
		if S[index:index+len(sources[i])] == sources[i] {
			matched[index] = i
		}
	}
	buf := strings.Builder{}
	buf.Grow(100000)
	for i := 0; i < len(S); {
		if matched[i] == -1 {
			buf.WriteByte(S[i])
			i++
			continue
		}
		j := matched[i]
		buf.WriteString(targets[j])
		i += len(sources[j])
	}
	return buf.String()
}
