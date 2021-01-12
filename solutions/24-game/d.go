/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package pointgame

import (
	"math"
)

/*
679. 24 点游戏
https://leetcode-cn.com/problems/24-game

你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。

示例 1:
输入: [4, 1, 8, 7]
输出: True
解释: (8-4) * (7-1) = 24

示例 2:
输入: [1, 2, 1, 2]
输出: False

注意:
除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
*/

/*
递归回溯
只有四张牌，四种运算， 总共的可能组合为：
先选两张牌，4*3，四种运算，可能有4*3*4种
已选2张牌运算后的结果与剩下两张牌组合，现在相当于有3张牌；选出其中两张运算的情况有3*2*4种
现在相当于有2张牌，有2*4种运算情况
综上，总共可能有(4*3*4) * (3*2*4) * (2*4) =9216种组合

时空复杂度都是O(1)

*/

func judgePoint24(nums []int) bool {
	return judge(parseFloats(nums))
}

func parseFloats(nums []int) []float64 {
	floats := make([]float64, len(nums))
	for i, v := range nums {
		floats[i] = float64(v)
	}
	return floats
}

func judge(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return math.Abs(nums[0]-24.0) < 1e-6
	}
	for i := range nums {
		for j := range nums {
			if i != j && try(i, j, nums) {
				return true
			}
		}
	}
	return false
}

type Operator func(i, j float64) float64

var (
	add   = func(i, j float64) float64 { return i + j }
	time  = func(i, j float64) float64 { return i * j }
	minus = func(i, j float64) float64 { return i - j }
	div   = func(i, j float64) float64 { return i / j }

	operators = [4]Operator{add, time, minus, div}
)

func try(i, j int, nums []float64) bool {
	// nums[i]和nums[j] 参与运算
	tmp := remove(nums, i, j)
	originSize := len(tmp)
	for k := 0; k < 4; k++ { // 尝试 "+*-/" 四种操作
		if k < 2 && j > i || k == 3 && nums[j] == 0 {
			continue
		}
		r := operators[k](nums[i], nums[j])
		if len(tmp) > originSize {
			tmp[len(tmp)-1] = r
		} else {
			tmp = append(tmp, r)
		}
		if judge(tmp) {
			return true
		}
	}
	return false
}

func remove(nums []float64, i, j int) []float64 {
	r := make([]float64, 0)
	for k, v := range nums {
		if k != i && k != j {
			r = append(r, v)
		}
	}
	return r
}
