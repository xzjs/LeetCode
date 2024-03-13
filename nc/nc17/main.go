package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param A string字符串
 * @return int整型
 */
func getLongestPalindrome(A string) int {
	// 先扩充字符串，变成奇数的串
	length := 2*len(A) + 1
	ss := make([]byte, length)
	for i := 0; i < len(ss); i++ {
		if i&1 == 0 {
			ss[i] = '#'
		} else {
			ss[i] = A[i>>1]
		}
	}
	p := make([]int, length)
	c, r := 0, 0
	maxLength := 0
	for i := 0; i < len(ss); i++ {
		// 计算镜像坐标
		mirror := (2 * c) - i
		if i < r {
			p[i] = min(r-i, p[mirror])
		}
		// 开始往两边扩展
		for a, b := i+(1+p[i]), i-(1+p[i]); a < length && b >= 0 && ss[a] == ss[b]; a, b = a+1, b-1 {
			p[i]++
		}

		// 移动c
		if i+p[i] > r {
			c = i
			r = 1 + p[i]
			if p[i] > maxLength {
				maxLength = p[i]
			}
		}
	}
	return maxLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	A := "baabccc"
	fmt.Println(getLongestPalindrome(A))
}
