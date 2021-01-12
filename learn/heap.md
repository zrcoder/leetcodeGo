## 怎么在切片的指定索引处插入一个元素？
假设有如下切片：
```go
s := []int{2, 8 , 6, 9}
```
怎么在索引2处插入新元素100？变成：
```go
{2, 8, 100, 6, 9}
```
### 朴素实现
```go
// 1.s扩容，长度加一
s = append(s, 0) // 追加的元素值是什么不重要，只是为了扩容
// 2. 从待插入索引2开始，元素一一后移
for i := len(s)-1; i > 2; i-- { // 为了防止覆盖，从后向前遍历，每一位的值等于前一位的值
	s[i] = s[i-1]
}
// 3.将要插入的100赋给s[2]
s[2] = 100
```
一般化，不妨把这个逻辑提取成一个函数
```go
// 在切片s的索引i处插入新元素v
func insert(s *[]int, i, v int) {
	*s = append(*s, 0)
	for j := len(*s) - 1; j > i; j-- {
		(*s)[j] = (*s)[j-1]
	}
	(*s)[i] = v
}
```
注意到函数接收的是切片的指针，有兴趣的话考虑下去掉*会怎么样？为什么会这样？<br>
当前要调用的话是这样：
```go
insert(&s, 2, 100)
```
*、&看起来有点吓人，可以这样修改下：让insert函数返回修改后的切片:
```go
func insert(s []int, i, v int) []int {
	s = append(s, 0)
	for j := len(s) - 1; j > i; j-- {
		s[j] = s[j-1]
	}
	s[i] = v
	return s
}
```
当然调用时变成这样了：
```go
s = insert(s, 2, 100)
```
至此我们完成了朴素实现的insert()
### 用上内置函数copy(),减少代码量
将从i开始的元素一一后移，这里可以用内置函数copy，减少那个循环：
```go
func insert(s []int, i, v int) []int {
	s = append(s, 0)
	_ = copy(s[i+1:], s[i:len(s)-1])
	s[i] = v
	return s
}
```
### 不用循环，不用copy，纯用append的实现
有一次看到别人实现这个功能的代码，惊艳了，先上代码：
```go
func insert(s []int, i, v int) []int {
	r := s[:i:i]
	r = append(r, v)
	r = append(r, s[i:]...)
	return r
}
```
也很容易理解，先获取i前边的部分，再追加v，再追加包括i在内的后边的部分<br>
有个细节需要注意，第一步为什么不写成s[:i]，明确要求r的容量为i是为什么？思考题~
基于上边的实现，甚至可以用一行搞定：
```go
func insert(s []int, i, v int) []int {
	return append(append(s[:i:i], v), s[i:]...)
}
```
简洁！同时也很明朗：前半部分+v+后半部分<br>
值得一提的是最后这个实现甚至允许i为s的长度，不像前边的实现，i为s长度时会索引越界。