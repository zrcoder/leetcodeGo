/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package water_and_jug_problem

/*
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
示例 1: (From the famous "Die Hard" example)
输入: x = 3, y = 5, z = 4
输出: True

示例 2:
输入: x = 2, y = 6, z = 5
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/water-and-jug-problem
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
裴蜀定理:
得名于法国数学家艾蒂安·裴蜀，说明了对任何整数a、b和它们的最大公约数d，关于未知数x和y的线性不定方程（称为裴蜀等式）：
若a,b是整数,且gcd(a,b)=d，一个比较明显的结论是：对于任意的整数x,y,ax+by都一定是d的倍数
特别地，一定存在整数x,y，使ax+by=d成立， 注意这里可能有负整数
它的一个重要推论是：a,b互质的充要条件是存在整数x,y使ax+by=1.

这个问题相当于找到整数a，b使得ax+by=z
当且仅当z是a和b的最大公约数的倍数
*/
func canMeasureWater(x int, y int, z int) bool {
	return z == 0 || (x+y >= z && z%gcd(x, y) == 0)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
