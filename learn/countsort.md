## 计数排序
计数排序并不基于比较来排序，而是基于计数，且其时间复杂度是线性的。<br>
假设n个元素中每一个都是[0,k]区间内的一个数字，可以使用计数排序，当k=O(n)时，时间复杂度为O(n)<br>
### 工作原理
可以先遍历待排序数组，统计到每个元素出现的个数，这里用长度为k+1的数组或者一个哈希表 <br>
```
s:      {2, 5, 3, 0, 2, 3, 0, 3}    k: 5

         0  1  2  3  4  5
count:  {2, 0, 2, 3, 0, 1} // 统计s中元素个数
```
再从0到k遍历刚才的统计数组（哈希表），通过累加得到小于等于各个元素的元素个数 <br>
```
count:  {2, 2, 4, 7, 7, 8} // 对于元素x， 统计小于等于x的个数
```
知道了待排序数组中任意元素x，小于等于x的元素个数，也就知道了**x排序后该在的位置**<br>
开辟一个和待排序数组同样大小的数组，存储最终的结果;如果有重复元素，简单处理下即可
### 实现
```go
// 计数排序s，假设s中元素值域为[0, k], 非原地的计数排序
func countSort(s []int) []int {
	if len(s) == 0 {
		return s
	}
	k := s[0]
	for _, v := range s {
		if v > k {
			k = v
		}
	}
	count := make([]int, k+1)
	for _, v := range s { // 统计s中每个元素出现的个数
		count[v]++
	}
	for i := 1; i <= k; i++ { // 统计s中每个元素x，小于等于x的元素个数，即统计排序后x应该出现的位置
		count[i] += count[i-1]
	}
	result := make([]int, len(s))
	for _, num := range s {
		index := count[num] - 1
		result[index] = num
		count[num]-- // 如果有重复的元素i，下一次插入的位置是当前插入位置的前一位
	}
	return result
}
```
count使用了数组，天然约束s中元素非负；如果s中有负数，可以调整映射关系。<br>
显然最小元素min应该对应count的0索引，元素i对应count的i+abs(min)。<br>
更普遍地，count数组可以用一个哈希表替代：
```go
// 计数排序s，假设s中元素值域为[min, max], 元素可以有负数且比较分散，非原地的计数排序
func countSort(s []int) []int {
	if len(s) == 0 {
		return s
	}
	count := make(map[int]int, 0)
	min, max := s[0], s[0]
	for _, v := range s { // 统计s中每个元素出现的个数
		count[v]++
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	for i := min+1; i <= max; i++ { // 统计s中每个元素x，小于等于x的元素个数，即统计排序后x应该出现的位置
		count[i] += count[i-1]
	}
	result := make([]int, len(s))
	for _, num := range s {
		index := count[num] - 1
		result[index] = num
		count[num]-- // 如果有重复的元素i，下一次插入的位置是当前插入位置的前一位
	}
	return result
}
```
假设数组大小为n， 值域为[Min, Max],设Max-Min+1=k <br>
计数排序的时间复杂度是O(k + n),<br>
空间复杂度O(k + n)，即count数组（map）的长度+结果数组result的长度<br>
如果待排序数组的元素不够分散，即n远大于k，计数排序的效率将不太好，比如待排序数组是{1, 1, 1000, 1, 1}<br>
无论时间空间，都有很大浪费。有办法优化吗？<br>
无论如何，都要遍历待排序数组，且需要一个结果数组。时间空间复杂度里的n是优化不了的，能优化的就是count的时空复杂度~<br>
至少count用哈希表，空间是可以优化的，像上面的例子，count的key只存1和1000，大小是2，也就是待排序数组中不同元素的个数<br>
将“ 统计s中每个元素x，小于等于x的元素个数，即统计排序后x应该出现的位置”这部分做一修改
```go
	sum := 0
	for i := min; i <= max; i++ { // 统计s中每个元素x，小于等于x的元素个数，即统计排序后x应该出现的位置
		if c, found := count[i]; found { // 只有count里存在i的时候才统计
			sum += c
			count[i] = sum
		}
	}
```
显然，待排序数组元素不够分散情况下的哈希表空间复杂度被降下来了；时间复杂度的话，怎么办？<br>
Go标准库并没有提供SortedMap之类的数据结构，我们可以做个轮子：<br>
比较简单的实现就是另外开辟一个数组切片，或者用一个小顶堆，遍历哈希表，将key排序（借用其他的排序算法）；使用哈希表时从切片或堆遍历就行<br>
这里封装一个函数：
```go
func rangeAsSorted(m map[int]int, f func(k, v int)) {
	keys := make([]int, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	for _, key := range keys {
		f(key, m[key])
	}
}
```
可以看到，为了能按顺序遍历哈希表，我们额外引入了一个切片，并对切片做了排序；<br>
哈希表比较小的情况下没什么，但如果哈希表很大,换句话说在这个问题中元素很分散，<br>
则空间复杂度还好，假设len(count)= c（已经对count空间做了优化，其大小不会大于n），空间复杂度为O(n+c)=O(n)，但是时间复杂度却是O(n+clg(c))<br>
可以用一个堆来代替切片，但是不会好太多。—— 鱼与熊掌不可兼得，所以计数排序要分情况：<br>
元素非常分散的话，就用开始的实现；元素非常聚焦的话， count部分的时间空间也可以优化。当然无论如何，我们可以控制count的空间大小与待排序数组相同<br>
假设数组元素比较聚焦，计数排序变成了这样：
```go
// 计数排序s，假设s中元素值域为[min, max], 元素可以有负数且比较聚焦，非原地的计数排序
func countSort(s []int) []int {
	if len(s) == 0 {
		return s
	}
	count := make(map[int]int, 0)
	for _, v := range s { // 统计s中每个元素出现的个数
		count[v]++
	}
	sum := 0
	rangeAsSorted(count, func(k, v int) {
		sum += v
		count[k] = sum
	})
	result := make([]int, len(s))
	for _, num := range s {
		index := count[num] - 1
		result[index] = num
		count[num]-- // 如果有重复的元素i，下一次插入的位置是当前插入位置的前一位
	}
	return result
}

func rangeAsSorted(m map[int]int, f func(k, v int)) {
	keys := make([]int, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	for _, key := range keys {
		f(key, m[key])
	}
}
```
## 应用
[1365. 有多少小于当前数字的数字](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number)
