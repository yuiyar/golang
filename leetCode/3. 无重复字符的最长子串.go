//给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
//
//示例1:
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//示例 2:
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//示例 3:
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
//示例 4:
//输入: s = ""
//输出: 0
//
//提示：
//0 <= s.length <= 5 * (10四次方)
//s 由英文字母、数字、符号和空格组成。

package main

import (
	"fmt"
)

func main() {
	str := "22a115f11"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var endPoint, res int

	for startPoint := 0; startPoint < len(s); startPoint++ {
		if startPoint != 0 {
			delete(m, s[startPoint-1])
		}

		for endPoint < len(s) && m[s[endPoint]] == 0 {
			m[s[endPoint]]++
			endPoint++
		}

		if res < endPoint-startPoint {
			res = endPoint - startPoint
		}
	}
	return res
}
