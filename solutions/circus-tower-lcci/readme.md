## [面试题 17.08. 马戏团人塔](https://leetcode-cn.com/problems/circus-tower-lcci/)
```text
有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。
出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。
已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

示例：
输入：height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
输出：6
解释：从上往下数，叠罗汉最多能叠 6 层：(56,90), (60,95), (65,100), (68,110), (70,150), (75,190)

提示：
height.length == weight.length <= 10000
```

### 解题思路
总体思路同[[406] 根据身高重建队列](../queue-reconstruction-by-height/d.go)，先粗排再细排

细排的思路同[[300] 最长上升子序列](../longest-increasing-subsequence/d.go)中贪心+二分搜索的解法

* 1.粗处理

排序，依据为：身高降序（这里就是人塔从下向上身高递减）

身高相等则体重轻的在前边（从下向上体重升序）——`问题1`：这里不符而直觉，直觉应该体重也是降序才对，看完下边步骤再回来看

* 2.细处理

遍历排序结果，将当前的人`p`放入到结果数组`result`中

当然`result`的所有元素符合题意，即前一个人的身高体重均比后一个人大

从前向后搜索`result`，找到第一个不应该排在`p`之前的人（即身高或体重不大于`p`的身高体重）

2.1 没找到，直接把`p`追加到`result`结尾

2.2 找到的话，把这个人排除，替换成``p——这是这个问题的关键所在，是个贪心策略，可以想一想为什么这样是合理的

最后的`result`数组即为所求；另外在搜索`result`的时候，用二分法会更快

* 回头看`问题1`：

考虑有身高相同的几个人，显然最终结果这几个人里最多有一个入选（也可能都不会入选），贪心策略，根据`步骤2`的处理，应该保留体重最大的人入选
所以一开始粗排，身高相等的情况下，体重升序，是为了2细排的时候能保留体重最重的那个人

### 初步实现
辅助代码
```go
type Person struct {
	height, weight int
}
```
主体代码
```go
func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	persons := make([]Person, n)
	for i := range persons {
		persons[i].height = height[i]
		persons[i].weight = weight[i]
	}
	sort.Slice(persons, func(i, j int) bool {
		// 身高高的在前边，身高相等则体重轻的在前边
		if persons[i].height == persons[j].height {
			return persons[i].weight < persons[j].weight
		}
		return persons[i].height > persons[j].height
	})
	var result []Person
	for _, p := range persons {
		// 在结果中找到第一个p不能叠在上面的人, 二分法
		j := sort.Search(len(result), func(i int) bool {
			c := result[i]
			return c.height <= p.height || c.weight <= p.weight
		})
		// 将第j个人替换成p
		if j == len(result) {
			result = append(result, p)
		} else {
			result[j] = p
		}
	}
	return len(result)
}
```
`步骤1`的排序时间复杂度为`O(n * lgn)`, `步骤2`遍历每个元素，对每个元素，在结果中用二分法找到应该替换的位置，时间复杂度`O(n * lgn)`
综合时间复杂度`O(n * lgn)`
`persons`数组和`result`数组为额外开辟的空间，空间复杂度`O(n)`

### 小做优化
上面的`步骤2`，可以在`persons`原地调整，节省`result`数组
```go
    // 粗处理的代码略
    k := 0
    for _, p := range persons {
        // 在结果中找到第一个不能叠在p上面的人, 二分法
        j := sort.Search(k, func(i int) bool {
            c := persons[i]
            return c.height <= p.height || c.weight <= p.weight
        })
        // 将这个个人替换成p
        persons[j] = p
        if j == k {
            k++
        }
    }
    return k
```
[参考实现](d.go)
另有一个dp解法，详见[[354] 俄罗斯套娃信封问题](../russian-doll-envelopes/d.go)