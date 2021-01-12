/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package number_of_burgers_with_no_waste_of_ingredients

import "sort"

/*
1276. 不浪费原料的汉堡制作方案 https://leetcode-cn.com/problems/number-of-burgers-with-no-waste-of-ingredients/

圣诞活动预热开始啦，汉堡店推出了全新的汉堡套餐。为了避免浪费原料，请你帮他们制定合适的制作计划。
给你两个整数 tomatoSlices 和 cheeseSlices，分别表示番茄片和奶酪片的数目。不同汉堡的原料搭配如下：

巨无霸汉堡：4 片番茄和 1 片奶酪
小皇堡：2 片番茄和 1 片奶酪
请你以 [total_jumbo, total_small]（[巨无霸汉堡总数，小皇堡总数]）的格式返回恰当的制作方案，
使得剩下的番茄片 tomatoSlices 和奶酪片 cheeseSlices 的数量都是 0。

如果无法使剩下的番茄片 tomatoSlices 和奶酪片 cheeseSlices 的数量为 0，就请返回 []。

示例 1：
输入：tomatoSlices = 16, cheeseSlices = 7
输出：[1,6]
解释：制作 1 个巨无霸汉堡和 6 个小皇堡需要 4*1 + 2*6 = 16 片番茄和 1 + 6 = 7 片奶酪。不会剩下原料。

示例 2：
输入：tomatoSlices = 17, cheeseSlices = 4
输出：[]
解释：只制作小皇堡和巨无霸汉堡无法用光全部原料。

示例 3：
输入：tomatoSlices = 4, cheeseSlices = 17
输出：[]
解释：制作 1 个巨无霸汉堡会剩下 16 片奶酪，制作 2 个小皇堡会剩下 15 片奶酪。

示例 4：
输入：tomatoSlices = 0, cheeseSlices = 0
输出：[0,0]

示例 5：
输入：tomatoSlices = 2, cheeseSlices = 1
输出：[0,1]

提示：
0 <= tomatoSlices <= 10^7
0 <= cheeseSlices <= 10^7
*/

/*
这是要解一个二元一次方程啊~~~
为叙述方便，t = tomatoSlices， c = cheeseSlices
设最终巨无霸x个，小皇堡y个，根据题意列出方程组：
4x + 2y = t
x + y = c
解得
x = t/2 -c
y = 2c - t/2
需要t为偶数，且最终的x、y非负就行了
时空复杂度都是常数级
*/
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 {
		return nil
	}
	x := tomatoSlices/2 - cheeseSlices
	y := 2*cheeseSlices - tomatoSlices/2
	if x < 0 || y < 0 {
		return nil
	}
	return []int{x, y}
}

/*
翻了翻leetcode题解，另有二分的解法
根据方程组：
x + y = c
4x + 2y = t
可知x的可能范围是[0, t/4]
这个问题使用二分法经典模板比较合适
*/
func numOfBurgers1(tomatoSlices int, cheeseSlices int) []int {
	lo, hi := 0, tomatoSlices/4
	for lo <= hi {
		x := lo + (hi-lo)/2
		y := cheeseSlices - x
		if 4*x+2*y == tomatoSlices {
			return []int{x, y}
		}
		if 4*x+2*y < tomatoSlices {
			lo = x + 1
		} else {
			hi = x - 1
		}
	}
	return nil
}

/*
使用标准库，减少点代码量，标准库的实现其实是模板二
*/
func numOfBurgers2(tomatoSlices int, cheeseSlices int) []int {
	x := sort.Search(tomatoSlices/4+1, func(x int) bool {
		y := cheeseSlices - x
		return 4*x+2*y >= tomatoSlices
	})
	y := cheeseSlices - x
	if 4*x+2*y == tomatoSlices {
		return []int{x, y}
	}
	return nil
}
