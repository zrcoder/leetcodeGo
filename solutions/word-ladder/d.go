package ladder

import "container/list"

/*
126. 单词接龙 II https://leetcode-cn.com/problems/word-ladder-ii/
给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。

说明:
如果不存在这样的转换序列，返回一个空列表。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]

示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: []

解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	hasBeginWord, hasEndWord := false, false
	for _, v := range wordList {
		if v == beginWord {
			hasBeginWord = true
		} else if v == endWord {
			hasEndWord = true
		}
	}
	var result [][]string
	if !hasEndWord {
		return result
	}
	if !hasBeginWord {
		wordList = append(wordList, beginWord)
	}

	neighbors := initNeighbors(wordList) // 记录每个单词改变一个字母能得到且存在与wordList的单词列表
	steps := initSteps(wordList)         // 记录从beginWord变换到当前单词需要到步数
	// bfs
	queue := list.New()
	queue.PushBack([]string{beginWord})
	steps[beginWord] = 0
	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]string)
		word := path[len(path)-1]
		if word == endWord {
			result = append(result, path)
			continue
		}
		for _, next := range neighbors[word] {
			if steps[word]+1 > steps[next] {
				continue
			}
			steps[next] = steps[word] + 1
			newPath := make([]string, len(path)+1)
			_ = copy(newPath, path)
			newPath[len(newPath)-1] = next
			queue.PushBack(newPath)
		}
	}
	return result
}

func initNeighbors(wordList []string) map[string][]string {
	neighbors := map[string][]string{}
	for i := range wordList {
		for j := i + 1; j < len(wordList); j++ {
			if canTransformByOneChar(wordList[i], wordList[j]) {
				neighbors[wordList[i]] = append(neighbors[wordList[i]], wordList[j])
				neighbors[wordList[j]] = append(neighbors[wordList[j]], wordList[i])
			}
		}
	}
	return neighbors
}

func canTransformByOneChar(s, t string) bool {
	for i := range s {
		if s[i] != t[i] {
			return s[i+1:] == t[i+1:]
		}
	}
	return false
}

func initSteps(wordList []string) map[string]int {
	steps := make(map[string]int, len(wordList))
	for _, word := range wordList {
		steps[word] = len(wordList)
	}
	return steps
}

/*
127. 单词接龙 https://leetcode-cn.com/problems/word-ladder
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:
如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]
输出: 5
解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
	 返回它的长度 5。

示例 2:
输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]
输出: 0
解释: endWord "cog" 不在字典中，所以无法进行转换。
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	hasBeginWord, hasEndWord := false, false
	for _, v := range wordList {
		if v == beginWord {
			hasBeginWord = true
		} else if v == endWord {
			hasEndWord = true
		}
	}
	if !hasEndWord {
		return 0
	}
	if !hasBeginWord {
		wordList = append(wordList, beginWord)
	}

	neighbors := initNeighbors(wordList) // 记录每个单词改变一个字母能得到且存在与wordList的单词列表
	steps := initSteps(wordList)         // 记录从beginWord变换到当前单词需要到步数
	// bfs
	queue := list.New()
	queue.PushBack(beginWord)
	steps[beginWord] = 0
	for queue.Len() > 0 {
		word := queue.Remove(queue.Front()).(string)
		if word == endWord {
			return steps[endWord] + 1
		}
		for _, next := range neighbors[word] {
			if steps[word]+1 > steps[next] {
				continue
			}
			steps[next] = steps[word] + 1
			queue.PushBack(next)
		}
	}
	return 0
}
