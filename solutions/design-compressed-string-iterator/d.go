/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design_compressed_string_iterator

/*
对于一个压缩字符串，设计一个数据结构，它支持如下两种操作： next 和 hasNext。

给定的压缩字符串格式为：每个字母后面紧跟一个正整数，这个整数表示该字母在解压后的字符串里连续出现的次数。

next() - 如果压缩字符串仍然有字母未被解压，则返回下一个字母，否则返回一个空格。
hasNext() - 判断是否还有字母仍然没被解压。

注意：

请记得将你的类在 StringIterator 中 初始化 ，因为静态变量或类变量在多组测试数据中不会被自动清空。更多细节请访问 这里 。

示例：

StringIterator iterator = new StringIterator("L1e2t1C1o1d1e1");

iterator.next(); // 返回 'L'
iterator.next(); // 返回 'e'
iterator.next(); // 返回 'e'
iterator.next(); // 返回 't'
iterator.next(); // 返回 'C'
iterator.next(); // 返回 'o'
iterator.next(); // 返回 'd'
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 'e'
iterator.hasNext(); // 返回 false
iterator.next(); // 返回 ' '

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-compressed-string-iterator
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
1 最朴素对实现是在初始化时还原字符串，不过这样可能内存占用太大
2 可用chars和nums两个数组记录原字符串里的字符及每个字符的出现次数
```
type StringIterator struct {
  chars []byte
  nums []int
  index int
}

func Constructor(compressedString string) StringIterator {
  r := StringIterator{}
  i := 0
  for i < len(compressedString) {
	  ch := compressedString[i]
	  r.chars = append(r.chars, ch)
	  i ++
	  num := 0
	  for i< len(compressedString) && compressedString[i] >= '0' && compressedString[i] <= '9' {
		  num = 10 * num + int(compressedString[i]-'0')
		  i ++
	  }
	  r.nums = append(r.nums, num)
  }
  return r
}

func (si *StringIterator) Next() byte {
  if si.HasNext() {
	  r := si.chars[si.index]
	  si.nums[si.index] --
	  if si.nums[si.index] == 0 {
		  si.index ++
	  }
	  return r
  }
  return ' '
}

func (si *StringIterator) HasNext() bool {
  return len(si.chars) > si.index
}
```
*/
/*
3 不对原字符串做处理；在实际对操作里处理
*/
type StringIterator struct {
	s     string
	num   int
	ch    byte
	index int
}

func Constructor(compressedString string) StringIterator {
	return StringIterator{s: compressedString}
}

func (si *StringIterator) Next() byte {
	if !si.HasNext() {
		return ' '
	}
	if si.num == 0 {
		si.ch = si.s[si.index]
		si.index++
		for si.index < len(si.s) && si.s[si.index] >= '0' && si.s[si.index] <= '9' {
			si.num = 10*si.num + int(si.s[si.index]-'0')
			si.index++
		}
	}
	si.num--
	return si.ch
}

func (si *StringIterator) HasNext() bool {
	return len(si.s) > si.index || si.num > 0
}
