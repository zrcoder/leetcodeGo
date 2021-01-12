package least_number_of_unique_integers_after_k_removals

import "sort"

/*
1481. 不同整数的最少数目
https://leetcode-cn.com/problems/least-number-of-unique-integers-after-k-removals

给你一个整数数组 arr 和一个整数 k 。现需要从数组中恰好移除 k 个元素，请找出移除后数组中不同整数的最少数目。

示例 1：
输入：arr = [5,5,4], k = 1
输出：1
解释：移除 1 个 4 ，数组中只剩下 5 一种整数。

示例 2：
输入：arr = [4,3,1,1,3,3,2], k = 3
输出：2
解释：先移除 4、2 ，然后再移除两个 1 中的任意 1 个或者三个 3 中的任意 1 个，最后剩下 1 和 3 两种整数。

提示：
1 <= arr.length <= 10^5
1 <= arr[i] <= 10^9
0 <= k <= arr.length
*/
func findLeastNumOfUniqueInts(arr []int, k int) int {
	count := map[int]int{}
	for _, v := range arr {
		count[v]++
	}
	frequencies := make([]int, 0, len(count))
	for _, v := range count {
		frequencies = append(frequencies, v)
	}
	sort.Ints(frequencies)
	res := len(frequencies)
	for _, v := range frequencies {
		if v > k {
			break
		}
		k -= v
		res--
	}
	return res
}

/*
java 代码

	public int findLeastNumOfUniqueInts(int[] arr, int k) {
        Map<Integer, Integer> count = new HashMap<>();
        for (int num : arr) {
            count.put(num, count.getOrDefault(num, 0) + 1);
        }
        int[] frequencies = new int[count.size()];
        int i = 0;
        for (int c : count.values()) {
            frequencies[i] = c;
            i++;
        }
        Arrays.sort(frequencies);
        int res = frequencies.length;
        for (int c : frequencies) {
            if (c > k) {
                break;
            }
            k -= c;
            res--;
        }
        return res;
    }
*/
