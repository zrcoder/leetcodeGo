package stickers_to_spell_word

import (
	"math"
)

/*
691. 贴纸拼词
https://leetcode-cn.com/problems/stickers-to-spell-word

我们给出了 N 种不同类型的贴纸。每个贴纸上都有一个小写的英文单词。
你希望从自己的贴纸集合中裁剪单个字母并重新排列它们，从而拼写出给定的目标字符串 target。
如果你愿意的话，你可以不止一次地使用每一张贴纸，而且每一张贴纸的数量都是无限的。
拼出目标 target 所需的最小贴纸数量是多少？如果任务不可能，则返回 -1。

示例 1：
输入：
["with", "example", "science"], "thehat"
输出：
3
解释：
我们可以使用 2 个 "with" 贴纸，和 1 个 "example" 贴纸。
把贴纸上的字母剪下来并重新排列后，就可以形成目标 “thehat“ 了。
此外，这是形成目标字符串所需的最小贴纸数量。

示例 2：
输入：
["notice", "possible"], "basicbasic"
输出：
-1
解释：
我们不能通过剪切给定贴纸的字母来形成目标“basicbasic”。

提示：
stickers 长度范围是 [1, 50]。
stickers 由小写英文单词组成（不带撇号）。
target 的长度在 [1, 15] 范围内，由小写字母组成。
在所有的测试案例中，所有的单词都是从 1000 个最常见的美国英语单词中随机选取的，目标是两个随机单词的串联。
时间限制可能比平时更具挑战性。预计 50 个贴纸的测试案例平均可在35ms内解决。
*/

/* 方法一 优化穷举搜索

时间复杂度：N 作为贴纸的数目，T 作为目标单词的字母数目。时间复杂度的界限是 O(N^(T+1)*T^2):
对于每个贴纸，须尝试使用最多T+1次，并更新目标计数成本 O(T)，最多做 T 次。
空间复杂度：O(N+T)

值得一提的是，infos可以用链表，每次实际去删除无用贴纸，实际测试比切片标记版本慢10倍左右
*/
var result int

type Sticker struct {
	Letters  map[rune]int
	IsUseful bool
}

func minStickers1(stickers []string, target string) int {
	// 记录target里有哪些字母，各有几个
	needed := make(map[rune]int, len(target))
	// 仅记录target里有哪些字母，便于后续判断stickers里是否不含target里的字母
	set := make(map[rune]bool, len(target))
	for _, v := range target {
		needed[v]++
		set[v] = true
	}
	// 只保留每张贴纸里target中包含的字母，并记录对应字母出现的次数
	infos := make([]Sticker, len(stickers))
	preprocess(infos, stickers, needed, set)
	if len(set) > 0 { // 目标存在贴纸上没有的字母
		return -1
	}

	result = -1 // 设置为-1，以区分是否从来没设置过
	backTrack(infos, needed, 0, 0)
	return result
}

func preprocess(infos []Sticker, stickers []string, needed map[rune]int, set map[rune]bool) {
	for i, s := range stickers {
		infos[i].Letters = make(map[rune]int, len(s))
		for _, char := range s {
			if needed[char] == 0 {
				continue
			}
			infos[i].Letters[char]++
			delete(set, char)
		}
		infos[i].IsUseful = len(infos[i].Letters) > 0
	}
}
func backTrack(stickers []Sticker, needed map[rune]int, used int, start int) {
	if len(needed) == 0 {
		if result == -1 || used < result {
			result = used
		}
		return
	}

	currStickers := filterStickers(stickers, needed)
	for i := start; i < len(currStickers); i++ {
		if !currStickers[i].IsUseful {
			continue
		}
		currNeeded := filterNeeded(stickers[i].Letters, needed)
		backTrack(currStickers, currNeeded, used+1, i)
	}
}

func filterStickers(stickers []Sticker, needed map[rune]int) []Sticker {
	currStickers := make([]Sticker, len(stickers))
	_ = copy(currStickers, stickers)
	markUseful(currStickers, needed)
	return currStickers
}

func markUseful(stickers []Sticker, needed map[rune]int) {
	markUsefulStickers(stickers, needed)
	removeUselessSubStickers(stickers, needed)
}

func markUsefulStickers(stickers []Sticker, needed map[rune]int) {
	for i := range stickers {
		if !stickers[i].IsUseful {
			continue
		}
		stickers[i].IsUseful = isNeeded(stickers[i].Letters, needed)
	}
}

func isNeeded(letters map[rune]int, needed map[rune]int) bool {
	for k := range needed {
		if letters[k] > 0 {
			return true
		}
	}
	return false
}

// 去除可以被其他贴纸完全覆盖的贴纸
func removeUselessSubStickers(stickers []Sticker, needed map[rune]int) {
	n := len(stickers)
	for i := 0; i < n-1; i++ {
		vi := stickers[i]
		if !vi.IsUseful {
			continue
		}
		for j := i + 1; j < n; j++ {
			vj := stickers[j]
			if !vj.IsUseful {
				continue
			}
			isISub, isJSub := isSubSticker(needed, vi.Letters, vj.Letters)
			if isISub {
				stickers[i].IsUseful = false
			} else if isJSub {
				stickers[j].IsUseful = false
			}
		}
	}
}

func isSubSticker(needed map[rune]int, leftSticker, rightSticker map[rune]int) (bool, bool) {
	isLeftSub, isRightSub := true, true
	// 判断左边或者右边的贴纸是否是子贴纸
	for k := range needed {
		if !isLeftSub && !isRightSub {
			break
		}
		if leftSticker[k] == 0 && rightSticker[k] == 0 {
			continue
		}
		if leftSticker[k] > rightSticker[k] {
			isLeftSub = false
		} else if leftSticker[k] < rightSticker[k] {
			isRightSub = false
		}
	}
	return isLeftSub, isRightSub
}

// 使用当前贴纸后,还需要多少字母
func filterNeeded(supplied, needed map[rune]int) map[rune]int {
	result := make(map[rune]int, len(needed))
	for k, v := range needed {
		if v-supplied[k] > 0 {
			result[k] = v - supplied[k]
		}
	}
	return result
}

/*
方法二：动态规划

假设 target 长度为 n，已经限定 n 不会超过 15，
可以用一个 int 来表示 target 被完成的状态，int 变量的二进制表示中，1表示对应位置的字母已经完成，0表示未完成
那么所有的状态将是

000...000
000...001
...
111......

每个状态的长度是 15，只用一个 int 的后 15 位，前边的 0 不予关注
这样的状态共有 2^n 种，即 1<<n 种

对于每一个状态 state ，在应用一个贴纸之后变成另一个状态 newState。
则 dp(newState) = min(dp(newSatate), dp(state)+1)

时间复杂度：O(2^T * S * T)其中 S 是所有贴纸中的字母总数，T 是目标单词中的字母数。可以仔细检查每个循环，得出这个结论。
空间复杂度：O(2^T)，dp 使用的空间。
*/

func minStickers(stickers []string, target string) int {
	total := 1 << len(target)
	dp := make([]int, total)
	for i := 1; i < total; i++ {
		dp[i] = math.MaxInt32 // dp[0] == 0, 先标记所有其他状态未尝试，同时为了方便后边的状态转移比较，标记为max
	}
	indices := calIndices(target)
	for _, sticker := range stickers {
		for state := 0; state < total; state++ {
			if dp[state] == math.MaxInt32 {
				continue
			}
			newState := getStateAfterUse(sticker, state, indices)
			if dp[newState] > dp[state]+1 {
				dp[newState] = dp[state] + 1 // 从state到newState，使用了一张卡片（就是sticker）
			}
		}
	}
	if dp[total-1] == math.MaxInt32 {
		return -1
	}
	return dp[total-1]
}

func calIndices(target string) map[rune][]int {
	indices := make(map[rune][]int, 0)
	for i, v := range target {
		indices[v] = append(indices[v], i)
	}
	return indices
}

// 返回在state状态下，使用贴纸sticker后得到的新状态
func getStateAfterUse(sticker string, state int, indexes map[rune][]int) int {
	for _, char := range sticker {
		for _, index := range indexes[char] {
			if state&(1<<index) == 0 {
				state |= 1 << index
				break
			}
		}
	}
	return state
}
