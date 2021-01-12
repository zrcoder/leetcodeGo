package split_array

/*
LCP 14. 切分数组
给定一个整数数组 nums ，小李想将 nums 切割成若干个非空子数组，使得每个子数组最左边的数和最右边的数的最大公约数大于 1 。
为了减少他的工作量，请求出最少可以切成多少个子数组。

示例 1：

输入：nums = [2,3,3,2,3,3]

输出：2

解释：最优切割为 [2,3,3,2] 和 [3,3] 。第一个子数组头尾数字的最大公约数为 2 ，第二个子数组头尾数字的最大公约数为 3 。

示例 2：

输入：nums = [2,3,5,7]

输出：4

解释：只有一种可行的切割：[2], [3], [5], [7]

限制：

1 <= nums.length <= 10^5
2 <= nums[i] <= 10^6
*/

/*
首先，直觉有一个贪心策略：
维护一个左指针，对于当前左指针i，从最右开始向左找到第一个右指针j，使得gcd(nums[i], nums[j]) > 1
那么区间[i, j]可以作为一个分组，之后更新i为j+1
可是这样的策略是错的，比如：

2，7，4，。。。，77 （中间省略的m个数都和已经有的数字互质）
显然，因为一开始确定[2,7,4]一组后，后边每个数一组，得到发分组是1+m+1；
但其实7和77可以作为左右端点组成一组，剩下2单独一组，共2组更好
那么如果把这个贪心策略既应用与从左到右的尝试，又从右到左尝试一遍，取两遍得到的最小值行不行呢？

还是不行，看这个例子`
2，7，4，。。。，25，77，5
从左到右贪心一遍，2和4废了7；从右到左贪心一遍，25和5废了77。
两种尝试都把7-77这个情况给漏掉了
*/

/*
动态规划，超时
有两层循环，且里边求gcd也有对数级（相对要求的数字而言）的复杂度

func splitArrayDp(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if gcd(nums[j], nums[i]) > 1 {
				if j == 0 {
					dp[i] = min(dp[i], 1)
				} else {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
		}
	}
	return dp[n-1]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
*/
/*
基本上这种题目直接去求gcd都会超时，需要用上素数筛的方法
基于上面dp的方法，来尝试用筛法解决下
*/
func splitArray(nums []int) int {
	// minPrimes[num]表示任意数字num的最小质因数, 2<= num <= max
	minPrimes := genMinPrimes(getMax(nums))

	n := len(nums)
	dp := make([]int, n)

	// 记录质因数的最佳位置，dp过程中调整
	primePos := make(map[int]int, 0)

	dp[0] = 1
	calFirstNum(nums[0], minPrimes, primePos)

	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + 1
		calNum(i, nums, dp, minPrimes, primePos)
	}
	return dp[n-1]
}

func getMax(nums []int) int {
	r := 0
	for _, v := range nums {
		r = max(r, v)
	}
	return r
}

func genMinPrimes(max int) []int {
	result := make([]int, max+1) // 用切片，map的话会超时
	for v := 2; v <= max; v++ {
		if result[v] != 0 {
			continue
		}
		for times := v; times <= max; times += v {
			if result[times] == 0 {
				result[times] = v
			}
		}
	}
	return result
}

func calFirstNum(num int, minPrimes []int, primePos map[int]int) {
	for num > 1 {
		prime := minPrimes[num]
		primePos[prime] = 0
		for num%prime == 0 {
			num /= prime
		}
	}
}

func calNum(i int, nums, dp, minPrimes []int, primePos map[int]int) {
	num := nums[i]
	for num > 1 {
		prime := minPrimes[num]
		if _, ok := primePos[prime]; !ok {
			primePos[prime] = i
		} else if dp[primePos[prime]] > dp[i-1]+1 {
			primePos[prime] = i
		}
		pos := primePos[prime]
		if pos == 0 {
			dp[i] = 1
			return
		}
		dp[i] = min(dp[i], dp[pos-1]+1)
		for num%prime == 0 {
			num /= prime
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
