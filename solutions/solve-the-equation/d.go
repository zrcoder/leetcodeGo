/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package equation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
640. 求解方程 https://leetcode-cn.com/problems/solve-the-equation

求解一个给定的方程，将x以字符串"x=#value"的形式返回。
该方程仅包含'+'，' - '操作，变量 x 和其对应系数。

如果方程没有解，请返回“No solution”。
如果方程有无限解，则返回“Infinite solutions”。
如果方程中只有一个解，要保证返回值 x 是一个整数。

示例 1：
输入: "x+5-3+x=6+x-2"
输出: "x=2"

示例 2:
输入: "x=x"
输出: "Infinite solutions"

示例 3:
输入: "2x=x"
输出: "x=0"

示例 4:
输入: "2x+3x-6x=x+2"
输出: "x=-1"

示例 5:
输入: "x=x+2"
输出: "No solution"
*/
func solveEquation(equation string) string {
	// 分别对方程两侧统计x的系数和及数字的和，处理成ax+b=cx+d的形式
	exps := strings.Split(equation, "=")
	a, b := help(exps[0])
	c, d := help(exps[1])
	if a-c == 0 && d-b == 0 {
		return "Infinite solutions"
	}
	if a-c == 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", (d-b)/(a-c))
}

/*
以下老老实实实现，从前向后遍历字符，代码复杂，要考虑的比较多
比较坑的是类似8+0x-2=8这样的用例，未知数前边一个0~~~~~~
以下代码处理不了这种情况
*/
func help0(exp string) (int, int) {
	// 异常预处理
	exp = strings.ReplaceAll(exp, "+0x", "")
	exp = strings.ReplaceAll(exp, "-0x", "")
	if strings.HasPrefix(exp, "0x") {
		exp = exp[2:]
	}
	xCount, nums := 0, 0
	flag := 1
	currNum := 0
	for i := range exp {
		v := exp[i]
		switch {
		case v == '+':
			if i > 0 && exp[i-1] != 'x' {
				nums += currNum * flag
			}
			flag = 1
			currNum = 0
		case v == '-':
			if i > 0 && exp[i-1] != 'x' {
				nums += currNum * flag
			}
			flag = -1
			currNum = 0
		case isNum(v):
			if i > 0 && exp[i-1] == 'x' {
				flag = 1
			}
			currNum = currNum*10 + int(v-'0')
		case v == 'x':
			if i == 0 {
				xCount = 1
			} else {
				if currNum == 0 {
					xCount += flag
				} else {
					xCount += flag * currNum
				}
				currNum = 0
			}
		}
	}
	return xCount, nums + flag*currNum
}

func isNum(v byte) bool {
	return v >= '0' && v <= '9'
}

func help(exp string) (int, int) {
	// 给所有"-"前边加一个"+"，然后用"+"切分整个表达式，方便处理
	exp = strings.ReplaceAll(exp, "-", "+-")
	s := strings.Split(exp, "+")
	xCount, num := 0, 0
	for _, v := range s {
		if len(v) > 0 && v[len(v)-1] == 'x' {
			if v == "x" || v == "+x" {
				xCount++
			} else if v == "-x" {
				xCount--
			} else {
				n, _ := strconv.Atoi(v[:len(v)-1])
				xCount += n
			}
		} else {
			n, _ := strconv.Atoi(v)
			num += n
		}
	}
	return xCount, num
}

func help1(exp string) (int, int) {
	// 用正则来切分表达式
	re, _ := regexp.Compile("[+-]?[0-9]*x?")
	s := re.FindAllString(exp, -1)
	xCount, num := 0, 0
	for _, v := range s {
		if len(v) > 0 && v[len(v)-1] == 'x' {
			if v == "x" || v == "+x" {
				xCount++
			} else if v == "-x" {
				xCount--
			} else {
				n, _ := strconv.Atoi(v[:len(v)-1])
				xCount += n
			}
		} else {
			n, _ := strconv.Atoi(v)
			num += n
		}
	}
	return xCount, num
}
