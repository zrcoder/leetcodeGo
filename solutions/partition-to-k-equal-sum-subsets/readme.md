许多问题，缩小视野、限于特例就会变得困难，如果跳出来发现一般形式，问题反而更明朗，更容易解决。<br>
以下三个leetcode问题，本质是一个，让我们从最普遍的问题开始逆序遍历这几个问题：<br>
`题目来源：力扣（LeetCode）;著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。`
## [698. 划分为k个相等的子集](https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets)
```text
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。


注意:

1 <= k <= len(nums) <= 16
0 < nums[i] < 10000
```
穷举搜索，适时回溯：
```text
对于 nums 中的每个数字，可以将其添加到 k 个子集中的一个，只要该子集的和不会超过目标值;
在不满足预期的时候做好回溯。
```
程序框架：
```go
func canPartitionKSubsets(nums []int, k int) bool {
	sum, max := 0, 0 // 注意到输入限制为非负，且和不会溢出
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / k // 尝试把所有元素放到k个组中，每组元素和为target
	if sum%k != 0 || max > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking(k, 0, 0, target, nums, used)
}
```
核心的穷举回溯函数backTracking()
```go
/*
尝试把所有元素放到k个组中，每组元素和为target
currSubSum 记录当前组累加的结果，start指明从nums的哪个位置开始
*/
func backTracking(k, currSubSum, start, target int, nums []int, used []bool) bool {
	if k == 0 { // 说明所有数字都放入了对应组
		return true
	}
	if currSubSum == target { // 已经构建了若干组
		// 构建下一组
		return backTracking(k-1, 0, 0, target, nums, used)
	}
	for i := start; i < len(nums); i++ {
		if !used[i] && currSubSum+nums[i] <= target {
			used[i] = true // 当前值放入组
			if backTracking(k, currSubSum+nums[i], i+1, target, nums, used) {
				return true
			}
			used[i] = false // 说明将当前值放入一组不能得到结果，回溯
		}
	}
	return false
}
```

## [473. 火柴拼正方形](https://leetcode-cn.com/problems/matchsticks-to-square)
```text
还记得童话《卖火柴的小女孩》吗？
现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。
不能折断火柴，可以把火柴连接起来，并且每根火柴都要用到。
输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。

示例 1:
输入: [1,1,2,2,2]
输出: true
解释: 能拼成一个边长为2的正方形，每边两根火柴。

示例 2:
输入: [3,3,3,3,4]
输出: false
解释: 不能用所有火柴拼成一个正方形。

注意:
给定的火柴长度和在 0 到 10^9之间。
火柴数组的长度不超过15。
```
即问题698中k为4的特例 [solution](partion4.go)
## [416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum)
```text
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.
```
即问题698中k为2的特例，但直接套框架会有部分用例超时，可以预先将数组从大到小排序来优化。<br>
另有一个01背包的解法 [solution](partition2.go)