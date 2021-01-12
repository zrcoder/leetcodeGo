/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package next_closest_time

import "strconv"

/*
681. 最近时刻 https://leetcode-cn.com/problems/next-closest-time
给定一个形如 “HH:MM” 表示的时刻，利用当前出现过的数字构造下一个距离当前时间最近的时刻。每个出现数字都可以被无限次使用。
你可以认为给定的字符串一定是合法的。例如，“01:34” 和 “12:09” 是合法的，“1:34” 和 “12:9” 是不合法的。

样例 1:
输入: "19:34"
输出: "19:39"
解释: 利用数字 1, 9, 3, 4 构造出来的最近时刻是 19:39，是 5 分钟之后。结果不是 19:33 因为这个时刻是 23 小时 59 分钟之后。

样例 2:
输入: "23:59"
输出: "22:22"
解释: 利用数字 2, 3, 5, 9 构造出来的最近时刻是 22:22。 答案一定是第二天的某一时刻，所以选择可构造的最小时刻。
*/

/*
模拟时钟前进一分钟。每次向前移动时，如果当前时间能够被构造，则返回当前时间
时间复杂度：O(1)——尝试最多24∗60的可能时间
空间复杂度：O(1)
*/
func nextClosestTime(time string) string {
	a, b, c, d := int(time[0]-'0'), int(time[1]-'0'), int(time[3]-'0'), int(time[4]-'0')
	set := GenSet(a, b, c, d)
	hour := a*10 + b
	minute := c*10 + d
	for i := 0; i < 24*60; i++ {
		hour, minute = add1minute(hour, minute)
		a, b, c, d = hour/10, hour%10, minute/10, minute%10
		if set.Has(a) && set.Has(b) && set.Has(c) && set.Has(d) {
			return genTimeStr(a, b, c, d)
		}
	}
	return ""
}

type Set map[int]struct{}

func GenSet(elements ...int) Set {
	s := make(map[int]struct{}, len(elements))
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

func (s Set) Has(i int) bool {
	_, ok := s[i]
	return ok
}

func add1minute(hour, minute int) (int, int) {
	if minute == 59 {
		return (hour + 1) % 24, 0
	}
	return hour, minute + 1
}

func genTimeStr(a, b, c, d int) string {
	h := a*10 + b
	hStr := strconv.Itoa(h)
	if h < 10 {
		hStr = "0" + hStr
	}
	m := c*10 + d
	mStr := strconv.Itoa(m)
	if m < 10 {
		mStr = "0" + mStr
	}
	return hStr + ":" + mStr
}
