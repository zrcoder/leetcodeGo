/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package ugly_number

import (
	"container/heap"
	"math"
)

/*
263.丑数 https://leetcode-cn.com/problems/ugly-number
编写一个程序判断给定的数是否为丑数。
丑数就是只包含质因数 2, 3, 5 的正整数，同时约定1是丑数

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
不断除以2、3、5最后得到1的话就是丑数~
*/
func isUgly1(num int) bool {
	if num <= 0 {
		return false
	}
	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1
}

/*
可以有个小优化, 先将num里所有的3、和5除掉；再判断剩余的数字是不是2的幂，这个可以在常数时间计算：n & (n-1) == 0就是2的幂
*/
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num&(num-1) == 0
}

/*
264. 丑数 II https://leetcode-cn.com/problems/ugly-number-ii
编写一个程序，找出第 n 个丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
朴素实现：
从1开始增长，判断当前数字是否为丑数
（用例超时，500 / 596 个通过测试用例）
*/
func nthUglyNumber1(n int) int {
	result := 1
	for n > 0 {
		if isUgly(result) {
			n--
		}
		result++
	}
	return result - 1 // 循环里多加了一次
}

/*
对朴素解法做一个优化：
受到素数筛选法启发，我们尝试用筛选法
所有的丑数都是之前的丑数乘以2、3、5生成的

（用例超时，500 / 596 个通过测试用例——成绩和朴素实现一样~~~）
*/
func nthUglyNumber2(n int) int {
	result := 1
	set := make(map[int]struct{})
	set[result] = struct{}{}
	for n > 0 {
		if _, ok := set[result]; ok {
			n--
			set[result*2] = struct{}{}
			set[result*3] = struct{}{}
			set[result*5] = struct{}{}
		}
		result++
	}
	return result - 1 // 循环里多加了一次
}

/*
其实每次result增加1过于老实，让result为set中最小元素会比较快；
为了能迅速找到最小值，可以用小顶堆代替set
还需要注意一点，当最小的元素出堆后，要把和它相等的元素都出堆。

其实把set改成小顶堆后，result得到了优化，但是如堆却变慢了，
综合来看，heap和set的表现应该不相上下，但是实际测试这次通过了全部用例
——在leetcode的用例里这个解法应该算是比上边的解法小有优化
*/
func nthUglyNumber3(n int) int {
	pq := &IntHeap{}
	result := 1
	heap.Push(pq, result)
	for ; n > 0; n-- {
		result = heap.Pop(pq).(int)
		// 去除重复;如丑数6，可以是2*3， 也可以是3*2
		for pq.Len() > 0 && result == (*pq)[0] {
			_ = heap.Pop(pq)
		}
		heap.Push(pq, result*2)
		heap.Push(pq, result*3)
		heap.Push(pq, result*5)
	}
	return result
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

/*
可以再优化下，每次入堆的时候就保证不要入堆里已有的元素
其他语言如java有个TreeSet可用，我们这里可以再堆外再用一个set来去重
实际测试收效甚微~~~想一想和上面单纯用个堆，在出堆的时候保证相同最小元素都出堆的做法没有太大区别，反而空间上多了一个set
这里如果实现一个treeset会怎么样？可以预见优化不会太多，这里不尝试了。
*/
func nthUglyNumber4(n int) int {
	pq := &IntHeap{}
	set := make(map[int]struct{})
	result := 1
	heap.Push(pq, result)
	set[result] = struct{}{}
	primes := []int{2, 3, 5}
	for ; n > 0; n-- {
		result = heap.Pop(pq).(int)
		for _, v := range primes {
			newUgly := result * v
			if _, ok := set[newUgly]; !ok {
				set[newUgly] = struct{}{}
				heap.Push(pq, newUgly)
			}
		}
	}
	return result
}

/*
动态规划

我们先模拟手写丑数的过程
1打头，1乘2 1乘3 1乘5，现在是{1,2,3,5}
轮到2，2乘2 2乘3 2乘5，现在是{1,2,3,4,5,6,10}
手写的过程和采用小顶堆的方法很像，但是怎么做到提前排序呢

小顶堆的方法是先存再排，dp的方法则是先排再存
我们设3个指针p_2,p_3,p_5
代表的是第几个数的2倍、第几个数3倍、第几个数5倍
动态方程dp[i]=min(dp[p_2]*2,dp[p_3]*3,dp[p_5]*5)
小顶堆是一个元素出来然后存3个元素
动态规划则是标识3个元素，通过比较他们的2倍、3倍、5倍的大小，来一个一个存

作者：LZH_Yves
链接：https://leetcode-cn.com/problems/ugly-number-ii/solution/bao-li-you-xian-dui-lie-xiao-ding-dui-dong-tai-gui/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

时空复杂度都是O(n)
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		t2, t3, t5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(t2, t3, t5)
		// t2, t3, t5不一定完全不相等
		if dp[i] == t2 {
			p2++
		}
		if dp[i] == t3 {
			p3++
		}
		if dp[i] == t5 {
			p5++
		}
	}
	return dp[n-1]
}

func min(a, b, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}

/*
1201. 丑数 III https://leetcode-cn.com/problems/ugly-number-iii
请你帮忙设计一个程序，用来找出第 n 个丑数。
丑数是可以被 a 或 b 或 c 整除的 正整数。
*/
