## [面试题 05.07. 配对交换](https://leetcode-cn.com/problems/exchange-lcci/)
配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位0与位1交换，位2与位3交换，以此类推）。

示例1:

 输入：num = 2（或者0b10）
 输出 1 (或者 0b01)
 
示例2:

 输入：num = 3
 输出：3
 
提示:
num的范围在[0, 2^30 - 1]之间，不会发生整数溢出。

### 解题思路
1. 字符串转换
2. 位运算
### 字符串转换
转成二进制的字符数组再做操作，最后转换回数字
```go
func exchangeBits(num int) int {	
	s := strconv.FormatUint(uint64(num), 2)
	if len(s) % 2 == 1 {
		s = "0" + s
	}
	b := []byte(s)
	for i := len(b)-1; i >= 0; i-=2 {
		b[i], b[i-1] = b[i-1], b[i]
	}
	r, _ := strconv.ParseUint(string(b), 2, 64)
	return int(r)
}
```
### 位运算
1.可以通过位运算将`num`的所有奇数位置为`0`得到`num`的偶数位数字`even`：

因num在`[0,2^30-1]`范围，可以用一个`30`位偶数位为`1`奇数位为`0`的数字与`num`做与运算得到`even`的值

2.同理可以得到`num`的奇数位数字`odd`

3.最后将`even >> 1`和`odd << 1`相加或做或运算得到结果
```go
func exchangeBits(num int) int {
	evenMask := 0b101010101010101010101010101010101010101010101010101010101010
	oddMask :=  0b010101010101010101010101010101010101010101010101010101010101
	even, odd := num&evenMask, num&oddMask
	return even>>1 | odd<<1
}
```
值得一提的是，较低的Go版本不支持evenMask和oddMask里`0bxxxxx`这样的语法，

可以改成十进制数值`768614336404564650`和`384307168202282325`

不过这样可读性不好