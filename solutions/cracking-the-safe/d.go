package cracking_the_safe

import (
	"math"
	"strings"
)

func crackSafe(n int, k int) string {
	total := int(math.Pow(float64(k), float64(n)))
	seen := make(map[int]bool, total)
	high := total / k
	res := strings.Builder{}
	var dfs func(node int)
	dfs = func(node int) {
		for i := 0; i < k; i++ {
			number := node*k + i
			if !seen[number] {
				seen[number] = true
				newNode := number % high
				dfs(newNode)
				res.WriteByte('0' + byte(i))
			}
		}
	}
	dfs(0)
	for i := 1; i < n; i++ {
		res.WriteByte('0')
	}
	return res.String()
}
