# [1. 两数之和](https://leetcode-cn.com/problems/two-sum)

## 题目

```text
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

## 分析

* 朴素实现， O(n^2)复杂度

```
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j ++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```
* 经典哈希表解法<br>
时间O(n), 空间O(n)的实现<br>
在遍历中，如果i在结果中，target-nums[i]必在nums中，否则i不满足要求<br>
为了迅速查找target-nums[i]是否在nums中，我们可以构造一个map，其键为nums里的元素，值为元素的索引<br>
构造需要遍历一遍nums，注意构造过程中，考虑有nums元素重复的情况<br>
查找另需一遍遍历<br>
```
func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, element := range nums {
		index[element] = i
	}
	for i, element := range nums {
		if j, found := index[target-element]; found {
			return []int{i, j}
		}
	}
	return nil
}
```
1. 进一步优化<br>
实际上可以边构建map，边做查找，总体只需遍历一遍<br>
```
func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, element := range nums {
		if j, found := index[target-element]; found {
			return []int{j, i} // not {i, j}, but {j, i}; let's think, j < i
		}
		index[element] = i
	}
	return nil
}

```
## 拓展
如果数组是已经排好序的呢？<br>
可以从两边往中间凑， 时间O(n), 不用额外空间~
```
func twoSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		}
		if sum < target {
			i ++
		} else {
			j --
		}
	}
	return nil
}
```
[参考实现](imp.go)及[测试用例](imp_test.go)