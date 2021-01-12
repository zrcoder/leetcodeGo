# 位运算
位运算巧用，能起到事半功倍的效果<br>
以下问题可借助哈希表解决，但用位运算更巧妙
## [136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/)
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现**两次**。
找出那个只出现了一次的元素。
```go
func singleNumber(nums []int) int {
    result := 0
    for _, v := range nums {
        result ^= v
    }
    return result
}
```
## [137. 只出现一次的数字 II](https://leetcode-cn.com/problems/single-number-ii/)
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了**三次**。
找出那个只出现了一次的元素。
```go
func singleNumber(nums []int) int {
    result := 0
    for i := uint(0); i < 64; i++ {
        ones := 0 // 统计第i位1出现的个数
        for _, v := range nums {
            ones += (v >> i) & 1
        }
        result |= (ones%3) << i
    }
    return result
}
```
## [260. 只出现一次的数字 III](https://leetcode-cn.com/problems/single-number-iii/)
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。<br> 
找出只出现一次的那两个元素。
```go
func singleNumber(nums []int) []int {
	t := 0
	for _, v := range nums {
		t ^= v
	}
	// 最后一个1的位置, 或写作 t & (-t)
	// 假设要求的两个数分别是x和y，最后一个1，要吗来自x，要吗来自y
	diff := (t & (t - 1)) ^ t
	x := 0
	for _, v := range nums {
		if diff&v != 0 {
			x ^= v
		}
	}
	return []int{x, t ^ x}
}
```