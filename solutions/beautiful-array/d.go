/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package beautiful_array

/*
对于某些固定的 N，如果数组 A 是整数 1, 2, ..., N 组成的排列，使得：
对于每个 i < j，都不存在 k 满足 i < k < j 使得 A[k] * 2 = A[i] + A[j]。
那么数组 A 是漂亮数组。
给定 N，返回任意漂亮数组 A（保证存在一个）。

示例 1：
输入：4
输出：[2,1,4,3]

示例 2：
输入：5
输出：[3,1,2,5,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/beautiful-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
官方题解： 自顶向下递归

首先我们可以发现一个不错的性质，如果某个数组[a1,a2,⋯,an] 是漂亮的，
那么对这个数组进行仿射变换，得到的新数组[ka1+b,ka2+b,⋯,kan+b] 也是漂亮的（其中 k ！= 0）。
那么我们就有了一个想法：将数组分成两部分 left 和 right，分别求出一个漂亮的数组，然后将它们进行仿射变换，使得不存在满足下面条件的三元组：
A[k] * 2 = A[i] + A[j], i < k < j；
A[i] 来自 left 部分，A[j] 来自 right 部分。
可以发现，等式 A[k] * 2 = A[i] + A[j] 的左侧是一个偶数，右侧的两个元素分别来自两个部分。
要想等式恒不成立，一个简单的办法就是让 left 部分的数都是奇数，right 部分的数都是偶数。
因此我们将所有的奇数放在 left 部分，所有的偶数放在 right 部分，这样可以保证等式恒不成立。
对于 [1..N] 的排列，left 部分包括 (N + 1) / 2 个奇数，right 部分包括 N / 2 个偶数。
对于 left 部分，我们进行 k = 1/2, b = 1/2 的仿射变换，把这些奇数一一映射到不超过 (N + 1) / 2 的整数。
对于 right 部分，我们进行 k = 1/2, b = 0 的仿射变换，把这些偶数一一映射到不超过 N / 2 的整数。
经过映射，left 和 right 部分变成了和原问题一样，但规模减少一半的子问题，这样就可以使用分治算法解决了。

算法
在 [1..N] 中有 (N + 1) / 2 个奇数和 N / 2 个偶数。
我们将其分治成两个子问题，其中一个为不超过 (N + 1) / 2 的整数，并映射到所有的奇数；另一个为不超过 N / 2 的整数，并映射到所有的偶数。

作者：LeetCode
链接：https://leetcode-cn.com/problems/beautiful-array/solution/piao-liang-shu-zu-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func beautifulArray(N int) []int {
	if N < 1 {
		return nil
	}
	if N == 1 {
		return []int{1}
	}
	var r []int
	for _, v := range beautifulArray((N + 1) / 2) { // 奇数们
		r = append(r, 2*v-1)
	}
	for _, v := range beautifulArray(N / 2) { // 偶数们
		r = append(r, 2*v)
	}
	return r
}

/*
可用一个map储存计算结果，优化递归栈及时间
*/
func beautifulArray2(N int) []int {
	if N < 1 {
		return nil
	}
	m := make(map[int][]int, 0) // 缓存中间计算结果，减少递归
	var helper func(n int) []int
	helper = func(n int) []int {
		if n == 1 {
			return []int{1}
		}
		if v, ok := m[n]; ok {
			return v
		}
		var r []int
		for _, v := range helper((n + 1) / 2) { // 奇数们
			r = append(r, 2*v-1)
		}
		for _, v := range helper(n / 2) { // 偶数们
			r = append(r, 2*v)
		}
		m[n] = r
		return r
	}
	return helper(N)
}

/*
自底向上，循环

漂亮数组有以下性质:
（1）A是一个漂亮数组，如果对A中所有元素添加一个常数，那么Ａ还是一个漂亮数组。
（2）A是一个漂亮数组，如果对A中所有元素乘以一个常数，那么A还是一个漂亮数组。
（3）A是一个漂亮数组，如果删除一些A中元素，那么A还是一个漂亮数组。
（4) A是一个奇数构成的漂亮数组，B是一个偶数构成的漂亮数组，那么A+B也是一个漂亮数组
比如:{1,5,3,7}+{2,6,4,8}={1,5,3,7,2,6,4,8}也是一个漂亮数组。

所以我们假设一个{1,...,m}的数组是漂亮数组，可以通过下面的方式将其规模翻倍，构造漂亮数组{1,...,2m}:
对{1,...,m}中所有的数*2并-1，构成一个奇数漂亮数组A。如{1,3,2,4}，变换为{1,5,3,7}
补齐偶数数组,如上面的{1,5,3,7}, 每个元素+1得到{2,6,4,8}
A+B构成了漂亮数组{1,...,2m}。这个例子里：{1,5,3,7}+{2,6,4,8}={1,5,3,7,2,6,4,8}
因每次翻倍构造，会有些多余的元素（大于N），最后需要剔除这些元素
在N比较小时的构造结果如下：
N == 1：{1};
N == 2：{1, 2};
N == 3：{1, 3, 2, 4} -> {1, 3, 2};  元素4剔除
N == 4：{1, 5, 3, 2, 6, 4} -> {1, 3, 2, 4};  5, 6被剔除

作者：zerorlis-2
链接：https://leetcode-cn.com/problems/beautiful-array/solution/piao-liang-shu-zu-de-yi-xie-xing-zhi-bing-qie-ji-y/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func beautifulArray0(N int) []int {
	if N < 1 {
		return nil
	}
	r := []int{1}
	for len(r) < N {
		n := len(r)
		for i := 0; i < n; i++ { // 得到左侧奇数漂亮数组
			r[i] = 2*r[i] - 1
		}
		for i := 0; i < n; i++ { // 得到右侧偶数漂亮数组
			r = append(r, r[i]+1)
		}
	}
	var result []int
	for _, v := range r {
		if v <= N { // 剔除多余， 只保留在N之内的元素
			result = append(result, v)
		}
	}
	return result
}

// 可以精确控制上边实现里的r和result的大小
func beautifulArray1(N int) []int {
	if N < 1 {
		return nil
	}
	r := make([]int, 2*N-1)
	r[0] = 1
	length := 1
	for length < N {
		for i := 0; i < length; i++ {
			r[i] = 2*r[i] - 1      // 得到左侧奇数漂亮数组
			r[i+length] = r[i] + 1 // 得到右侧偶数漂亮数组
		}
		length *= 2
	}
	if length == N {
		return r[:N]
	}
	result := make([]int, N)
	k := 0
	for _, v := range r {
		if k == N {
			break
		}
		if v <= N { // 剔除多余， 只保留在N之内的元素
			result[k] = v
			k++
		}
	}
	return result
}
