package restore_ip_addresses

import (
	"strconv"
	"strings"
)

/*
93. 复原IP地址
https://leetcode-cn.com/problems/restore-ip-addresses

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。

示例:
输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/

const (
	maxIp    = 255
	count    = 4
	minIpLen = 4
	maxLen   = 12
)

/*
常规dfs
*/
func restoreIpAddresses(s string) []string {
	if len(s) < minIpLen || len(s) > maxLen {
		return nil
	}
	var result, path []string
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) && len(path) == count {
			result = append(result, strings.Join(path, "."))
			return
		}
		end := min(len(s), start+minIpLen)
		for i := start + 1; i <= end; i++ {
			seg := s[start:i]
			if !isValid(seg) {
				continue
			}
			path = append(path, seg)
			dfs(i)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	return num <= maxIp
}
