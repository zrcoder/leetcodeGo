/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package count_primes

/*
204. Count Primes
https://leetcode.com/problems/count-primes

统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/

//朴素实现，超时；时间复杂度O(n*sqrt(n)), 空间复杂度O（1）
func countPrimes1(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

/* 空间换时间
用一个bool数组来存某个数字是否素数；用类似动态规划的方法标记这个数组
如：
标记2 为素数，然后标记2倍数（4，6，8。。。）为非素数
标记3为素数，然后标记3但倍数为非素数
。。。
不断重复以上的过程
*/
func countPrimes2(n int) int {
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ { // 也可以认为false代表素数，这里就不用做这个标记里
		isPrime[i] = true
	}
	for i := 2; i < n; i++ { // 可以优化，上界到sqrt（n）即可；即i*i < n
		if !isPrime[i] {
			continue
		}
		for j := i + i; j < n; j += i { // j为i的倍数; 可以优化，j到初始值从i*i开始
			isPrime[j] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count

}

// 初始化及循环条件优化后的实现：
func countPrimes3(n int) int {
	isComposite := make([]bool, n) // 默认全部不是合数，即全部是素数
	for i := 2; i*i < n; i++ {
		if isComposite[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isComposite[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			count++
		}
	}
	return count
}

// 可以用bitset替换isPrime数组，优化空间；bitset的空间是bool数组的1/8
func countPrimes(n int) int {
	isComposite := NewBitsetWithSize(n)
	for i := 2; i*i < n; i++ {
		if isComposite.Get(i) {
			continue
		}
		for j := i * i; j < n; j += i {
			isComposite.Set(j) // 标记为非素数
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isComposite.Get(i) {
			count++
		}
	}
	return count
}

type BitSet []byte

func NewBitsetWithSize(size int) BitSet {
	return make([]byte, 1+(size-1)/8)
}

// Set true at the index
func (bs BitSet) Set(index int) {
	bs[index/8] |= 1 << uint(index%8)
}

// Returns the bool value at the index
func (bs BitSet) Get(index int) bool {
	return bs[index/8]&(1<<uint(index%8)) != 0
}
