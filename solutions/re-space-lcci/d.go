package lcci

/*
面试题 17.13. 恢复空格 https://leetcode-cn.com/problems/re-space-lcci

哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。
像句子"I reset the computer. It still didn’t boot!"
已经变成了"iresetthecomputeritstilldidntboot"。
在处理标点符号和大小写之前，你得先把它断成词语。
当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。
假设文章用sentence表示，设计一个算法，把文章断开，
要求未识别的字符最少，返回未识别的字符数。

注意：本题相对原题稍作改动，只需返回未识别的字符数

示例：
输入：
dictionary = ["looked","just","like","her","brother"]
sentence = "jesslookedjustliketimherbrother"
输出： 7
解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。

提示：
0 <= len(sentence) <= 1000
dictionary中总字符数不超过 150000。
*/

/*
常规dp
定义长度为sentence长度+1的dp数组，dp[i]表示sentence[:i]在字典中未识别的字符数
判断sentence[j:i]是否在词典中，其中0<= j < i。为方便迅速查找，事先将dictionary数组转成map
如果sentence[j:i]在字典中，则dp[i] = min(dp[i], dp[j])，否则dp[i] = dp[i-1]+1

时间复杂度 O(n^2), 空间复杂度O(max(n, m)), 其中n、m分别为sentence长度和字典大小
*/
func respace(dictionary []string, sentence string) int {
	dic := make(map[string]bool, len(dictionary))
	for _, s := range dictionary {
		dic[s] = true
	}
	// dp[i]代表sentence[:i]中未识别的字符数
	dp := make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			word := sentence[j:i]
			if dic[word] && dp[j] < dp[i] {
				dp[i] = dp[j]
			}
		}
	}
	return dp[len(sentence)]
}
