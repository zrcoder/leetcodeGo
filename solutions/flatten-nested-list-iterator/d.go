package flatten_nested_list_iterator

/*
341. 扁平化嵌套列表迭代器
给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。

列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。



示例 1:

输入: [[1,1],2,[1,1]]
输出: [1,1,2,1,1]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
示例 2:

输入: [1,[4,[6]]]
输出: [1,4,6]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
*/

/*
简单实现下NestedInteger结构体，解决编译问题；不是这个问题重点
*/
// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val  interface{}
	list []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (i NestedInteger) IsInteger() bool {
	return i.val != nil
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (i NestedInteger) GetInteger() int {
	return i.val.(int)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (i NestedInteger) GetList() []*NestedInteger {
	return i.list
}

// 在初始化Iterator的时候一次性扁平化
type NestedIterator1 struct {
	s []int
}

func Constructor1(nestedList []*NestedInteger) *NestedIterator1 {
	var flattern func(nestedList []*NestedInteger) []int
	flattern = func(nestedList []*NestedInteger) []int {
		var s []int
		for _, v := range nestedList {
			if v.IsInteger() {
				s = append(s, v.GetInteger())
			} else {
				s = append(s, flattern(v.GetList())...)
			}
		}
		return s
	}
	return &NestedIterator1{s: flattern(nestedList)}
}

func (it *NestedIterator1) Next() int {
	r := it.s[0]
	it.s = it.s[1:]
	return r
}

func (it *NestedIterator1) HasNext() bool {
	return len(it.s) > 0
}

// 不在初始化的时候一次性扁平化，而是在具体调用HasNext或Next的时候做部分扁平化处理
type NestedIterator struct {
	s []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{s: nestedList}
}

func (it *NestedIterator) Next() int {
	r := it.s[0].GetInteger()
	it.s = it.s[1:]
	return r
}

func (it *NestedIterator) HasNext() bool {
	for len(it.s) > 0 && !it.s[0].IsInteger() {
		var tmp []*NestedInteger
		first := it.s[0].GetList()
		it.s = it.s[1:]
		for _, v := range first {
			tmp = append(tmp, v)
		}
		tmp = append(tmp, it.s...)
		it.s = tmp
	}
	return len(it.s) > 0
}
